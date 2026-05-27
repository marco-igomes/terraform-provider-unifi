package unifi

// Generated from go-unifi/unifi/settings/rsyslogd.generated.go.
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
	_ resource.Resource                = &settingRsyslogdResource{}
	_ resource.ResourceWithImportState = &settingRsyslogdResource{}
)

func NewSettingRsyslogdResource() resource.Resource {
	return &settingRsyslogdResource{}
}

type settingRsyslogdResource struct {
	client *Client
}

type settingRsyslogdModel struct {
	ID                          types.String `tfsdk:"id"`
	Site                        types.String `tfsdk:"site"`
	Contents                    types.List   `tfsdk:"contents"`
	Debug                       types.Bool   `tfsdk:"debug"`
	Enabled                     types.Bool   `tfsdk:"enabled"`
	IP                          types.String `tfsdk:"ip"`
	LogAllContents              types.Bool   `tfsdk:"log_all_contents"`
	NetconsoleEnabled           types.Bool   `tfsdk:"netconsole_enabled"`
	NetconsoleHost              types.String `tfsdk:"netconsole_host"`
	NetconsolePort              types.Int64  `tfsdk:"netconsole_port"`
	Port                        types.Int64  `tfsdk:"port"`
	ThisController              types.Bool   `tfsdk:"this_controller"`
	ThisControllerEncryptedOnly types.Bool   `tfsdk:"this_controller_encrypted_only"`
}

func (r *settingRsyslogdResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_setting_rsyslogd"
}

func (r *settingRsyslogdResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Remote syslog (rsyslogd) configuration.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"site": schema.StringAttribute{
				Optional: true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace(), stringplanmodifier.UseStateForUnknown()},
			},
			"contents": schema.ListAttribute{
				MarkdownDescription: "device|client|firewall_default_policy|triggers|updates|admin_activity|critical|security_detections|vpn",
				Optional:            true, Computed: true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"debug": schema.BoolAttribute{
				MarkdownDescription: "debug field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"ip": schema.StringAttribute{
				MarkdownDescription: "ip field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"log_all_contents": schema.BoolAttribute{
				MarkdownDescription: "log_all_contents field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"netconsole_enabled": schema.BoolAttribute{
				MarkdownDescription: "netconsole_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"netconsole_host": schema.StringAttribute{
				MarkdownDescription: "netconsole_host field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"netconsole_port": schema.Int64Attribute{
				MarkdownDescription: "[1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5]",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: "[1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5]",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"this_controller": schema.BoolAttribute{
				MarkdownDescription: "this_controller field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"this_controller_encrypted_only": schema.BoolAttribute{
				MarkdownDescription: "this_controller_encrypted_only field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *settingRsyslogdResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *settingRsyslogdResource) siteOrDefault(site types.String) string {
	if s := site.ValueString(); s != "" {
		return s
	}
	return r.client.Site
}

func (r *settingRsyslogdResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan settingRsyslogdModel
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

func (r *settingRsyslogdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state settingRsyslogdModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	site := r.siteOrDefault(state.Site)

	meta, current, err := ui.GetSetting[*settings.Rsyslogd](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading Rsyslogd Setting", err.Error())
		return
	}
	resp.Diagnostics.Append(r.settingToModel(ctx, meta, current, &state, site)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *settingRsyslogdResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state settingRsyslogdModel
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

func (r *settingRsyslogdResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Settings cannot be deleted from the controller; dropping from state only.
}

func (r *settingRsyslogdResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *settingRsyslogdResource) writeAndRefresh(ctx context.Context, site string, m *settingRsyslogdModel) diag.Diagnostics {
	var diags diag.Diagnostics

	// Read current so unmodified fields aren't dropped to zero values.
	_, current, err := ui.GetSetting[*settings.Rsyslogd](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); !ok {
			diags.AddError("Error Reading Rsyslogd Setting", err.Error())
			return diags
		}
		current = &settings.Rsyslogd{}
	}

	r.applyModelToSetting(ctx, m, current, &diags)
	if diags.HasError() {
		return diags
	}

	if err := r.client.UpdateSetting(ctx, site, current); err != nil {
		diags.AddError("Error Updating Rsyslogd Setting", err.Error())
		return diags
	}

	meta, refreshed, err := ui.GetSetting[*settings.Rsyslogd](r.client.ApiClient, ctx, site)
	if err != nil {
		diags.AddError("Error Re-reading Rsyslogd Setting", err.Error())
		return diags
	}
	diags.Append(r.settingToModel(ctx, meta, refreshed, m, site)...)
	return diags
}

func (r *settingRsyslogdResource) settingToModel(ctx context.Context, meta *ui.Setting, s *settings.Rsyslogd, m *settingRsyslogdModel, site string) diag.Diagnostics {
	var diags diag.Diagnostics
	if meta != nil {
		m.ID = types.StringValue(meta.Id)
	}
	m.Site = types.StringValue(site)
	_ = ctx
	_ = diags
	m.Contents = stringListOrNull(ctx, s.Contents, &diags)
	m.Debug = types.BoolValue(s.Debug)
	m.Enabled = types.BoolValue(s.Enabled)
	m.IP = stringOrNull(s.IP)
	m.LogAllContents = types.BoolValue(s.LogAllContents)
	m.NetconsoleEnabled = types.BoolValue(s.NetconsoleEnabled)
	m.NetconsoleHost = stringOrNull(s.NetconsoleHost)
	m.NetconsolePort = types.Int64PointerValue(s.NetconsolePort)
	m.Port = types.Int64PointerValue(s.Port)
	m.ThisController = types.BoolValue(s.ThisController)
	m.ThisControllerEncryptedOnly = types.BoolValue(s.ThisControllerEncryptedOnly)
	return diags
}

func (r *settingRsyslogdResource) applyModelToSetting(ctx context.Context, m *settingRsyslogdModel, s *settings.Rsyslogd, diags *diag.Diagnostics) {
	_ = ctx
	_ = diags
	if !m.Contents.IsNull() && !m.Contents.IsUnknown() {
		var v []string
		diags.Append(m.Contents.ElementsAs(ctx, &v, false)...)
		s.Contents = v
	}
	s.Debug = m.Debug.ValueBool()
	s.Enabled = m.Enabled.ValueBool()
	s.IP = m.IP.ValueString()
	s.LogAllContents = m.LogAllContents.ValueBool()
	s.NetconsoleEnabled = m.NetconsoleEnabled.ValueBool()
	s.NetconsoleHost = m.NetconsoleHost.ValueString()
	s.NetconsolePort = int64PointerOrNil(m.NetconsolePort)
	s.Port = int64PointerOrNil(m.Port)
	s.ThisController = m.ThisController.ValueBool()
	s.ThisControllerEncryptedOnly = m.ThisControllerEncryptedOnly.ValueBool()
}
