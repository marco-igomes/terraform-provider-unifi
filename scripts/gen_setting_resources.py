#!/usr/bin/env python3
"""Generate one terraform-provider-unifi resource file per SDK setting type.

Reads SDK struct definitions from go-unifi/unifi/settings/*.generated.go,
emits unifi/setting_<key>_resource.go using a hand-rolled framework template.
"""
import os
import re
import sys

SDK_DIR = "/Users/migomes/Projects/migomes/external/terraform/go-unifi/unifi/settings"
OUT_DIR = "/Users/migomes/Projects/migomes/external/terraform/terraform-provider-unifi/unifi"

# (key, GoStructName, ResourceTypeName, MarkdownDescription)
TARGETS = [
    ("auto_speedtest","AutoSpeedtest","auto_speedtest","Scheduled WAN speedtest configuration."),
    ("connectivity","Connectivity","connectivity","WAN uplink connectivity monitor (used by failover logic)."),
    ("country","Country","country","Regulatory country/region code (drives radio limits and DFS)."),
    ("doh","Doh","doh","DNS-over-HTTPS upstream provider configuration."),
    ("dpi","Dpi","dpi","Deep Packet Inspection (DPI) toggles + traffic fingerprinting."),
    ("global_nat","GlobalNat","global_nat","Global NAT mode toggle."),
    ("global_network","GlobalNetwork","global_network","Site-wide default security posture (allow-all vs zero-trust)."),
    ("global_switch","GlobalSwitch","global_switch","Switch-wide defaults: STP version, DHCP snooping, RADIUS profile."),
    ("guest_access","GuestAccess","guest_access","Captive portal / guest hotspot configuration."),
    ("igmp_snooping","IgmpSnooping","igmp_snooping","Multicast (IGMP) snooping config: querier mode, subscription mode, flood behaviour."),
    ("ips","Ips","ips","IDS/IPS configuration: mode, enabled networks, categories, DNS filtering, ad-blocking."),
    ("ipsec","Ipsec","ipsec","Global IPsec parameters (e.g. IKEv2 reauthentication method)."),
    ("locale","Locale","locale","Controller timezone."),
    ("magic_site_to_site_vpn","MagicSiteToSiteVpn","magic_site_to_site_vpn","UniFi SD-WAN (Magic Site-to-Site) overlay VPN configuration."),
    ("mdns","Mdns","mdns","mDNS reflector configuration (cross-VLAN service discovery)."),
    ("netflow","Netflow","netflow","NetFlow/sFlow flow exporter configuration."),
    ("ntp","Ntp","ntp","NTP servers and manual/auto mode for controller time sync."),
    ("roaming_assistant","RoamingAssistant","roaming_assistant","WiFi roaming assistant RSSI threshold (sticky-client mitigation)."),
    ("rsyslogd","Rsyslogd","rsyslogd","Remote syslog (rsyslogd) configuration."),
    ("ssl_inspection","SslInspection","ssl_inspection","TLS/SSL inspection on/off."),
    ("traffic_flow","TrafficFlow","traffic_flow","Traffic flow capture toggles."),
]

FIELD_RE = re.compile(r"^\t([A-Z]\w+)\s+(\[?\]?\*?[A-Za-z0-9_.]+)\s+`json:\"([^\"]+)\"`(.*)$", re.M)


def snake(name: str) -> str:
    s = re.sub(r"(.)([A-Z][a-z]+)", r"\1_\2", name)
    s = re.sub(r"([a-z0-9])([A-Z])", r"\1_\2", s)
    return s.lower()


def parse_struct(path: str, struct_name: str):
    text = open(path).read()
    m = re.search(rf"type {struct_name} struct \{{(.*?)^\}}", text, re.M | re.S)
    if not m:
        raise RuntimeError(f"struct {struct_name} not found in {path}")
    body = m.group(1)
    fields = []
    for fm in FIELD_RE.finditer(body):
        name, gotype, jsontag, rest = fm.groups()
        json_name = jsontag.split(",")[0]
        comment = rest.strip().lstrip("//").strip() or ""
        # tfsdk attribute names must be snake_case lowercase alphanumeric.
        tf_name = snake(json_name) if any(c.isupper() for c in json_name) else json_name
        fields.append({"go": name, "gotype": gotype, "json": json_name, "tf": tf_name, "comment": comment})
    return fields


