package unifi

// Generated from go-unifi/unifi/settings/roaming_assistant.generated.go.
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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ui "github.com/ubiquiti-community/go-unifi/unifi"
	"github.com/ubiquiti-community/go-unifi/unifi/settings"
)

var (
	_ resource.Resource                = &settingRoamingAssistantResource{}
	_ resource.ResourceWithImportState = &settingRoamingAssistantResource{}
)

func NewSettingRoamingAssistantResource() resource.Resource {
	return &settingRoamingAssistantResource{}
}

type settingRoamingAssistantResource struct {
	client *Client
}

type settingRoamingAssistantModel struct {
	ID      types.String `tfsdk:"id"`
	Site    types.String `tfsdk:"site"`
	Enabled types.Bool   `tfsdk:"enabled"`
	Rssi    types.Int64  `tfsdk:"rssi"`
}

func (r *settingRoamingAssistantResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_setting_roaming_assistant"
}

func (r *settingRoamingAssistantResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "WiFi roaming assistant RSSI threshold (sticky-client mitigation).",
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
			"rssi": schema.Int64Attribute{
				MarkdownDescription: "^-([6-7][0-9]|80)$",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *settingRoamingAssistantResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *settingRoamingAssistantResource) siteOrDefault(site types.String) string {
	if s := site.ValueString(); s != "" {
		return s
	}
	return r.client.Site
}

func (r *settingRoamingAssistantResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan settingRoamingAssistantModel
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

func (r *settingRoamingAssistantResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state settingRoamingAssistantModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	site := r.siteOrDefault(state.Site)

	meta, current, err := ui.GetSetting[*settings.RoamingAssistant](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading RoamingAssistant Setting", err.Error())
		return
	}
	resp.Diagnostics.Append(r.settingToModel(ctx, meta, current, &state, site)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *settingRoamingAssistantResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state settingRoamingAssistantModel
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

func (r *settingRoamingAssistantResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Settings cannot be deleted from the controller; dropping from state only.
}

func (r *settingRoamingAssistantResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *settingRoamingAssistantResource) writeAndRefresh(ctx context.Context, site string, m *settingRoamingAssistantModel) diag.Diagnostics {
	var diags diag.Diagnostics

	// Read current so unmodified fields aren't dropped to zero values.
	_, current, err := ui.GetSetting[*settings.RoamingAssistant](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); !ok {
			diags.AddError("Error Reading RoamingAssistant Setting", err.Error())
			return diags
		}
		current = &settings.RoamingAssistant{}
	}

	r.applyModelToSetting(ctx, m, current, &diags)
	if diags.HasError() {
		return diags
	}

	if err := r.client.UpdateSetting(ctx, site, current); err != nil {
		diags.AddError("Error Updating RoamingAssistant Setting", err.Error())
		return diags
	}

	meta, refreshed, err := ui.GetSetting[*settings.RoamingAssistant](r.client.ApiClient, ctx, site)
	if err != nil {
		diags.AddError("Error Re-reading RoamingAssistant Setting", err.Error())
		return diags
	}
	diags.Append(r.settingToModel(ctx, meta, refreshed, m, site)...)
	return diags
}

func (r *settingRoamingAssistantResource) settingToModel(ctx context.Context, meta *ui.Setting, s *settings.RoamingAssistant, m *settingRoamingAssistantModel, site string) diag.Diagnostics {
	var diags diag.Diagnostics
	if meta != nil {
		m.ID = types.StringValue(meta.Id)
	}
	m.Site = types.StringValue(site)
	_ = ctx
	_ = diags
	m.Enabled = types.BoolValue(s.Enabled)
	m.Rssi = types.Int64PointerValue(s.Rssi)
	return diags
}

func (r *settingRoamingAssistantResource) applyModelToSetting(ctx context.Context, m *settingRoamingAssistantModel, s *settings.RoamingAssistant, diags *diag.Diagnostics) {
	_ = ctx
	_ = diags
	if !m.Enabled.IsNull() && !m.Enabled.IsUnknown() {
		s.Enabled = m.Enabled.ValueBool()
	}
	if !m.Rssi.IsNull() && !m.Rssi.IsUnknown() {
		s.Rssi = int64PointerOrNil(m.Rssi)
	}
}
