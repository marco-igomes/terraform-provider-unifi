package unifi

// Generated from go-unifi/unifi/settings/netflow.generated.go.
// DO NOT EDIT MANUALLY — regenerate via scripts/gen_setting_resources.py.

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ui "github.com/ubiquiti-community/go-unifi/unifi"
	"github.com/ubiquiti-community/go-unifi/unifi/settings"
)

var (
	_ resource.Resource                = &settingNetflowResource{}
	_ resource.ResourceWithImportState = &settingNetflowResource{}
)

func NewSettingNetflowResource() resource.Resource {
	return &settingNetflowResource{}
}

type settingNetflowResource struct {
	client *Client
}

type settingNetflowModel struct {
	ID                  types.String `tfsdk:"id"`
	Site                types.String `tfsdk:"site"`
	AutoEngineIDEnabled types.Bool   `tfsdk:"auto_engine_id_enabled"`
	Enabled             types.Bool   `tfsdk:"enabled"`
	EngineID            types.Int64  `tfsdk:"engine_id"`
	ExportFrequency     types.Int64  `tfsdk:"export_frequency"`
	NetworkIDs          types.List   `tfsdk:"network_ids"`
	Port                types.Int64  `tfsdk:"port"`
	RefreshRate         types.Int64  `tfsdk:"refresh_rate"`
	SamplingMode        types.String `tfsdk:"sampling_mode"`
	SamplingRate        types.Int64  `tfsdk:"sampling_rate"`
	Server              types.String `tfsdk:"server"`
	Version             types.Int64  `tfsdk:"version"`
}

func (r *settingNetflowResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_setting_netflow"
}