SIMPLE_TYPES = {"bool", "string", "int", "int64", "*int64", "[]string"}


def tf_type_for(gotype: str):
    """Return (tfsdk_go_type, schema_attribute_constructor, marshal_helpers)."""
    if gotype == "bool":
        return ("types.Bool", "schema.BoolAttribute",
                {"to_model": "types.BoolValue({src}.{Go})",
                 "to_sdk":   "{model}.{Go}.ValueBool()",
                 "plan_mod": "boolplanmodifier.UseStateForUnknown()",
                 "plan_mod_pkg": "boolplanmodifier",
                 "list_kw": ""})
    if gotype == "string":
        return ("types.String", "schema.StringAttribute",
                {"to_model": "stringOrNull({src}.{Go})",
                 "to_sdk":   "{model}.{Go}.ValueString()",
                 "plan_mod": "stringplanmodifier.UseStateForUnknown()",
                 "plan_mod_pkg": "stringplanmodifier",
                 "list_kw": ""})
    if gotype in ("int", "int64"):
        return ("types.Int64", "schema.Int64Attribute",
                {"to_model": "types.Int64Value(int64({src}.{Go}))",
                 "to_sdk":   "int({model}.{Go}.ValueInt64())" if gotype == "int" else "{model}.{Go}.ValueInt64()",
                 "plan_mod": "int64planmodifier.UseStateForUnknown()",
                 "plan_mod_pkg": "int64planmodifier",
                 "list_kw": ""})
    if gotype == "*int64":
        return ("types.Int64", "schema.Int64Attribute",
                {"to_model": "types.Int64PointerValue({src}.{Go})",
                 "to_sdk":   "int64PointerOrNil({model}.{Go})",
                 "plan_mod": "int64planmodifier.UseStateForUnknown()",
                 "plan_mod_pkg": "int64planmodifier",
                 "list_kw": ""})
    if gotype == "[]string":
        return ("types.List", "schema.ListAttribute",
                {"to_model": "stringListOrNull(ctx, {src}.{Go}, &diags)",
                 "to_sdk":   "stringSliceFromList(ctx, {model}.{Go}, &diags)",
                 "plan_mod": "listplanmodifier.UseStateForUnknown()",
                 "plan_mod_pkg": "listplanmodifier",
                 "list_kw": "ElementType: types.StringType,"})
    # complex types -> JSON-string attribute (Computed-only, lossless round trip)
    # Prefix referenced types with `settings.` package qualifier.
    qualified = gotype
    if qualified.startswith("[]"):
        qualified = "[]settings." + qualified[2:]
    elif qualified.startswith("*"):
        qualified = "*settings." + qualified[1:]
    else:
        qualified = "settings." + qualified
    return ("types.String", "schema.StringAttribute",
            {"to_model": "jsonStringFrom({src}.{Go})",
             "to_sdk":   None,  # complex -> JSON; needs custom decoding, see emit
             "plan_mod": "stringplanmodifier.UseStateForUnknown()",
             "plan_mod_pkg": "stringplanmodifier",
             "list_kw": "",
             "complex": True,
             "complex_type": qualified})


HEADER = '''package unifi

// Code generated from go-unifi/unifi/settings/{key}.generated.go.
// DO NOT EDIT MANUALLY — regenerate via scripts/gen_setting_resources.py.

import (
\t"context"
\t"encoding/json"
\t"fmt"

\t"github.com/hashicorp/terraform-plugin-framework/diag"
\t"github.com/hashicorp/terraform-plugin-framework/path"
\t"github.com/hashicorp/terraform-plugin-framework/resource"
\t"github.com/hashicorp/terraform-plugin-framework/resource/schema"
{plan_mod_imports}
\t"github.com/hashicorp/terraform-plugin-framework/types"
\tui "github.com/ubiquiti-community/go-unifi/unifi"
\t"github.com/ubiquiti-community/go-unifi/unifi/settings"
)
'''


