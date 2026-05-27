package unifi

// Generated from go-unifi/unifi/settings/global_switch.generated.go.
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
	_ resource.Resource                = &settingGlobalSwitchResource{}
	_ resource.ResourceWithImportState = &settingGlobalSwitchResource{}
)

func NewSettingGlobalSwitchResource() resource.Resource {
	return &settingGlobalSwitchResource{}
}

type settingGlobalSwitchResource struct {
	client *Client
}

type settingGlobalSwitchModel struct {
	ID                             types.String `tfsdk:"id"`
	Site                           types.String `tfsdk:"site"`
	AclDeviceIsolation             types.List   `tfsdk:"acl_device_isolation"`
	AclL3Isolation                 types.String `tfsdk:"acl_l3_isolation"`
	DHCPSnoop                      types.Bool   `tfsdk:"dhcp_snoop"`
	Dot1XFallbackNetworkID         types.String `tfsdk:"dot1x_fallback_networkconf_id"`
	Dot1XPortctrlEnabled           types.Bool   `tfsdk:"dot1x_portctrl_enabled"`
	FloodKnownProtocols            types.Bool   `tfsdk:"flood_known_protocols"`
	FlowctrlEnabled                types.Bool   `tfsdk:"flowctrl_enabled"`
	ForwardUnknownMcastRouterPorts types.Bool   `tfsdk:"forward_unknown_mcast_router_ports"`
	JumboframeEnabled              types.Bool   `tfsdk:"jumboframe_enabled"`
	RADIUSProfileID                types.String `tfsdk:"radiusprofile_id"`
	StpVersion                     types.String `tfsdk:"stp_version"`
	SwitchExclusions               types.List   `tfsdk:"switch_exclusions"`
}

func (r *settingGlobalSwitchResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_setting_global_switch"
}

