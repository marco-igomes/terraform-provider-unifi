package unifi

// Generated from go-unifi/unifi/settings/igmp_snooping.generated.go.
// DO NOT EDIT MANUALLY — regenerate via scripts/gen_setting_resources.py.

import (
	"context"
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
	_ resource.Resource                = &settingIgmpSnoopingResource{}
	_ resource.ResourceWithImportState = &settingIgmpSnoopingResource{}
)

func NewSettingIgmpSnoopingResource() resource.Resource {
	return &settingIgmpSnoopingResource{}
}

type settingIgmpSnoopingResource struct {
	client *Client
}

type settingIgmpSnoopingModel struct {
	ID                                 types.String `tfsdk:"id"`
	Site                               types.String `tfsdk:"site"`
	Enabled                            types.Bool   `tfsdk:"enabled"`
	FailoverQuerier                    types.String `tfsdk:"failover_querier"`
	FastleaveForNetworkIDs             types.List   `tfsdk:"fastleave_for_network_ids"`
	FloodKnownProtocols                types.Bool   `tfsdk:"flood_known_protocols"`
	FloodUnknownMulticastForNetworkIDs types.List   `tfsdk:"flood_unknown_multicast_for_network_ids"`
	ForwardUnknownMcastRouterPorts     types.Bool   `tfsdk:"forward_unknown_mcast_router_ports"`
	NetworkIDs                         types.List   `tfsdk:"network_ids"`
	PrimaryQuerier                     types.String `tfsdk:"primary_querier"`
	QuerierAddresses                   types.List   `tfsdk:"querier_addresses"`
	QuerierMode                        types.String `tfsdk:"querier_mode"`
	QuerierSubscriptionMode            types.String `tfsdk:"querier_subscription_mode"`
	QuerierSwitches                    types.List   `tfsdk:"querier_switches"`
	SubscriptionMode                   types.String `tfsdk:"subscription_mode"`
	Switches                           types.List   `tfsdk:"switches"`
}

func (r *settingIgmpSnoopingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_setting_igmp_snooping"
}