def emit(key, struct, type_name, desc, fields):
    plan_mod_pkgs = {"stringplanmodifier"}
    for f in fields:
        tf = tf_type_for(f["gotype"])
        plan_mod_pkgs.add(tf[2]["plan_mod_pkg"])
    # also ensure list import if used
    plan_mod_imports = "\n".join(
        f'\t"github.com/hashicorp/terraform-plugin-framework/resource/schema/{p}"'
        for p in sorted(plan_mod_pkgs)
    )

    parts = [HEADER.format(key=key, plan_mod_imports=plan_mod_imports)]

    res_name = f"setting{struct}Resource"
    model_name = f"setting{struct}Model"
    new_name = f"NewSetting{struct}Resource"

    parts.append(f"""
var (
\t_ resource.Resource                = &{res_name}{{}}
\t_ resource.ResourceWithImportState = &{res_name}{{}}
)

func {new_name}() resource.Resource {{
\treturn &{res_name}{{}}
}}

type {res_name} struct {{
\tclient *Client
}}

type {model_name} struct {{
\tID   types.String `tfsdk:"id"`
\tSite types.String `tfsdk:"site"`
""")

    for f in fields:
        tfsdk, _, helpers = tf_type_for(f["gotype"])
        parts.append(f"\t{f['go']} {tfsdk} `tfsdk:\"{f['tf']}\"`\n")
    parts.append("}\n")

    # Metadata
    parts.append(f"""
func (r *{res_name}) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {{
\tresp.TypeName = req.ProviderTypeName + "_setting_{type_name}"
}}

func (r *{res_name}) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {{
\tresp.Schema = schema.Schema{{
\t\tMarkdownDescription: "{desc}",
\t\tAttributes: map[string]schema.Attribute{{
\t\t\t"id": schema.StringAttribute{{Computed: true, PlanModifiers: []planmodifier.String{{stringplanmodifier.UseStateForUnknown()}}}},
\t\t\t"site": schema.StringAttribute{{
\t\t\t\tOptional: true, Computed: true,
\t\t\t\tPlanModifiers: []planmodifier.String{{stringplanmodifier.RequiresReplace(), stringplanmodifier.UseStateForUnknown()}},
\t\t\t}},
""")

    # Per-field schema
    for f in fields:
        _, attr_ctor, helpers = tf_type_for(f["gotype"])
        comment = f["comment"].replace("\\", "\\\\").replace('"', '\\"').strip()
        if not comment:
            comment = f"{f['json']} field"
        pm_pkg = helpers["plan_mod_pkg"]
        pm_type = {"boolplanmodifier": "Bool", "stringplanmodifier": "String",
                   "int64planmodifier": "Int64", "listplanmodifier": "List"}[pm_pkg]
        list_kw = helpers["list_kw"]
        list_kw_line = f"\n\t\t\t\t{list_kw}" if list_kw else ""
        parts.append(f'\t\t\t"{f["tf"]}": {attr_ctor}{{\n')
        parts.append(f'\t\t\t\tMarkdownDescription: "{comment}",\n')
        parts.append(f"\t\t\t\tOptional: true, Computed: true,{list_kw_line}\n")
        parts.append(f"\t\t\t\tPlanModifiers: []planmodifier.{pm_type}{{{pm_pkg}.UseStateForUnknown()}},\n")
        parts.append("\t\t\t},\n")

    parts.append("""\t\t},
\t}
}

""")

    # planmodifier base import
    parts.append("""func (r *""" + res_name + """) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
\tif req.ProviderData == nil { return }
\tclient, ok := req.ProviderData.(*Client)
\tif !ok {
\t\tresp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *Client, got: %T.", req.ProviderData))
\t\treturn
\t}
\tr.client = client
}

func (r *""" + res_name + """) siteOrDefault(site types.String) string {
\tif s := site.ValueString(); s != "" { return s }
\treturn r.client.Site
}

""")

    # CRUD: Create/Update both write via UpdateSetting; Read uses GetSetting; Delete drops state (settings can't be deleted)
    parts.append(f"""func (r *{res_name}) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {{
\tvar plan {model_name}
\tresp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
\tif resp.Diagnostics.HasError() {{ return }}
\tsite := r.siteOrDefault(plan.Site)
\tresp.Diagnostics.Append(r.writeAndRefresh(ctx, site, &plan)...)
\tif resp.Diagnostics.HasError() {{ return }}
\tresp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}}

func (r *{res_name}) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {{
\tvar state {model_name}
\tresp.Diagnostics.Append(req.State.Get(ctx, &state)...)
\tif resp.Diagnostics.HasError() {{ return }}
\tsite := r.siteOrDefault(state.Site)

\tmeta, current, err := ui.GetSetting[*settings.{struct}](r.client.ApiClient, ctx, site)
\tif err != nil {{
\t\tif _, ok := err.(*ui.NotFoundError); ok {{ resp.State.RemoveResource(ctx); return }}
\t\tresp.Diagnostics.AddError("Error Reading {struct} Setting", err.Error())
\t\treturn
\t}}
\tresp.Diagnostics.Append(r.settingToModel(ctx, meta, current, &state, site)...)
\tresp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}}

func (r *{res_name}) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {{
\tvar plan, state {model_name}
\tresp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
\tresp.Diagnostics.Append(req.State.Get(ctx, &state)...)
\tif resp.Diagnostics.HasError() {{ return }}
\tsite := r.siteOrDefault(state.Site)
\tresp.Diagnostics.Append(r.writeAndRefresh(ctx, site, &plan)...)
\tif resp.Diagnostics.HasError() {{ return }}
\tresp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}}

func (r *{res_name}) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {{
\t// Settings cannot be deleted from the controller; dropping from state only.
}}

func (r *{res_name}) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {{
\tresource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}}

func (r *{res_name}) writeAndRefresh(ctx context.Context, site string, m *{model_name}) diag.Diagnostics {{
\tvar diags diag.Diagnostics

\t// Read current so unmodified fields aren't dropped to zero values.
\t_, current, err := ui.GetSetting[*settings.{struct}](r.client.ApiClient, ctx, site)
\tif err != nil {{
\t\tif _, ok := err.(*ui.NotFoundError); !ok {{
\t\t\tdiags.AddError("Error Reading {struct} Setting", err.Error())
\t\t\treturn diags
\t\t}}
\t\tcurrent = &settings.{struct}{{}}
\t}}

\tr.applyModelToSetting(ctx, m, current, &diags)
\tif diags.HasError() {{ return diags }}

\tif err := r.client.UpdateSetting(ctx, site, current); err != nil {{
\t\tdiags.AddError("Error Updating {struct} Setting", err.Error())
\t\treturn diags
\t}}

\tmeta, refreshed, err := ui.GetSetting[*settings.{struct}](r.client.ApiClient, ctx, site)
\tif err != nil {{
\t\tdiags.AddError("Error Re-reading {struct} Setting", err.Error())
\t\treturn diags
\t}}
\tdiags.Append(r.settingToModel(ctx, meta, refreshed, m, site)...)
\treturn diags
}}

""")

    # settingToModel
    parts.append(f"func (r *{res_name}) settingToModel(ctx context.Context, meta *ui.Setting, s *settings.{struct}, m *{model_name}, site string) diag.Diagnostics {{\n")
    parts.append("\tvar diags diag.Diagnostics\n")
    parts.append("\tif meta != nil { m.ID = types.StringValue(meta.Id) }\n")
    parts.append("\tm.Site = types.StringValue(site)\n")
    parts.append("\t_ = ctx\n\t_ = diags\n")
    for f in fields:
        _, _, helpers = tf_type_for(f["gotype"])
        expr = helpers["to_model"].format(src="s", Go=f["go"], model="m")
        parts.append(f"\tm.{f['go']} = {expr}\n")
    parts.append("\treturn diags\n}\n\n")

    # applyModelToSetting
    # Only overwrite SDK fields where the model has a real value (not null/unknown).
    # This way fields the user did not declare retain their controller-current value
    # (read into the SDK struct beforehand by writeAndRefresh).
    parts.append(f"func (r *{res_name}) applyModelToSetting(ctx context.Context, m *{model_name}, s *settings.{struct}, diags *diag.Diagnostics) {{\n")
    parts.append("\t_ = ctx\n\t_ = diags\n")
    for f in fields:
        _, _, helpers = tf_type_for(f["gotype"])
        if helpers.get("complex"):
            ctype = helpers["complex_type"]
            parts.append(f"\tif !m.{f['go']}.IsNull() && !m.{f['go']}.IsUnknown() {{\n")
            parts.append(f"\t\tvar val {ctype}\n")
            parts.append(f"\t\tif err := json.Unmarshal([]byte(m.{f['go']}.ValueString()), &val); err != nil {{\n")
            parts.append(f"\t\t\tdiags.AddError(\"Invalid {f['tf']}\", err.Error())\n")
            parts.append(f"\t\t}} else {{ s.{f['go']} = val }}\n")
            parts.append("\t}\n")
        elif f["gotype"] == "[]string":
            parts.append(f"\tif !m.{f['go']}.IsNull() && !m.{f['go']}.IsUnknown() {{\n")
            parts.append(f"\t\tvar v []string\n")
            parts.append(f"\t\tdiags.Append(m.{f['go']}.ElementsAs(ctx, &v, false)...)\n")
            parts.append(f"\t\ts.{f['go']} = v\n")
            parts.append("\t}\n")
        else:
            expr = helpers["to_sdk"].format(model="m", Go=f["go"])
            parts.append(f"\tif !m.{f['go']}.IsNull() && !m.{f['go']}.IsUnknown() {{ s.{f['go']} = {expr} }}\n")
    parts.append("}\n")

    # Imports tidy: if no complex fields, json/diag may be unused; reference them via _ stmts
    has_complex = any(tf_type_for(f["gotype"])[2].get("complex") for f in fields)
    has_list = any(f["gotype"] == "[]string" for f in fields)
    if not has_complex:
        # use a no-op var to keep json imported is wasteful; instead let's only import json when needed
        pass

    return "".join(parts), has_complex, has_list, plan_mod_pkgs


