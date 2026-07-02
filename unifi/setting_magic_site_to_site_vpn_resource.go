package unifi

// Generated from go-unifi/unifi/settings/magic_site_to_site_vpn.generated.go.
// DO NOT EDIT MANUALLY — regenerate via scripts/gen_setting_resources.py.

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ui "github.com/ubiquiti-community/go-unifi/unifi"
	"github.com/ubiquiti-community/go-unifi/unifi/settings"
)

var (
	_ resource.Resource                = &settingMagicSiteToSiteVpnResource{}
	_ resource.ResourceWithImportState = &settingMagicSiteToSiteVpnResource{}
)

func NewSettingMagicSiteToSiteVpnResource() resource.Resource {
	return &settingMagicSiteToSiteVpnResource{}
}

type settingMagicSiteToSiteVpnResource struct {
	client *Client
}

type settingMagicSiteToSiteVpnModel struct {
	ID          types.String `tfsdk:"id"`
	Site        types.String `tfsdk:"site"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	PublicKey   types.String `tfsdk:"public_key"`
	XPrivateKey types.String `tfsdk:"x_private_key"`
}

func (r *settingMagicSiteToSiteVpnResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_setting_magic_site_to_site_vpn"
}

func (r *settingMagicSiteToSiteVpnResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "UniFi SD-WAN (Magic Site-to-Site) overlay VPN configuration.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"site": schema.StringAttribute{
				Optional: true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace(), stringplanmodifier.UseStateForUnknown()},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether the UID-Magic site-to-site VPN feature is enabled on this site.",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"public_key": schema.StringAttribute{
				MarkdownDescription: "Controller-generated.",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"x_private_key": schema.StringAttribute{
				MarkdownDescription: "Controller-generated.",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *settingMagicSiteToSiteVpnResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *Client, got: %T.", req.ProviderData))
		return
	}
	r.client = client
}

func (r *settingMagicSiteToSiteVpnResource) siteOrDefault(site types.String) string {
	if s := site.ValueString(); s != "" {
		return s
	}
	return r.client.Site
}

func (r *settingMagicSiteToSiteVpnResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan settingMagicSiteToSiteVpnModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	site := r.siteOrDefault(plan.Site)
	resp.Diagnostics.Append(r.writeAndRefresh(ctx, site, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *settingMagicSiteToSiteVpnResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state settingMagicSiteToSiteVpnModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	site := r.siteOrDefault(state.Site)

	meta, current, err := ui.GetSetting[*settings.MagicSiteToSiteVpn](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading MagicSiteToSiteVpn Setting", err.Error())
		return
	}
	resp.Diagnostics.Append(r.settingToModel(ctx, meta, current, &state, site)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *settingMagicSiteToSiteVpnResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state settingMagicSiteToSiteVpnModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	site := r.siteOrDefault(state.Site)
	resp.Diagnostics.Append(r.writeAndRefresh(ctx, site, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *settingMagicSiteToSiteVpnResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Settings cannot be deleted from the controller; dropping from state only.
}

func (r *settingMagicSiteToSiteVpnResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *settingMagicSiteToSiteVpnResource) writeAndRefresh(ctx context.Context, site string, m *settingMagicSiteToSiteVpnModel) diag.Diagnostics {
	var diags diag.Diagnostics

	// Read current so unmodified fields aren't dropped to zero values.
	_, current, err := ui.GetSetting[*settings.MagicSiteToSiteVpn](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); !ok {
			diags.AddError("Error Reading MagicSiteToSiteVpn Setting", err.Error())
			return diags
		}
		current = &settings.MagicSiteToSiteVpn{}
	}

	r.applyModelToSetting(ctx, m, current, &diags)
	if diags.HasError() {
		return diags
	}

	if err := r.client.UpdateSetting(ctx, site, current); err != nil {
		diags.AddError("Error Updating MagicSiteToSiteVpn Setting", err.Error())
		return diags
	}

	meta, refreshed, err := ui.GetSetting[*settings.MagicSiteToSiteVpn](r.client.ApiClient, ctx, site)
	if err != nil {
		diags.AddError("Error Re-reading MagicSiteToSiteVpn Setting", err.Error())
		return diags
	}
	diags.Append(r.settingToModel(ctx, meta, refreshed, m, site)...)
	return diags
}

func (r *settingMagicSiteToSiteVpnResource) settingToModel(ctx context.Context, meta *ui.Setting, s *settings.MagicSiteToSiteVpn, m *settingMagicSiteToSiteVpnModel, site string) diag.Diagnostics {
	var diags diag.Diagnostics
	if meta != nil {
		m.ID = types.StringValue(meta.Id)
	}
	m.Site = types.StringValue(site)
	_ = ctx
	_ = diags
	m.Enabled = types.BoolValue(s.Enabled)
	m.PublicKey = stringOrNull(s.PublicKey)
	m.XPrivateKey = stringOrNull(s.XPrivateKey)
	return diags
}

func (r *settingMagicSiteToSiteVpnResource) applyModelToSetting(ctx context.Context, m *settingMagicSiteToSiteVpnModel, s *settings.MagicSiteToSiteVpn, diags *diag.Diagnostics) {
	_ = ctx
	_ = diags
	if !m.Enabled.IsNull() && !m.Enabled.IsUnknown() {
		s.Enabled = m.Enabled.ValueBool()
	}
	if !m.PublicKey.IsNull() && !m.PublicKey.IsUnknown() {
		s.PublicKey = m.PublicKey.ValueString()
	}
	if !m.XPrivateKey.IsNull() && !m.XPrivateKey.IsUnknown() {
		s.XPrivateKey = m.XPrivateKey.ValueString()
	}
}
