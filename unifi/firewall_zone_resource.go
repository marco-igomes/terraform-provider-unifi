package unifi

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/ubiquiti-community/go-unifi/unifi"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &firewallZoneResource{}
	_ resource.ResourceWithImportState = &firewallZoneResource{}
)

func NewFirewallZoneResource() resource.Resource {
	return &firewallZoneResource{}
}

type firewallZoneResource struct {
	client *Client
}

type firewallZoneResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Site        types.String `tfsdk:"site"`
	Name        types.String `tfsdk:"name"`
	NetworkIDs  types.Set    `tfsdk:"network_ids"`
	ZoneKey     types.String `tfsdk:"zone_key"`
	DefaultZone types.Bool   `tfsdk:"default_zone"`
}

func (r *firewallZoneResource) Metadata(
	_ context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_firewall_zone"
}

func (r *firewallZoneResource) Schema(
	_ context.Context,
	_ resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a firewall zone (UniFi 9+ zone-based firewall). Zones group networks together so policies can reference them as source/destination.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the firewall zone.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"site": schema.StringAttribute{
				MarkdownDescription: "The site to associate the zone with.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Human-readable name of the zone.",
				Required:            true,
			},
			"network_ids": schema.SetAttribute{
				MarkdownDescription: "Network IDs that belong to this zone.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
			},
			"zone_key": schema.StringAttribute{
				MarkdownDescription: "Stable controller-assigned key for the zone (e.g. `main`, `iot`). Read-only after creation.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"default_zone": schema.BoolAttribute{
				MarkdownDescription: "Whether this is a system-managed default zone (e.g. `internal`, `external`, `dmz`, `gateway`). System zones cannot be deleted.",
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *firewallZoneResource) Configure(
	_ context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *Client, got: %T.", req.ProviderData),
		)
		return
	}
	r.client = client
}

func (r *firewallZoneResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var plan firewallZoneResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	site := plan.Site.ValueString()
	if site == "" {
		site = r.client.Site
	}

	zone, diags := r.modelToZone(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	created, err := r.client.CreateFirewallZone(ctx, site, zone)
	if err != nil {
		resp.Diagnostics.AddError("Error Creating Firewall Zone", err.Error())
		return
	}

	resp.Diagnostics.Append(r.zoneToModel(created, &plan, site)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *firewallZoneResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var state firewallZoneResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	site := state.Site.ValueString()
	if site == "" {
		site = r.client.Site
	}

	zone, err := r.client.GetFirewallZone(ctx, site, state.ID.ValueString())
	if err != nil {
		if _, ok := err.(*unifi.NotFoundError); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading Firewall Zone", err.Error())
		return
	}

	resp.Diagnostics.Append(r.zoneToModel(zone, &state, site)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *firewallZoneResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var plan, state firewallZoneResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	site := state.Site.ValueString()
	if site == "" {
		site = r.client.Site
	}

	zone, diags := r.modelToZone(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	zone.ID = state.ID.ValueString()

	updated, err := r.client.UpdateFirewallZone(ctx, site, zone)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating Firewall Zone", err.Error())
		return
	}

	resp.Diagnostics.Append(r.zoneToModel(updated, &plan, site)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *firewallZoneResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var state firewallZoneResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	site := state.Site.ValueString()
	if site == "" {
		site = r.client.Site
	}

	if state.DefaultZone.ValueBool() {
		// System zones can't be deleted; just drop from state.
		return
	}

	if err := r.client.DeleteFirewallZone(ctx, site, state.ID.ValueString()); err != nil {
		if _, ok := err.(*unifi.NotFoundError); !ok {
			resp.Diagnostics.AddError("Error Deleting Firewall Zone", err.Error())
		}
	}
}

func (r *firewallZoneResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *firewallZoneResource) modelToZone(
	ctx context.Context,
	model *firewallZoneResourceModel,
) (*unifi.FirewallZone, diag.Diagnostics) {
	var diags diag.Diagnostics
	zone := &unifi.FirewallZone{
		Name: model.Name.ValueString(),
	}
	if !model.NetworkIDs.IsNull() && !model.NetworkIDs.IsUnknown() {
		var ids []string
		diags.Append(model.NetworkIDs.ElementsAs(ctx, &ids, false)...)
		if !diags.HasError() {
			zone.NetworkIDs = ids
		}
	}
	return zone, diags
}

func (r *firewallZoneResource) zoneToModel(
	zone *unifi.FirewallZone,
	model *firewallZoneResourceModel,
	site string,
) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = types.StringValue(zone.ID)
	model.Site = types.StringValue(site)
	model.Name = types.StringValue(zone.Name)
	model.ZoneKey = stringOrNull(zone.ZoneKey)
	model.DefaultZone = types.BoolValue(zone.DefaultZone)

	if len(zone.NetworkIDs) == 0 {
		model.NetworkIDs = types.SetValueMust(types.StringType, nil)
		return diags
	}
	vals := make([]attr.Value, len(zone.NetworkIDs))
	for i, id := range zone.NetworkIDs {
		vals[i] = types.StringValue(id)
	}
	setVal, d := types.SetValue(types.StringType, vals)
	diags.Append(d...)
	if diags.HasError() {
		model.NetworkIDs = types.SetNull(types.StringType)
	} else {
		model.NetworkIDs = setVal
	}
	return diags
}