func (r *settingNetflowResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "NetFlow/sFlow flow exporter configuration.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"site": schema.StringAttribute{
				Optional: true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace(), stringplanmodifier.UseStateForUnknown()},
			},
			"auto_engine_id_enabled": schema.BoolAttribute{
				MarkdownDescription: "auto_engine_id_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"engine_id": schema.Int64Attribute{
				MarkdownDescription: "^$|[1-9][0-9]*",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"export_frequency": schema.Int64Attribute{
				MarkdownDescription: "export_frequency field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: "network_ids field",
				Optional:            true, Computed: true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: "102[4-9]|10[3-9][0-9]|1[1-9][0-9]{2}|[2-9][0-9]{3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5]",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"refresh_rate": schema.Int64Attribute{
				MarkdownDescription: "refresh_rate field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"sampling_mode": schema.StringAttribute{
				MarkdownDescription: "off|hash|random|deterministic",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"sampling_rate": schema.Int64Attribute{
				MarkdownDescription: "[2-9]|[1-9][0-9]{1,3}|1[0-5][0-9]{3}|16[0-2][0-9]{2}|163[0-7][0-9]|1638[0-3]|^$",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"server": schema.StringAttribute{
				MarkdownDescription: ".{0,252}[^\\.]$",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "5|9|10",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *settingNetflowResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *settingNetflowResource) siteOrDefault(site types.String) string {
	if s := site.ValueString(); s != "" {
		return s
	}
	return r.client.Site
}

func (r *settingNetflowResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan settingNetflowModel
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

func (r *settingNetflowResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state settingNetflowModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	site := r.siteOrDefault(state.Site)

	meta, current, err := ui.GetSetting[*settings.Netflow](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading Netflow Setting", err.Error())
		return
	}
	resp.Diagnostics.Append(r.settingToModel(ctx, meta, current, &state, site)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *settingNetflowResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state settingNetflowModel
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

func (r *settingNetflowResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Settings cannot be deleted from the controller; dropping from state only.
}

func (r *settingNetflowResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *settingNetflowResource) writeAndRefresh(ctx context.Context, site string, m *settingNetflowModel) diag.Diagnostics {
	var diags diag.Diagnostics

	// Read current so unmodified fields aren't dropped to zero values.
	_, current, err := ui.GetSetting[*settings.Netflow](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); !ok {
			diags.AddError("Error Reading Netflow Setting", err.Error())
			return diags
		}
		current = &settings.Netflow{}
	}

	r.applyModelToSetting(ctx, m, current, &diags)
	if diags.HasError() {
		return diags
	}

	if err := r.client.UpdateSetting(ctx, site, current); err != nil {
		diags.AddError("Error Updating Netflow Setting", err.Error())
		return diags
	}

	meta, refreshed, err := ui.GetSetting[*settings.Netflow](r.client.ApiClient, ctx, site)
	if err != nil {
		diags.AddError("Error Re-reading Netflow Setting", err.Error())
		return diags
	}
	diags.Append(r.settingToModel(ctx, meta, refreshed, m, site)...)
	return diags
}

func (r *settingNetflowResource) settingToModel(ctx context.Context, meta *ui.Setting, s *settings.Netflow, m *settingNetflowModel, site string) diag.Diagnostics {
	var diags diag.Diagnostics
	if meta != nil {
		m.ID = types.StringValue(meta.Id)
	}
	m.Site = types.StringValue(site)
	_ = ctx
	_ = diags
	m.AutoEngineIDEnabled = types.BoolValue(s.AutoEngineIDEnabled)
	m.Enabled = types.BoolValue(s.Enabled)
	m.EngineID = types.Int64PointerValue(s.EngineID)
	m.ExportFrequency = types.Int64PointerValue(s.ExportFrequency)
	m.NetworkIDs = stringListOrNull(ctx, s.NetworkIDs, &diags)
	m.Port = types.Int64PointerValue(s.Port)
	m.RefreshRate = types.Int64PointerValue(s.RefreshRate)
	m.SamplingMode = stringOrNull(s.SamplingMode)
	m.SamplingRate = types.Int64PointerValue(s.SamplingRate)
	m.Server = stringOrNull(s.Server)
	m.Version = types.Int64PointerValue(s.Version)
	return diags
}

func (r *settingNetflowResource) applyModelToSetting(ctx context.Context, m *settingNetflowModel, s *settings.Netflow, diags *diag.Diagnostics) {
	_ = ctx
	_ = diags
	if !m.AutoEngineIDEnabled.IsNull() && !m.AutoEngineIDEnabled.IsUnknown() {
		s.AutoEngineIDEnabled = m.AutoEngineIDEnabled.ValueBool()
	}
	if !m.Enabled.IsNull() && !m.Enabled.IsUnknown() {
		s.Enabled = m.Enabled.ValueBool()
	}
	if !m.EngineID.IsNull() && !m.EngineID.IsUnknown() {
		s.EngineID = int64PointerOrNil(m.EngineID)
	}
	if !m.ExportFrequency.IsNull() && !m.ExportFrequency.IsUnknown() {
		s.ExportFrequency = int64PointerOrNil(m.ExportFrequency)
	}
	if !m.NetworkIDs.IsNull() && !m.NetworkIDs.IsUnknown() {
		var v []string
		diags.Append(m.NetworkIDs.ElementsAs(ctx, &v, false)...)
		s.NetworkIDs = v
	}
	if !m.Port.IsNull() && !m.Port.IsUnknown() {
		s.Port = int64PointerOrNil(m.Port)
	}
	if !m.RefreshRate.IsNull() && !m.RefreshRate.IsUnknown() {
		s.RefreshRate = int64PointerOrNil(m.RefreshRate)
	}
	if !m.SamplingMode.IsNull() && !m.SamplingMode.IsUnknown() {
		s.SamplingMode = m.SamplingMode.ValueString()
	}
	if !m.SamplingRate.IsNull() && !m.SamplingRate.IsUnknown() {
		s.SamplingRate = int64PointerOrNil(m.SamplingRate)
	}
	if !m.Server.IsNull() && !m.Server.IsUnknown() {
		s.Server = m.Server.ValueString()
	}
	if !m.Version.IsNull() && !m.Version.IsUnknown() {
		s.Version = int64PointerOrNil(m.Version)
	}
}