def main():
    os.makedirs(OUT_DIR, exist_ok=True)
    new_funcs = []
    for key, struct, type_name, desc in TARGETS:
        spath = os.path.join(SDK_DIR, f"{key}.generated.go")
        fields = parse_struct(spath, struct)
        code, has_complex, has_list, plan_mod_pkgs = emit(key, struct, type_name, desc, fields)

        # rewrite header to drop unused imports
        imports = ['\t"context"', '\t"fmt"', '\t"github.com/hashicorp/terraform-plugin-framework/diag"',
                   '\t"github.com/hashicorp/terraform-plugin-framework/path"',
                   '\t"github.com/hashicorp/terraform-plugin-framework/resource"',
                   '\t"github.com/hashicorp/terraform-plugin-framework/resource/schema"',
                   '\t"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"']
        for p in sorted(plan_mod_pkgs):
            imports.append(f'\t"github.com/hashicorp/terraform-plugin-framework/resource/schema/{p}"')
        imports.append('\t"github.com/hashicorp/terraform-plugin-framework/types"')
        imports.append('\tui "github.com/ubiquiti-community/go-unifi/unifi"')
        imports.append('\t"github.com/ubiquiti-community/go-unifi/unifi/settings"')
        if has_complex:
            imports.insert(2, '\t"encoding/json"')
        header = f"package unifi\n\n// Generated from go-unifi/unifi/settings/{key}.generated.go.\n// DO NOT EDIT MANUALLY — regenerate via scripts/gen_setting_resources.py.\n\nimport (\n" + "\n".join(imports) + "\n)\n"
        # strip the inline header from emit() and prepend our cleaned one
        body = code.split(")\n", 1)[1]
        out = header + body

        outpath = os.path.join(OUT_DIR, f"setting_{key}_resource.go")
        open(outpath, "w").write(out)
        new_funcs.append(f"NewSetting{struct}Resource")
        print(f"WROTE {outpath}  ({len(fields)} fields)")

    print()
    print("Add to provider.go Resources():")
    for n in new_funcs:
        print(f"\t\t{n},")


if __name__ == "__main__":
    main()
