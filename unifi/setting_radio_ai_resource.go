package unifi

// Generated from go-unifi/unifi/settings/radio_ai.generated.go.
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
	_ resource.Resource                = &settingRadioAiResource{}
	_ resource.ResourceWithImportState = &settingRadioAiResource{}
)

func NewSettingRadioAiResource() resource.Resource {
	return &settingRadioAiResource{}
}

type settingRadioAiResource struct {
	client *Client
}

type settingRadioAiModel struct {
	ID                          types.String `tfsdk:"id"`
	Site                        types.String `tfsdk:"site"`
	AutoAdjustChannelsToCountry types.Bool   `tfsdk:"auto_adjust_channels_to_country"`
	AutoChannelPresetsType      types.String `tfsdk:"auto_channel_presets_type"`
	Channels6E                  types.String `tfsdk:"channels_6e"`
	ChannelsBlacklist           types.String `tfsdk:"channels_blacklist"`
	ChannelsNa                  types.String `tfsdk:"channels_na"`
	ChannelsNg                  types.String `tfsdk:"channels_ng"`
	CronExpr                    types.String `tfsdk:"cron_expr"`
	Default                     types.Bool   `tfsdk:"default"`
	Enabled                     types.Bool   `tfsdk:"enabled"`
	ExcludeDevices              types.List   `tfsdk:"exclude_devices"`
	HighPriorityDevices         types.List   `tfsdk:"high_priority_devices"`
	HtModesNa                   types.String `tfsdk:"ht_modes_na"`
	HtModesNg                   types.String `tfsdk:"ht_modes_ng"`
	Optimize                    types.List   `tfsdk:"optimize"`
	Radios                      types.List   `tfsdk:"radios"`
	RadiosConfiguration         types.String `tfsdk:"radios_configuration"`
	SettingPreference           types.String `tfsdk:"setting_preference"`
	UseXy                       types.Bool   `tfsdk:"use_xy"`
}

func (r *settingRadioAiResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_setting_radio_ai"
}

