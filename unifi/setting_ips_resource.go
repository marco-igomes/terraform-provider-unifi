package unifi

// Generated from go-unifi/unifi/settings/ips.generated.go.
// DO NOT EDIT MANUALLY — regenerate via scripts/gen_setting_resources.py.

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ui "github.com/ubiquiti-community/go-unifi/unifi"
	"github.com/ubiquiti-community/go-unifi/unifi/settings"
)

var (
	_ resource.Resource                = &settingIpsResource{}
	_ resource.ResourceWithImportState = &settingIpsResource{}
)

func NewSettingIpsResource() resource.Resource {
	return &settingIpsResource{}
}

type settingIpsResource struct {
	client *Client
}

type settingIpsModel struct {
	ID                                  types.String `tfsdk:"id"`
	Site                                types.String `tfsdk:"site"`
	AdBlockingConfigurations            types.String `tfsdk:"ad_blocking_configurations"`
	AdvancedFilteringPreference         types.String `tfsdk:"advanced_filtering_preference"`
	ContentFilteringBlockingPageEnabled types.Bool   `tfsdk:"content_filtering_blocking_page_enabled"`
	DnsFiltering                        types.Bool   `tfsdk:"dns_filtering"`
	DnsFilters                          types.String `tfsdk:"dns_filters"`
	EnabledCategories                   types.List   `tfsdk:"enabled_categories"`
	EnabledNetworks                     types.List   `tfsdk:"enabled_networks"`
	Honeypot                            types.String `tfsdk:"honeypot"`
	HoneypotEnabled                     types.Bool   `tfsdk:"honeypot_enabled"`
	IPsMode                             types.String `tfsdk:"ips_mode"`
	MemoryOptimized                     types.Bool   `tfsdk:"memory_optimized"`
	RestrictTorrents                    types.Bool   `tfsdk:"restrict_torrents"`
	Suppression                         types.String `tfsdk:"suppression"`
}

func (r *settingIpsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_setting_ips"
}