func (r *settingIgmpSnoopingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Multicast (IGMP) snooping config: querier mode, subscription mode, flood behaviour.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"site": schema.StringAttribute{
				Optional: true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace(), stringplanmodifier.UseStateForUnknown()},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"failover_querier": schema.StringAttribute{
				MarkdownDescription: "failover_querier field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"fastleave_for_network_ids": schema.ListAttribute{
				MarkdownDescription: "fastleave_for_network_ids field",
				Optional:            true, Computed: true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"flood_known_protocols": schema.BoolAttribute{
				MarkdownDescription: "flood_known_protocols field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"flood_unknown_multicast_for_network_ids": schema.ListAttribute{
				MarkdownDescription: "flood_unknown_multicast_for_network_ids field",
				Optional:            true, Computed: true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"forward_unknown_mcast_router_ports": schema.BoolAttribute{
				MarkdownDescription: "forward_unknown_mcast_router_ports field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: "network_ids field",
				Optional:            true, Computed: true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"primary_querier": schema.StringAttribute{
				MarkdownDescription: "primary_querier field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"querier_addresses": schema.ListAttribute{
				MarkdownDescription: "querier_addresses field",
				Optional:            true, Computed: true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"querier_mode": schema.StringAttribute{
				MarkdownDescription: "OFF|ON|AUTO",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"querier_subscription_mode": schema.StringAttribute{
				MarkdownDescription: "ALL|CUSTOM",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"querier_switches": schema.ListAttribute{
				MarkdownDescription: "querier_switches field",
				Optional:            true, Computed: true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"subscription_mode": schema.StringAttribute{
				MarkdownDescription: "ALL|CUSTOM",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"switches": schema.ListAttribute{
				MarkdownDescription: "switches field",
				Optional:            true, Computed: true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *settingIgmpSnoopingResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *settingIgmpSnoopingResource) siteOrDefault(site types.String) string {
	if s := site.ValueString(); s != "" {
		return s
	}
	return r.client.Site
}

func (r *settingIgmpSnoopingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan settingIgmpSnoopingModel
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

func (r *settingIgmpSnoopingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state settingIgmpSnoopingModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	site := r.siteOrDefault(state.Site)

	meta, current, err := ui.GetSetting[*settings.IgmpSnooping](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading IgmpSnooping Setting", err.Error())
		return
	}
	resp.Diagnostics.Append(r.settingToModel(ctx, meta, current, &state, site)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *settingIgmpSnoopingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state settingIgmpSnoopingModel
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

func (r *settingIgmpSnoopingResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Settings cannot be deleted from the controller; dropping from state only.
}

func (r *settingIgmpSnoopingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *settingIgmpSnoopingResource) writeAndRefresh(ctx context.Context, site string, m *settingIgmpSnoopingModel) diag.Diagnostics {
	var diags diag.Diagnostics

	// Read current so unmodified fields aren't dropped to zero values.
	_, current, err := ui.GetSetting[*settings.IgmpSnooping](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); !ok {
			diags.AddError("Error Reading IgmpSnooping Setting", err.Error())
			return diags
		}
		current = &settings.IgmpSnooping{}
	}

	r.applyModelToSetting(ctx, m, current, &diags)
	if diags.HasError() {
		return diags
	}

	if err := r.client.UpdateSetting(ctx, site, current); err != nil {
		diags.AddError("Error Updating IgmpSnooping Setting", err.Error())
		return diags
	}

	meta, refreshed, err := ui.GetSetting[*settings.IgmpSnooping](r.client.ApiClient, ctx, site)
	if err != nil {
		diags.AddError("Error Re-reading IgmpSnooping Setting", err.Error())
		return diags
	}
	diags.Append(r.settingToModel(ctx, meta, refreshed, m, site)...)
	return diags
}

func (r *settingIgmpSnoopingResource) settingToModel(ctx context.Context, meta *ui.Setting, s *settings.IgmpSnooping, m *settingIgmpSnoopingModel, site string) diag.Diagnostics {
	var diags diag.Diagnostics
	if meta != nil {
		m.ID = types.StringValue(meta.Id)
	}
	m.Site = types.StringValue(site)
	_ = ctx
	_ = diags
	m.Enabled = types.BoolValue(s.Enabled)
	m.FailoverQuerier = stringOrNull(s.FailoverQuerier)
	m.FastleaveForNetworkIDs = stringListOrNull(ctx, s.FastleaveForNetworkIDs, &diags)
	m.FloodKnownProtocols = types.BoolValue(s.FloodKnownProtocols)
	m.FloodUnknownMulticastForNetworkIDs = stringListOrNull(ctx, s.FloodUnknownMulticastForNetworkIDs, &diags)
	m.ForwardUnknownMcastRouterPorts = types.BoolValue(s.ForwardUnknownMcastRouterPorts)
	m.NetworkIDs = stringListOrNull(ctx, s.NetworkIDs, &diags)
	m.PrimaryQuerier = stringOrNull(s.PrimaryQuerier)
	m.QuerierAddresses = stringListOrNull(ctx, s.QuerierAddresses, &diags)
	m.QuerierMode = stringOrNull(s.QuerierMode)
	m.QuerierSubscriptionMode = stringOrNull(s.QuerierSubscriptionMode)
	m.QuerierSwitches = stringListOrNull(ctx, s.QuerierSwitches, &diags)
	m.SubscriptionMode = stringOrNull(s.SubscriptionMode)
	m.Switches = stringListOrNull(ctx, s.Switches, &diags)
	return diags
}

func (r *settingIgmpSnoopingResource) applyModelToSetting(ctx context.Context, m *settingIgmpSnoopingModel, s *settings.IgmpSnooping, diags *diag.Diagnostics) {
	_ = ctx
	_ = diags
	s.Enabled = m.Enabled.ValueBool()
	s.FailoverQuerier = m.FailoverQuerier.ValueString()
	if !m.FastleaveForNetworkIDs.IsNull() && !m.FastleaveForNetworkIDs.IsUnknown() {
		var v []string
		diags.Append(m.FastleaveForNetworkIDs.ElementsAs(ctx, &v, false)...)
		s.FastleaveForNetworkIDs = v
	}
	s.FloodKnownProtocols = m.FloodKnownProtocols.ValueBool()
	if !m.FloodUnknownMulticastForNetworkIDs.IsNull() && !m.FloodUnknownMulticastForNetworkIDs.IsUnknown() {
		var v []string
		diags.Append(m.FloodUnknownMulticastForNetworkIDs.ElementsAs(ctx, &v, false)...)
		s.FloodUnknownMulticastForNetworkIDs = v
	}
	s.ForwardUnknownMcastRouterPorts = m.ForwardUnknownMcastRouterPorts.ValueBool()
	if !m.NetworkIDs.IsNull() && !m.NetworkIDs.IsUnknown() {
		var v []string
		diags.Append(m.NetworkIDs.ElementsAs(ctx, &v, false)...)
		s.NetworkIDs = v
	}
	s.PrimaryQuerier = m.PrimaryQuerier.ValueString()
	if !m.QuerierAddresses.IsNull() && !m.QuerierAddresses.IsUnknown() {
		var v []string
		diags.Append(m.QuerierAddresses.ElementsAs(ctx, &v, false)...)
		s.QuerierAddresses = v
	}
	s.QuerierMode = m.QuerierMode.ValueString()
	s.QuerierSubscriptionMode = m.QuerierSubscriptionMode.ValueString()
	if !m.QuerierSwitches.IsNull() && !m.QuerierSwitches.IsUnknown() {
		var v []string
		diags.Append(m.QuerierSwitches.ElementsAs(ctx, &v, false)...)
		s.QuerierSwitches = v
	}
	s.SubscriptionMode = m.SubscriptionMode.ValueString()
	if !m.Switches.IsNull() && !m.Switches.IsUnknown() {
		var v []string
		diags.Append(m.Switches.ElementsAs(ctx, &v, false)...)
		s.Switches = v
	}
}
