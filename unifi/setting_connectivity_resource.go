package unifi

// Generated from go-unifi/unifi/settings/connectivity.generated.go.
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
	_ resource.Resource                = &settingConnectivityResource{}
	_ resource.ResourceWithImportState = &settingConnectivityResource{}
)

func NewSettingConnectivityResource() resource.Resource {
	return &settingConnectivityResource{}
}

type settingConnectivityResource struct {
	client *Client
}

type settingConnectivityModel struct {
	ID                 types.String `tfsdk:"id"`
	Site               types.String `tfsdk:"site"`
	EnableIsolatedWLAN types.Bool   `tfsdk:"enable_isolated_wlan"`
	Enabled            types.Bool   `tfsdk:"enabled"`
	MeshEssid          types.String `tfsdk:"x_mesh_essid"`
	MeshPsk            types.String `tfsdk:"x_mesh_psk"`
	UplinkHost         types.String `tfsdk:"uplink_host"`
	UplinkType         types.String `tfsdk:"uplink_type"`
}

func (r *settingConnectivityResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_setting_connectivity"
}

func (r *settingConnectivityResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "WAN uplink connectivity monitor (used by failover logic).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"site": schema.StringAttribute{
				Optional: true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace(), stringplanmodifier.UseStateForUnknown()},
			},
			"enable_isolated_wlan": schema.BoolAttribute{
				MarkdownDescription: "enable_isolated_wlan field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"x_mesh_essid": schema.StringAttribute{
				MarkdownDescription: "x_mesh_essid field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"x_mesh_psk": schema.StringAttribute{
				MarkdownDescription: "x_mesh_psk field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"uplink_host": schema.StringAttribute{
				MarkdownDescription: "uplink_host field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"uplink_type": schema.StringAttribute{
				MarkdownDescription: "uplink_type field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *settingConnectivityResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *settingConnectivityResource) siteOrDefault(site types.String) string {
	if s := site.ValueString(); s != "" {
		return s
	}
	return r.client.Site
}

func (r *settingConnectivityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan settingConnectivityModel
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

func (r *settingConnectivityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state settingConnectivityModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	site := r.siteOrDefault(state.Site)

	meta, current, err := ui.GetSetting[*settings.Connectivity](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading Connectivity Setting", err.Error())
		return
	}
	resp.Diagnostics.Append(r.settingToModel(ctx, meta, current, &state, site)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *settingConnectivityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state settingConnectivityModel
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

func (r *settingConnectivityResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Settings cannot be deleted from the controller; dropping from state only.
}

func (r *settingConnectivityResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *settingConnectivityResource) writeAndRefresh(ctx context.Context, site string, m *settingConnectivityModel) diag.Diagnostics {
	var diags diag.Diagnostics

	// Read current so unmodified fields aren't dropped to zero values.
	_, current, err := ui.GetSetting[*settings.Connectivity](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); !ok {
			diags.AddError("Error Reading Connectivity Setting", err.Error())
			return diags
		}
		current = &settings.Connectivity{}
	}

	r.applyModelToSetting(ctx, m, current, &diags)
	if diags.HasError() {
		return diags
	}

	if err := r.client.UpdateSetting(ctx, site, current); err != nil {
		diags.AddError("Error Updating Connectivity Setting", err.Error())
		return diags
	}

	meta, refreshed, err := ui.GetSetting[*settings.Connectivity](r.client.ApiClient, ctx, site)
	if err != nil {
		diags.AddError("Error Re-reading Connectivity Setting", err.Error())
		return diags
	}
	diags.Append(r.settingToModel(ctx, meta, refreshed, m, site)...)
	return diags
}

func (r *settingConnectivityResource) settingToModel(ctx context.Context, meta *ui.Setting, s *settings.Connectivity, m *settingConnectivityModel, site string) diag.Diagnostics {
	var diags diag.Diagnostics
	if meta != nil {
		m.ID = types.StringValue(meta.Id)
	}
	m.Site = types.StringValue(site)
	_ = ctx
	_ = diags
	m.EnableIsolatedWLAN = types.BoolValue(s.EnableIsolatedWLAN)
	m.Enabled = types.BoolValue(s.Enabled)
	m.MeshEssid = stringOrNull(s.MeshEssid)
	m.MeshPsk = stringOrNull(s.MeshPsk)
	m.UplinkHost = stringOrNull(s.UplinkHost)
	m.UplinkType = stringOrNull(s.UplinkType)
	return diags
}

func (r *settingConnectivityResource) applyModelToSetting(ctx context.Context, m *settingConnectivityModel, s *settings.Connectivity, diags *diag.Diagnostics) {
	_ = ctx
	_ = diags
	s.EnableIsolatedWLAN = m.EnableIsolatedWLAN.ValueBool()
	s.Enabled = m.Enabled.ValueBool()
	s.MeshEssid = m.MeshEssid.ValueString()
	s.MeshPsk = m.MeshPsk.ValueString()
	s.UplinkHost = m.UplinkHost.ValueString()
	s.UplinkType = m.UplinkType.ValueString()
}