func (r *settingIpsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "IDS/IPS configuration: mode, enabled networks, categories, DNS filtering, ad-blocking.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"site": schema.StringAttribute{
				Optional: true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace(), stringplanmodifier.UseStateForUnknown()},
			},
			"ad_blocking_configurations": schema.StringAttribute{
				MarkdownDescription: "ad_blocking_configurations field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"advanced_filtering_preference": schema.StringAttribute{
				MarkdownDescription: "|manual|disabled",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"content_filtering_blocking_page_enabled": schema.BoolAttribute{
				MarkdownDescription: "Show the UniFi blocking page when content filtering blocks a request.",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"dns_filtering": schema.BoolAttribute{
				MarkdownDescription: "Whether per-network DNS content filtering is enabled.",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"dns_filters": schema.StringAttribute{
				MarkdownDescription: "dns_filters field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"enabled_categories": schema.ListAttribute{
				MarkdownDescription: "emerging-activex|emerging-attackresponse|botcc|emerging-chat|ciarmy|compromised|emerging-dns|emerging-dos|dshield|emerging-exploit|emerging-ftp|emerging-games|emerging-icmp|emerging-icmpinfo|emerging-imap|emerging-inappropriate|emerging-info|emerging-malware|emerging-misc|emerging-mobile|emerging-netbios|emerging-p2p|emerging-policy|emerging-pop3|emerging-rpc|emerging-scada|emerging-scan|emerging-shellcode|emerging-smtp|emerging-snmp|emerging-sql|emerging-telnet|emerging-tftp|tor|emerging-useragent|emerging-voip|emerging-webapps|emerging-webclient|emerging-webserver|emerging-worm|exploit-kit|adware-pup|botcc-portgrouped|phishing|threatview-cs-c2|3coresec|chat|coinminer|current-events|drop|hunting|icmp-info|inappropriate|info|ja3|policy|scada|dark-web-blocker-list|malicious-hosts",
				Optional:            true, Computed: true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"enabled_networks": schema.ListAttribute{
				MarkdownDescription: "enabled_networks field",
				Optional:            true, Computed: true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"honeypot": schema.StringAttribute{
				MarkdownDescription: "honeypot field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"honeypot_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether the IDS/IPS honeypot is enabled.",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"ips_mode": schema.StringAttribute{
				MarkdownDescription: "ids|ips|ipsInline|disabled",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"memory_optimized": schema.BoolAttribute{
				MarkdownDescription: "Use the memory-optimised ruleset path (smaller rule footprint, slightly fewer signatures).",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"restrict_torrents": schema.BoolAttribute{
				MarkdownDescription: "IDS/IPS: block BitTorrent / P2P traffic.",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"suppression": schema.StringAttribute{
				MarkdownDescription: "suppression field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *settingIpsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *settingIpsResource) siteOrDefault(site types.String) string {
	if s := site.ValueString(); s != "" {
		return s
	}
	return r.client.Site
}

func (r *settingIpsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan settingIpsModel
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

func (r *settingIpsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state settingIpsModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	site := r.siteOrDefault(state.Site)

	meta, current, err := ui.GetSetting[*settings.Ips](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading Ips Setting", err.Error())
		return
	}
	resp.Diagnostics.Append(r.settingToModel(ctx, meta, current, &state, site)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *settingIpsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state settingIpsModel
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

func (r *settingIpsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Settings cannot be deleted from the controller; dropping from state only.
}

func (r *settingIpsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *settingIpsResource) writeAndRefresh(ctx context.Context, site string, m *settingIpsModel) diag.Diagnostics {
	var diags diag.Diagnostics

	// Read current so unmodified fields aren't dropped to zero values.
	_, current, err := ui.GetSetting[*settings.Ips](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); !ok {
			diags.AddError("Error Reading Ips Setting", err.Error())
			return diags
		}
		current = &settings.Ips{}
	}

	r.applyModelToSetting(ctx, m, current, &diags)
	if diags.HasError() {
		return diags
	}

	if err := r.client.UpdateSetting(ctx, site, current); err != nil {
		diags.AddError("Error Updating Ips Setting", err.Error())
		return diags
	}

	meta, refreshed, err := ui.GetSetting[*settings.Ips](r.client.ApiClient, ctx, site)
	if err != nil {
		diags.AddError("Error Re-reading Ips Setting", err.Error())
		return diags
	}
	diags.Append(r.settingToModel(ctx, meta, refreshed, m, site)...)
	return diags
}

func (r *settingIpsResource) settingToModel(ctx context.Context, meta *ui.Setting, s *settings.Ips, m *settingIpsModel, site string) diag.Diagnostics {
	var diags diag.Diagnostics
	if meta != nil {
		m.ID = types.StringValue(meta.Id)
	}
	m.Site = types.StringValue(site)
	_ = ctx
	_ = diags
	m.AdBlockingConfigurations = jsonStringFrom(s.AdBlockingConfigurations)
	m.AdvancedFilteringPreference = stringOrNull(s.AdvancedFilteringPreference)
	m.ContentFilteringBlockingPageEnabled = types.BoolValue(s.ContentFilteringBlockingPageEnabled)
	m.DnsFiltering = types.BoolValue(s.DnsFiltering)
	m.DnsFilters = jsonStringFrom(s.DnsFilters)
	m.EnabledCategories = stringList(ctx, s.EnabledCategories, &diags)
	m.EnabledNetworks = stringList(ctx, s.EnabledNetworks, &diags)
	m.Honeypot = jsonStringFrom(s.Honeypot)
	m.HoneypotEnabled = types.BoolValue(s.HoneypotEnabled)
	m.IPsMode = stringOrNull(s.IPsMode)
	m.MemoryOptimized = types.BoolValue(s.MemoryOptimized)
	m.RestrictTorrents = types.BoolValue(s.RestrictTorrents)
	m.Suppression = jsonStringFrom(s.Suppression)
	return diags
}

func (r *settingIpsResource) applyModelToSetting(ctx context.Context, m *settingIpsModel, s *settings.Ips, diags *diag.Diagnostics) {
	_ = ctx
	_ = diags
	if !m.AdBlockingConfigurations.IsNull() && !m.AdBlockingConfigurations.IsUnknown() {
		var val []settings.SettingIpsAdBlocking
		if err := json.Unmarshal([]byte(m.AdBlockingConfigurations.ValueString()), &val); err != nil {
			diags.AddError("Invalid ad_blocking_configurations", err.Error())
		} else {
			s.AdBlockingConfigurations = val
		}
	}
	if !m.AdvancedFilteringPreference.IsNull() && !m.AdvancedFilteringPreference.IsUnknown() {
		s.AdvancedFilteringPreference = m.AdvancedFilteringPreference.ValueString()
	}
	if !m.ContentFilteringBlockingPageEnabled.IsNull() && !m.ContentFilteringBlockingPageEnabled.IsUnknown() {
		s.ContentFilteringBlockingPageEnabled = m.ContentFilteringBlockingPageEnabled.ValueBool()
	}
	if !m.DnsFiltering.IsNull() && !m.DnsFiltering.IsUnknown() {
		s.DnsFiltering = m.DnsFiltering.ValueBool()
	}
	if !m.DnsFilters.IsNull() && !m.DnsFilters.IsUnknown() {
		var val []settings.SettingIpsDnsFilter
		if err := json.Unmarshal([]byte(m.DnsFilters.ValueString()), &val); err != nil {
			diags.AddError("Invalid dns_filters", err.Error())
		} else {
			s.DnsFilters = val
		}
	}
	if !m.EnabledCategories.IsNull() && !m.EnabledCategories.IsUnknown() {
		var v []string
		diags.Append(m.EnabledCategories.ElementsAs(ctx, &v, false)...)
		s.EnabledCategories = v
	}
	if !m.EnabledNetworks.IsNull() && !m.EnabledNetworks.IsUnknown() {
		var v []string
		diags.Append(m.EnabledNetworks.ElementsAs(ctx, &v, false)...)
		s.EnabledNetworks = v
	}
	if !m.Honeypot.IsNull() && !m.Honeypot.IsUnknown() {
		var val []settings.SettingIpsHoneypot
		if err := json.Unmarshal([]byte(m.Honeypot.ValueString()), &val); err != nil {
			diags.AddError("Invalid honeypot", err.Error())
		} else {
			s.Honeypot = val
		}
	}
	if !m.HoneypotEnabled.IsNull() && !m.HoneypotEnabled.IsUnknown() {
		s.HoneypotEnabled = m.HoneypotEnabled.ValueBool()
	}
	if !m.IPsMode.IsNull() && !m.IPsMode.IsUnknown() {
		s.IPsMode = m.IPsMode.ValueString()
	}
	if !m.MemoryOptimized.IsNull() && !m.MemoryOptimized.IsUnknown() {
		s.MemoryOptimized = m.MemoryOptimized.ValueBool()
	}
	if !m.RestrictTorrents.IsNull() && !m.RestrictTorrents.IsUnknown() {
		s.RestrictTorrents = m.RestrictTorrents.ValueBool()
	}
	if !m.Suppression.IsNull() && !m.Suppression.IsUnknown() {
		var val *settings.SettingIpsSuppression
		if err := json.Unmarshal([]byte(m.Suppression.ValueString()), &val); err != nil {
			diags.AddError("Invalid suppression", err.Error())
		} else {
			s.Suppression = val
		}
	}
}