func (r *settingGlobalSwitchResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Switch-wide defaults: STP version, DHCP snooping, RADIUS profile.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"site": schema.StringAttribute{
				Optional: true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace(), stringplanmodifier.UseStateForUnknown()},
			},
			"acl_device_isolation": schema.ListAttribute{
				MarkdownDescription: "acl_device_isolation field",
				Optional:            true, Computed: true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"acl_l3_isolation": schema.StringAttribute{
				MarkdownDescription: "acl_l3_isolation field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"dhcp_snoop": schema.BoolAttribute{
				MarkdownDescription: "dhcp_snoop field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"dot1x_fallback_networkconf_id": schema.StringAttribute{
				MarkdownDescription: "[\\d\\w]+|",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"dot1x_portctrl_enabled": schema.BoolAttribute{
				MarkdownDescription: "dot1x_portctrl_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"flood_known_protocols": schema.BoolAttribute{
				MarkdownDescription: "flood_known_protocols field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"flowctrl_enabled": schema.BoolAttribute{
				MarkdownDescription: "flowctrl_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"forward_unknown_mcast_router_ports": schema.BoolAttribute{
				MarkdownDescription: "forward_unknown_mcast_router_ports field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"jumboframe_enabled": schema.BoolAttribute{
				MarkdownDescription: "jumboframe_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"radiusprofile_id": schema.StringAttribute{
				MarkdownDescription: "radiusprofile_id field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"stp_version": schema.StringAttribute{
				MarkdownDescription: "stp|rstp|disabled",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"switch_exclusions": schema.ListAttribute{
				MarkdownDescription: "^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$",
				Optional:            true, Computed: true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *settingGlobalSwitchResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *settingGlobalSwitchResource) siteOrDefault(site types.String) string {
	if s := site.ValueString(); s != "" {
		return s
	}
	return r.client.Site
}

func (r *settingGlobalSwitchResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan settingGlobalSwitchModel
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

func (r *settingGlobalSwitchResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state settingGlobalSwitchModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	site := r.siteOrDefault(state.Site)

	meta, current, err := ui.GetSetting[*settings.GlobalSwitch](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading GlobalSwitch Setting", err.Error())
		return
	}
	resp.Diagnostics.Append(r.settingToModel(ctx, meta, current, &state, site)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *settingGlobalSwitchResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state settingGlobalSwitchModel
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

func (r *settingGlobalSwitchResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Settings cannot be deleted from the controller; dropping from state only.
}

func (r *settingGlobalSwitchResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *settingGlobalSwitchResource) writeAndRefresh(ctx context.Context, site string, m *settingGlobalSwitchModel) diag.Diagnostics {
	var diags diag.Diagnostics

	// Read current so unmodified fields aren't dropped to zero values.
	_, current, err := ui.GetSetting[*settings.GlobalSwitch](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); !ok {
			diags.AddError("Error Reading GlobalSwitch Setting", err.Error())
			return diags
		}
		current = &settings.GlobalSwitch{}
	}

	r.applyModelToSetting(ctx, m, current, &diags)
	if diags.HasError() {
		return diags
	}

	if err := r.client.UpdateSetting(ctx, site, current); err != nil {
		diags.AddError("Error Updating GlobalSwitch Setting", err.Error())
		return diags
	}

	meta, refreshed, err := ui.GetSetting[*settings.GlobalSwitch](r.client.ApiClient, ctx, site)
	if err != nil {
		diags.AddError("Error Re-reading GlobalSwitch Setting", err.Error())
		return diags
	}
	diags.Append(r.settingToModel(ctx, meta, refreshed, m, site)...)
	return diags
}

func (r *settingGlobalSwitchResource) settingToModel(ctx context.Context, meta *ui.Setting, s *settings.GlobalSwitch, m *settingGlobalSwitchModel, site string) diag.Diagnostics {
	var diags diag.Diagnostics
	if meta != nil {
		m.ID = types.StringValue(meta.Id)
	}
	m.Site = types.StringValue(site)
	_ = ctx
	_ = diags
	m.AclDeviceIsolation = stringListOrNull(ctx, s.AclDeviceIsolation, &diags)
	m.AclL3Isolation = jsonStringFrom(s.AclL3Isolation)
	m.DHCPSnoop = types.BoolValue(s.DHCPSnoop)
	m.Dot1XFallbackNetworkID = stringOrNull(s.Dot1XFallbackNetworkID)
	m.Dot1XPortctrlEnabled = types.BoolValue(s.Dot1XPortctrlEnabled)
	m.FloodKnownProtocols = types.BoolValue(s.FloodKnownProtocols)
	m.FlowctrlEnabled = types.BoolValue(s.FlowctrlEnabled)
	m.ForwardUnknownMcastRouterPorts = types.BoolValue(s.ForwardUnknownMcastRouterPorts)
	m.JumboframeEnabled = types.BoolValue(s.JumboframeEnabled)
	m.RADIUSProfileID = stringOrNull(s.RADIUSProfileID)
	m.StpVersion = stringOrNull(s.StpVersion)
	m.SwitchExclusions = stringListOrNull(ctx, s.SwitchExclusions, &diags)
	return diags
}

func (r *settingGlobalSwitchResource) applyModelToSetting(ctx context.Context, m *settingGlobalSwitchModel, s *settings.GlobalSwitch, diags *diag.Diagnostics) {
	_ = ctx
	_ = diags
	if !m.AclDeviceIsolation.IsNull() && !m.AclDeviceIsolation.IsUnknown() {
		var v []string
		diags.Append(m.AclDeviceIsolation.ElementsAs(ctx, &v, false)...)
		s.AclDeviceIsolation = v
	}
	if !m.AclL3Isolation.IsNull() && !m.AclL3Isolation.IsUnknown() {
		var val []settings.SettingGlobalSwitchAclL3Isolation
		if err := json.Unmarshal([]byte(m.AclL3Isolation.ValueString()), &val); err != nil {
			diags.AddError("Invalid acl_l3_isolation", err.Error())
		} else {
			s.AclL3Isolation = val
		}
	}
	if !m.DHCPSnoop.IsNull() && !m.DHCPSnoop.IsUnknown() {
		s.DHCPSnoop = m.DHCPSnoop.ValueBool()
	}
	if !m.Dot1XFallbackNetworkID.IsNull() && !m.Dot1XFallbackNetworkID.IsUnknown() {
		s.Dot1XFallbackNetworkID = m.Dot1XFallbackNetworkID.ValueString()
	}
	if !m.Dot1XPortctrlEnabled.IsNull() && !m.Dot1XPortctrlEnabled.IsUnknown() {
		s.Dot1XPortctrlEnabled = m.Dot1XPortctrlEnabled.ValueBool()
	}
	if !m.FloodKnownProtocols.IsNull() && !m.FloodKnownProtocols.IsUnknown() {
		s.FloodKnownProtocols = m.FloodKnownProtocols.ValueBool()
	}
	if !m.FlowctrlEnabled.IsNull() && !m.FlowctrlEnabled.IsUnknown() {
		s.FlowctrlEnabled = m.FlowctrlEnabled.ValueBool()
	}
	if !m.ForwardUnknownMcastRouterPorts.IsNull() && !m.ForwardUnknownMcastRouterPorts.IsUnknown() {
		s.ForwardUnknownMcastRouterPorts = m.ForwardUnknownMcastRouterPorts.ValueBool()
	}
	if !m.JumboframeEnabled.IsNull() && !m.JumboframeEnabled.IsUnknown() {
		s.JumboframeEnabled = m.JumboframeEnabled.ValueBool()
	}
	if !m.RADIUSProfileID.IsNull() && !m.RADIUSProfileID.IsUnknown() {
		s.RADIUSProfileID = m.RADIUSProfileID.ValueString()
	}
	if !m.StpVersion.IsNull() && !m.StpVersion.IsUnknown() {
		s.StpVersion = m.StpVersion.ValueString()
	}
	if !m.SwitchExclusions.IsNull() && !m.SwitchExclusions.IsUnknown() {
		var v []string
		diags.Append(m.SwitchExclusions.ElementsAs(ctx, &v, false)...)
		s.SwitchExclusions = v
	}
}
