package unifi

// Generated from go-unifi/unifi/settings/traffic_flow.generated.go.
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
	_ resource.Resource                = &settingTrafficFlowResource{}
	_ resource.ResourceWithImportState = &settingTrafficFlowResource{}
)

func NewSettingTrafficFlowResource() resource.Resource {
	return &settingTrafficFlowResource{}
}

type settingTrafficFlowResource struct {
	client *Client
}

type settingTrafficFlowModel struct {
	ID                           types.String `tfsdk:"id"`
	Site                         types.String `tfsdk:"site"`
	EnabledAllowedTraffic        types.Bool   `tfsdk:"enabled_allowed_traffic"`
	GatewayDNSEnabled            types.Bool   `tfsdk:"gateway_dns_enabled"`
	UnifiDeviceManagementEnabled types.Bool   `tfsdk:"unifi_device_management_enabled"`
	UnifiServicesEnabled         types.Bool   `tfsdk:"unifi_services_enabled"`
}

func (r *settingTrafficFlowResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_setting_traffic_flow"
}

func (r *settingTrafficFlowResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Traffic flow capture toggles.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"site": schema.StringAttribute{
				Optional: true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace(), stringplanmodifier.UseStateForUnknown()},
			},
			"enabled_allowed_traffic": schema.BoolAttribute{
				MarkdownDescription: "enabled_allowed_traffic field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"gateway_dns_enabled": schema.BoolAttribute{
				MarkdownDescription: "gateway_dns_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"unifi_device_management_enabled": schema.BoolAttribute{
				MarkdownDescription: "unifi_device_management_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"unifi_services_enabled": schema.BoolAttribute{
				MarkdownDescription: "unifi_services_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *settingTrafficFlowResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *settingTrafficFlowResource) siteOrDefault(site types.String) string {
	if s := site.ValueString(); s != "" {
		return s
	}
	return r.client.Site
}

func (r *settingTrafficFlowResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan settingTrafficFlowModel
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

func (r *settingTrafficFlowResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state settingTrafficFlowModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	site := r.siteOrDefault(state.Site)

	meta, current, err := ui.GetSetting[*settings.TrafficFlow](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading TrafficFlow Setting", err.Error())
		return
	}
	resp.Diagnostics.Append(r.settingToModel(ctx, meta, current, &state, site)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *settingTrafficFlowResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state settingTrafficFlowModel
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

func (r *settingTrafficFlowResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Settings cannot be deleted from the controller; dropping from state only.
}

func (r *settingTrafficFlowResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *settingTrafficFlowResource) writeAndRefresh(ctx context.Context, site string, m *settingTrafficFlowModel) diag.Diagnostics {
	var diags diag.Diagnostics

	// Read current so unmodified fields aren't dropped to zero values.
	_, current, err := ui.GetSetting[*settings.TrafficFlow](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); !ok {
			diags.AddError("Error Reading TrafficFlow Setting", err.Error())
			return diags
		}
		current = &settings.TrafficFlow{}
	}

	r.applyModelToSetting(ctx, m, current, &diags)
	if diags.HasError() {
		return diags
	}

	if err := r.client.UpdateSetting(ctx, site, current); err != nil {
		diags.AddError("Error Updating TrafficFlow Setting", err.Error())
		return diags
	}

	meta, refreshed, err := ui.GetSetting[*settings.TrafficFlow](r.client.ApiClient, ctx, site)
	if err != nil {
		diags.AddError("Error Re-reading TrafficFlow Setting", err.Error())
		return diags
	}
	diags.Append(r.settingToModel(ctx, meta, refreshed, m, site)...)
	return diags
}

func (r *settingTrafficFlowResource) settingToModel(ctx context.Context, meta *ui.Setting, s *settings.TrafficFlow, m *settingTrafficFlowModel, site string) diag.Diagnostics {
	var diags diag.Diagnostics
	if meta != nil {
		m.ID = types.StringValue(meta.Id)
	}
	m.Site = types.StringValue(site)
	_ = ctx
	_ = diags
	m.EnabledAllowedTraffic = types.BoolValue(s.EnabledAllowedTraffic)
	m.GatewayDNSEnabled = types.BoolValue(s.GatewayDNSEnabled)
	m.UnifiDeviceManagementEnabled = types.BoolValue(s.UnifiDeviceManagementEnabled)
	m.UnifiServicesEnabled = types.BoolValue(s.UnifiServicesEnabled)
	return diags
}

func (r *settingTrafficFlowResource) applyModelToSetting(ctx context.Context, m *settingTrafficFlowModel, s *settings.TrafficFlow, diags *diag.Diagnostics) {
	_ = ctx
	_ = diags
	s.EnabledAllowedTraffic = m.EnabledAllowedTraffic.ValueBool()
	s.GatewayDNSEnabled = m.GatewayDNSEnabled.ValueBool()
	s.UnifiDeviceManagementEnabled = m.UnifiDeviceManagementEnabled.ValueBool()
	s.UnifiServicesEnabled = m.UnifiServicesEnabled.ValueBool()
}
