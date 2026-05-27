package unifi

// Generated from go-unifi/unifi/settings/locale.generated.go.
// DO NOT EDIT MANUALLY — regenerate via scripts/gen_setting_resources.py.

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ui "github.com/ubiquiti-community/go-unifi/unifi"
	"github.com/ubiquiti-community/go-unifi/unifi/settings"
)

var (
	_ resource.Resource                = &settingLocaleResource{}
	_ resource.ResourceWithImportState = &settingLocaleResource{}
)

func NewSettingLocaleResource() resource.Resource {
	return &settingLocaleResource{}
}

type settingLocaleResource struct {
	client *Client
}

type settingLocaleModel struct {
	ID       types.String `tfsdk:"id"`
	Site     types.String `tfsdk:"site"`
	Timezone types.String `tfsdk:"timezone"`
}

func (r *settingLocaleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_setting_locale"
}

func (r *settingLocaleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Controller timezone.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"site": schema.StringAttribute{
				Optional: true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace(), stringplanmodifier.UseStateForUnknown()},
			},
			"timezone": schema.StringAttribute{
				MarkdownDescription: "timezone field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *settingLocaleResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *settingLocaleResource) siteOrDefault(site types.String) string {
	if s := site.ValueString(); s != "" {
		return s
	}
	return r.client.Site
}

func (r *settingLocaleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan settingLocaleModel
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

func (r *settingLocaleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state settingLocaleModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	site := r.siteOrDefault(state.Site)

	meta, current, err := ui.GetSetting[*settings.Locale](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading Locale Setting", err.Error())
		return
	}
	resp.Diagnostics.Append(r.settingToModel(ctx, meta, current, &state, site)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *settingLocaleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state settingLocaleModel
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

func (r *settingLocaleResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Settings cannot be deleted from the controller; dropping from state only.
}

func (r *settingLocaleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *settingLocaleResource) writeAndRefresh(ctx context.Context, site string, m *settingLocaleModel) diag.Diagnostics {
	var diags diag.Diagnostics

	// Read current so unmodified fields aren't dropped to zero values.
	_, current, err := ui.GetSetting[*settings.Locale](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); !ok {
			diags.AddError("Error Reading Locale Setting", err.Error())
			return diags
		}
		current = &settings.Locale{}
	}

	r.applyModelToSetting(ctx, m, current, &diags)
	if diags.HasError() {
		return diags
	}

	if err := r.client.UpdateSetting(ctx, site, current); err != nil {
		diags.AddError("Error Updating Locale Setting", err.Error())
		return diags
	}

	meta, refreshed, err := ui.GetSetting[*settings.Locale](r.client.ApiClient, ctx, site)
	if err != nil {
		diags.AddError("Error Re-reading Locale Setting", err.Error())
		return diags
	}
	diags.Append(r.settingToModel(ctx, meta, refreshed, m, site)...)
	return diags
}

func (r *settingLocaleResource) settingToModel(ctx context.Context, meta *ui.Setting, s *settings.Locale, m *settingLocaleModel, site string) diag.Diagnostics {
	var diags diag.Diagnostics
	if meta != nil {
		m.ID = types.StringValue(meta.Id)
	}
	m.Site = types.StringValue(site)
	_ = ctx
	_ = diags
	m.Timezone = stringOrNull(s.Timezone)
	return diags
}

func (r *settingLocaleResource) applyModelToSetting(ctx context.Context, m *settingLocaleModel, s *settings.Locale, diags *diag.Diagnostics) {
	_ = ctx
	_ = diags
	if !m.Timezone.IsNull() && !m.Timezone.IsUnknown() {
		s.Timezone = m.Timezone.ValueString()
	}
}