func (r *settingRadioAiResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "RF auto-optimization (RF Scanning): auto channel/power selection, channel pools, blacklist.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"site": schema.StringAttribute{
				Optional: true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace(), stringplanmodifier.UseStateForUnknown()},
			},
			"auto_adjust_channels_to_country": schema.BoolAttribute{
				MarkdownDescription: "Constrain auto-selected channels to those permitted by the site's regulatory country.",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"auto_channel_presets_type": schema.StringAttribute{
				MarkdownDescription: "Channel-selection preset. One of: \"maximum_speed\", \"conservative\", \"custom\".",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"channels_6e": schema.StringAttribute{
				MarkdownDescription: "6 GHz channel pool the optimizer may select from (JSON array of channel numbers).",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"channels_blacklist": schema.StringAttribute{
				MarkdownDescription: "Channels excluded from auto-selection (JSON array of {channel, channel_width, radio}).",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"channels_na": schema.StringAttribute{
				MarkdownDescription: "5 GHz channel pool the optimizer may select from (JSON array of channel numbers).",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"channels_ng": schema.StringAttribute{
				MarkdownDescription: "2.4 GHz channel pool the optimizer may select from (JSON array of channel numbers).",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"cron_expr": schema.StringAttribute{
				MarkdownDescription: "Cron expression for the scheduled optimization run.",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"default": schema.BoolAttribute{
				MarkdownDescription: "Whether this is the controller's default RF-optimization profile.",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether RF auto-optimization (RF Scanning) is enabled.",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"exclude_devices": schema.ListAttribute{
				MarkdownDescription: "MAC addresses of access points excluded from optimization.",
				Optional:            true, Computed: true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"high_priority_devices": schema.ListAttribute{
				MarkdownDescription: "MAC addresses of access points prioritized during optimization.",
				Optional:            true, Computed: true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"ht_modes_na": schema.StringAttribute{
				MarkdownDescription: "Allowed 5 GHz channel widths in MHz. One or more of: \"20\", \"40\", \"80\", \"160\".",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ht_modes_ng": schema.StringAttribute{
				MarkdownDescription: "Allowed 2.4 GHz channel widths in MHz. One or more of: \"20\", \"40\".",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"optimize": schema.ListAttribute{
				MarkdownDescription: "What RF auto-optimization adjusts. One or more of: \"channel\", \"power\".",
				Optional:            true, Computed: true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"radios": schema.ListAttribute{
				MarkdownDescription: "Radio bands to optimize. One or more of: \"na\" (5 GHz), \"ng\" (2.4 GHz), \"6e\" (6 GHz).",
				Optional:            true, Computed: true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"radios_configuration": schema.StringAttribute{
				MarkdownDescription: "Per-radio optimization settings (JSON array of {radio, channel_width, dfs}).",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"setting_preference": schema.StringAttribute{
				MarkdownDescription: "Whether these settings are auto-managed or manual. One of: \"auto\", \"manual\".",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"use_xy": schema.BoolAttribute{
				MarkdownDescription: "Use the X/Y placement-based optimizer.",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *settingRadioAiResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *settingRadioAiResource) siteOrDefault(site types.String) string {
	if s := site.ValueString(); s != "" {
		return s
	}
	return r.client.Site
}

func (r *settingRadioAiResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan settingRadioAiModel
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

func (r *settingRadioAiResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state settingRadioAiModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	site := r.siteOrDefault(state.Site)

	meta, current, err := ui.GetSetting[*settings.RadioAi](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading RadioAi Setting", err.Error())
		return
	}
	resp.Diagnostics.Append(r.settingToModel(ctx, meta, current, &state, site)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *settingRadioAiResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state settingRadioAiModel
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

func (r *settingRadioAiResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Settings cannot be deleted from the controller; dropping from state only.
}

func (r *settingRadioAiResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *settingRadioAiResource) writeAndRefresh(ctx context.Context, site string, m *settingRadioAiModel) diag.Diagnostics {
	var diags diag.Diagnostics

	// Read current so unmodified fields aren't dropped to zero values.
	_, current, err := ui.GetSetting[*settings.RadioAi](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); !ok {
			diags.AddError("Error Reading RadioAi Setting", err.Error())
			return diags
		}
		current = &settings.RadioAi{}
	}

	r.applyModelToSetting(ctx, m, current, &diags)
	if diags.HasError() {
		return diags
	}

	if err := r.client.UpdateSetting(ctx, site, current); err != nil {
		diags.AddError("Error Updating RadioAi Setting", err.Error())
		return diags
	}

	meta, refreshed, err := ui.GetSetting[*settings.RadioAi](r.client.ApiClient, ctx, site)
	if err != nil {
		diags.AddError("Error Re-reading RadioAi Setting", err.Error())
		return diags
	}
	diags.Append(r.settingToModel(ctx, meta, refreshed, m, site)...)
	return diags
}

func (r *settingRadioAiResource) settingToModel(ctx context.Context, meta *ui.Setting, s *settings.RadioAi, m *settingRadioAiModel, site string) diag.Diagnostics {
	var diags diag.Diagnostics
	if meta != nil {
		m.ID = types.StringValue(meta.Id)
	}
	m.Site = types.StringValue(site)
	_ = ctx
	_ = diags
	m.AutoAdjustChannelsToCountry = types.BoolValue(s.AutoAdjustChannelsToCountry)
	m.AutoChannelPresetsType = stringOrNull(s.AutoChannelPresetsType)
	m.Channels6E = jsonStringFrom(s.Channels6E)
	m.ChannelsBlacklist = jsonStringFrom(s.ChannelsBlacklist)
	m.ChannelsNa = jsonStringFrom(s.ChannelsNa)
	m.ChannelsNg = jsonStringFrom(s.ChannelsNg)
	m.CronExpr = stringOrNull(s.CronExpr)
	m.Default = types.BoolValue(s.Default)
	m.Enabled = types.BoolValue(s.Enabled)
	m.ExcludeDevices = stringList(ctx, s.ExcludeDevices, &diags)
	m.HighPriorityDevices = stringList(ctx, s.HighPriorityDevices, &diags)
	m.HtModesNa = jsonStringFrom(s.HtModesNa)
	m.HtModesNg = jsonStringFrom(s.HtModesNg)
	m.Optimize = stringList(ctx, s.Optimize, &diags)
	m.Radios = stringList(ctx, s.Radios, &diags)
	m.RadiosConfiguration = jsonStringFrom(s.RadiosConfiguration)
	m.SettingPreference = stringOrNull(s.SettingPreference)
	m.UseXy = types.BoolValue(s.UseXy)
	return diags
}

func (r *settingRadioAiResource) applyModelToSetting(ctx context.Context, m *settingRadioAiModel, s *settings.RadioAi, diags *diag.Diagnostics) {
	_ = ctx
	_ = diags
	if !m.AutoAdjustChannelsToCountry.IsNull() && !m.AutoAdjustChannelsToCountry.IsUnknown() {
		s.AutoAdjustChannelsToCountry = m.AutoAdjustChannelsToCountry.ValueBool()
	}
	if !m.AutoChannelPresetsType.IsNull() && !m.AutoChannelPresetsType.IsUnknown() {
		s.AutoChannelPresetsType = m.AutoChannelPresetsType.ValueString()
	}
	if !m.Channels6E.IsNull() && !m.Channels6E.IsUnknown() {
		var val []int64
		if err := json.Unmarshal([]byte(m.Channels6E.ValueString()), &val); err != nil {
			diags.AddError("Invalid channels_6e", err.Error())
		} else {
			s.Channels6E = val
		}
	}
	if !m.ChannelsBlacklist.IsNull() && !m.ChannelsBlacklist.IsUnknown() {
		var val []settings.SettingRadioAiChannelsBlacklist
		if err := json.Unmarshal([]byte(m.ChannelsBlacklist.ValueString()), &val); err != nil {
			diags.AddError("Invalid channels_blacklist", err.Error())
		} else {
			s.ChannelsBlacklist = val
		}
	}
	if !m.ChannelsNa.IsNull() && !m.ChannelsNa.IsUnknown() {
		var val []int64
		if err := json.Unmarshal([]byte(m.ChannelsNa.ValueString()), &val); err != nil {
			diags.AddError("Invalid channels_na", err.Error())
		} else {
			s.ChannelsNa = val
		}
	}
	if !m.ChannelsNg.IsNull() && !m.ChannelsNg.IsUnknown() {
		var val []int64
		if err := json.Unmarshal([]byte(m.ChannelsNg.ValueString()), &val); err != nil {
			diags.AddError("Invalid channels_ng", err.Error())
		} else {
			s.ChannelsNg = val
		}
	}
	if !m.CronExpr.IsNull() && !m.CronExpr.IsUnknown() {
		s.CronExpr = m.CronExpr.ValueString()
	}
	if !m.Default.IsNull() && !m.Default.IsUnknown() {
		s.Default = m.Default.ValueBool()
	}
	if !m.Enabled.IsNull() && !m.Enabled.IsUnknown() {
		s.Enabled = m.Enabled.ValueBool()
	}
	if !m.ExcludeDevices.IsNull() && !m.ExcludeDevices.IsUnknown() {
		var v []string
		diags.Append(m.ExcludeDevices.ElementsAs(ctx, &v, false)...)
		s.ExcludeDevices = v
	}
	if !m.HighPriorityDevices.IsNull() && !m.HighPriorityDevices.IsUnknown() {
		var v []string
		diags.Append(m.HighPriorityDevices.ElementsAs(ctx, &v, false)...)
		s.HighPriorityDevices = v
	}
	if !m.HtModesNa.IsNull() && !m.HtModesNa.IsUnknown() {
		var val []int64
		if err := json.Unmarshal([]byte(m.HtModesNa.ValueString()), &val); err != nil {
			diags.AddError("Invalid ht_modes_na", err.Error())
		} else {
			s.HtModesNa = val
		}
	}
	if !m.HtModesNg.IsNull() && !m.HtModesNg.IsUnknown() {
		var val []int64
		if err := json.Unmarshal([]byte(m.HtModesNg.ValueString()), &val); err != nil {
			diags.AddError("Invalid ht_modes_ng", err.Error())
		} else {
			s.HtModesNg = val
		}
	}
	if !m.Optimize.IsNull() && !m.Optimize.IsUnknown() {
		var v []string
		diags.Append(m.Optimize.ElementsAs(ctx, &v, false)...)
		s.Optimize = v
	}
	if !m.Radios.IsNull() && !m.Radios.IsUnknown() {
		var v []string
		diags.Append(m.Radios.ElementsAs(ctx, &v, false)...)
		s.Radios = v
	}
	if !m.RadiosConfiguration.IsNull() && !m.RadiosConfiguration.IsUnknown() {
		var val []settings.SettingRadioAiRadiosConfiguration
		if err := json.Unmarshal([]byte(m.RadiosConfiguration.ValueString()), &val); err != nil {
			diags.AddError("Invalid radios_configuration", err.Error())
		} else {
			s.RadiosConfiguration = val
		}
	}
	if !m.SettingPreference.IsNull() && !m.SettingPreference.IsUnknown() {
		s.SettingPreference = m.SettingPreference.ValueString()
	}
	if !m.UseXy.IsNull() && !m.UseXy.IsUnknown() {
		s.UseXy = m.UseXy.ValueBool()
	}
}
