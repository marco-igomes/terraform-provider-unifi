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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/ubiquiti-community/go-unifi/unifi"
)

var (
	_ resource.Resource                = &firewallPolicyResource{}
	_ resource.ResourceWithImportState = &firewallPolicyResource{}
)

func NewFirewallPolicyResource() resource.Resource {
	return &firewallPolicyResource{}
}

type firewallPolicyResource struct {
	client *Client
}

type firewallPolicyResourceModel struct {
	ID                    types.String `tfsdk:"id"`
	Site                  types.String `tfsdk:"site"`
	Name                  types.String `tfsdk:"name"`
	Description           types.String `tfsdk:"description"`
	Enabled               types.Bool   `tfsdk:"enabled"`
	Action                types.String `tfsdk:"action"`
	Protocol              types.String `tfsdk:"protocol"`
	IPVersion             types.String `tfsdk:"ip_version"`
	Index                 types.Int64  `tfsdk:"index"`
	Logging               types.Bool   `tfsdk:"logging"`
	MatchIPSec            types.Bool   `tfsdk:"match_ip_sec"`
	MatchIPSecType        types.String `tfsdk:"match_ip_sec_type"`
	MatchOppositeProtocol types.Bool   `tfsdk:"match_opposite_protocol"`
	ConnectionStateType   types.String `tfsdk:"connection_state_type"`
	ConnectionStates      types.List   `tfsdk:"connection_states"`
	CreateAllowRespond    types.Bool   `tfsdk:"create_allow_respond"`
	ICMPTypename          types.String `tfsdk:"icmp_typename"`
	ICMPV6Typename        types.String `tfsdk:"icmp_v6_typename"`
	Predefined            types.Bool   `tfsdk:"predefined"`
	OriginID              types.String `tfsdk:"origin_id"`
	OriginType            types.String `tfsdk:"origin_type"`
	Source                types.Object `tfsdk:"source"`
	Destination           types.Object `tfsdk:"destination"`
	Schedule              types.Object `tfsdk:"schedule"`
}

type firewallPolicyEndpointModel struct {
	ZoneID                types.String `tfsdk:"zone_id"`
	MatchingTarget        types.String `tfsdk:"matching_target"`
	MatchingTargetType    types.String `tfsdk:"matching_target_type"`
	IPGroupID             types.String `tfsdk:"ip_group_id"`
	IPs                   types.List   `tfsdk:"ips"`
	NetworkIDs            types.List   `tfsdk:"network_ids"`
	MatchMAC              types.Bool   `tfsdk:"match_mac"`
	MatchOppositeIPs      types.Bool   `tfsdk:"match_opposite_ips"`
	MatchOppositeNetworks types.Bool   `tfsdk:"match_opposite_networks"`
	MatchOppositePorts    types.Bool   `tfsdk:"match_opposite_ports"`
	Port                  types.Int64  `tfsdk:"port"`
	PortGroupID           types.String `tfsdk:"port_group_id"`
	PortMatchingType      types.String `tfsdk:"port_matching_type"`
}

func firewallPolicyEndpointAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"zone_id":                 types.StringType,
		"matching_target":         types.StringType,
		"matching_target_type":    types.StringType,
		"ip_group_id":             types.StringType,
		"ips":                     types.ListType{ElemType: types.StringType},
		"network_ids":             types.ListType{ElemType: types.StringType},
		"match_mac":               types.BoolType,
		"match_opposite_ips":      types.BoolType,
		"match_opposite_networks": types.BoolType,
		"match_opposite_ports":    types.BoolType,
		"port":                    types.Int64Type,
		"port_group_id":           types.StringType,
		"port_matching_type":      types.StringType,
	}
}

type firewallPolicyScheduleModel struct {
	Mode           types.String `tfsdk:"mode"`
	Date           types.String `tfsdk:"date"`
	DateStart      types.String `tfsdk:"date_start"`
	DateEnd        types.String `tfsdk:"date_end"`
	RepeatOnDays   types.List   `tfsdk:"repeat_on_days"`
	TimeAllDay     types.Bool   `tfsdk:"time_all_day"`
	TimeRangeStart types.String `tfsdk:"time_range_start"`
	TimeRangeEnd   types.String `tfsdk:"time_range_end"`
}

func firewallPolicyScheduleAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"mode":             types.StringType,
		"date":             types.StringType,
		"date_start":       types.StringType,
		"date_end":         types.StringType,
		"repeat_on_days":   types.ListType{ElemType: types.StringType},
		"time_all_day":     types.BoolType,
		"time_range_start": types.StringType,
		"time_range_end":   types.StringType,
	}
}

func (r *firewallPolicyResource) Metadata(
	_ context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_firewall_policy"
}

func endpointAttributes(label string) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"zone_id": schema.StringAttribute{
			MarkdownDescription: fmt.Sprintf("ID of the %s firewall zone.", label),
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"matching_target": schema.StringAttribute{
			MarkdownDescription: "What to match: `ANY`, `DEVICE`, `IP`, `NETWORK`, or `MAC`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"matching_target_type": schema.StringAttribute{
			MarkdownDescription: "How matching is specified: `ANY`, `SPECIFIC`, `LIST`, or `OBJECT`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"ip_group_id": schema.StringAttribute{
			MarkdownDescription: "ID of an address-group `unifi_firewall_group` (used when matching_target=IP and matching_target_type=OBJECT).",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"ips": schema.ListAttribute{
			MarkdownDescription: "IPs (used when matching_target is `IP` and matching_target_type is `SPECIFIC` or `LIST`).",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.List{listplanmodifier.UseStateForUnknown()},
		},
		"network_ids": schema.ListAttribute{
			MarkdownDescription: "Network IDs (used when matching_target is `NETWORK`). Reference `unifi_network` resources via `.id`.",
			Optional:            true,
			Computed:            true,
			ElementType:         types.StringType,
			PlanModifiers:       []planmodifier.List{listplanmodifier.UseStateForUnknown()},
		},
		"match_opposite_networks": schema.BoolAttribute{
			MarkdownDescription: "Invert the network match (match everything except the specified networks).",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"match_mac": schema.BoolAttribute{
			MarkdownDescription: "Match by MAC address.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"match_opposite_ips": schema.BoolAttribute{
			MarkdownDescription: "Invert the IP match (match everything except the specified IPs).",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"match_opposite_ports": schema.BoolAttribute{
			MarkdownDescription: "Invert the port match.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
		},
		"port": schema.Int64Attribute{
			MarkdownDescription: "Specific port (used with port_matching_type = `SPECIFIC`).",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
		},
		"port_group_id": schema.StringAttribute{
			MarkdownDescription: "ID of a port-group firewall group (used with port_matching_type = `OBJECT`).",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"port_matching_type": schema.StringAttribute{
			MarkdownDescription: "Port match type: `ANY`, `SPECIFIC`, `LIST`, or `OBJECT`.",
			Optional:            true,
			Computed:            true,
			PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
	}
}

func (r *firewallPolicyResource) Schema(
	_ context.Context,
	_ resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a zone-based firewall policy (UniFi 9+). Policies reference `unifi_firewall_zone` resources via `source.zone_id` / `destination.zone_id`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"site": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Policy name.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"enabled": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"action": schema.StringAttribute{
				MarkdownDescription: "Policy action: `ALLOW`, `BLOCK`, or `REJECT`.",
				Required:            true,
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: "Protocol: `all`, `tcp`, `udp`, or `tcp_udp`.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ip_version": schema.StringAttribute{
				MarkdownDescription: "IP version: `BOTH`, `IPV4`, or `IPV6`.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"index": schema.Int64Attribute{
				MarkdownDescription: "Ordering index. Lower-numbered policies are evaluated first.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"logging": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"match_ip_sec": schema.BoolAttribute{
				MarkdownDescription: "Match only IPsec-encapsulated traffic.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"match_ip_sec_type": schema.StringAttribute{
				MarkdownDescription: "IPsec match type: `MATCH_IP_SEC` or `MATCH_NON_IP_SEC`.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"origin_id": schema.StringAttribute{
				MarkdownDescription: "System-set: links the policy back to a UniFi-generated source (e.g. wifiman). Read-only.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"origin_type": schema.StringAttribute{
				MarkdownDescription: "System-set: classification of `origin_id`. Read-only.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"match_opposite_protocol": schema.BoolAttribute{
				MarkdownDescription: "Invert the protocol match.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"connection_state_type": schema.StringAttribute{
				MarkdownDescription: "Connection state matching: `ALL` or `RESPOND_ONLY`.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"connection_states": schema.ListAttribute{
				MarkdownDescription: "Connection state names to match (e.g. `NEW`, `ESTABLISHED`, `RELATED`, `INVALID`).",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				PlanModifiers:       []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"create_allow_respond": schema.BoolAttribute{
				MarkdownDescription: "Auto-create the reverse allow-respond rule.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"icmp_typename": schema.StringAttribute{
				MarkdownDescription: "ICMPv4 type name filter: `ANY`, `SPECIFIC`, `LIST`, or `OBJECT`.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"icmp_v6_typename": schema.StringAttribute{
				MarkdownDescription: "ICMPv6 type name filter: `ANY`, `SPECIFIC`, `LIST`, or `OBJECT`.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"predefined": schema.BoolAttribute{
				MarkdownDescription: "Whether this is a system-managed predefined policy. Read-only.",
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"source": schema.SingleNestedAttribute{
				MarkdownDescription: "Source matching.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes:          endpointAttributes("source"),
			},
			"destination": schema.SingleNestedAttribute{
				MarkdownDescription: "Destination matching.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes:          endpointAttributes("destination"),
			},
			"schedule": schema.SingleNestedAttribute{
				MarkdownDescription: "When the policy applies. Omit for `ALWAYS`.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"mode": schema.StringAttribute{
						MarkdownDescription: "`ALWAYS`, `EVERY_DAY`, `EVERY_WEEK`, or `ONE_TIME_ONLY`.",
						Optional:            true,
						Computed:            true,
						PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"date": schema.StringAttribute{
						MarkdownDescription: "Date for ONE_TIME_ONLY mode (YYYY-MM-DD).",
						Optional:            true,
						Computed:            true,
						PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"date_start": schema.StringAttribute{
						MarkdownDescription: "Start date for windowed schedules (ISO8601).",
						Optional:            true,
						Computed:            true,
						PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"date_end": schema.StringAttribute{
						MarkdownDescription: "End date for windowed schedules (ISO8601).",
						Optional:            true,
						Computed:            true,
						PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"repeat_on_days": schema.ListAttribute{
						MarkdownDescription: "Days of week for EVERY_WEEK mode (`mon`, `tue`, …).",
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
						PlanModifiers:       []planmodifier.List{listplanmodifier.UseStateForUnknown()},
					},
					"time_all_day": schema.BoolAttribute{
						Optional:      true,
						Computed:      true,
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"time_range_start": schema.StringAttribute{
						MarkdownDescription: "Start time `HH:MM` when time_all_day is false.",
						Optional:            true,
						Computed:            true,
						PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"time_range_end": schema.StringAttribute{
						MarkdownDescription: "End time `HH:MM` when time_all_day is false.",
						Optional:            true,
						Computed:            true,
						PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
		},
	}
}

func (r *firewallPolicyResource) Configure(
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

func (r *firewallPolicyResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var plan firewallPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	site := plan.Site.ValueString()
	if site == "" {
		site = r.client.Site
	}

	policy, diags := r.modelToPolicy(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	created, err := r.client.CreateFirewallPolicy(ctx, site, policy)
	if err != nil {
		resp.Diagnostics.AddError("Error Creating Firewall Policy", err.Error())
		return
	}

	resp.Diagnostics.Append(r.policyToModel(ctx, created, &plan, site)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *firewallPolicyResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var state firewallPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	site := state.Site.ValueString()
	if site == "" {
		site = r.client.Site
	}

	policy, err := r.client.GetFirewallPolicy(ctx, site, state.ID.ValueString())
	if err != nil {
		if _, ok := err.(*unifi.NotFoundError); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading Firewall Policy", err.Error())
		return
	}

	resp.Diagnostics.Append(r.policyToModel(ctx, policy, &state, site)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *firewallPolicyResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var plan, state firewallPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	site := state.Site.ValueString()
	if site == "" {
		site = r.client.Site
	}

	policy, diags := r.modelToPolicy(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	policy.ID = state.ID.ValueString()

	updated, err := r.client.UpdateFirewallPolicy(ctx, site, policy)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating Firewall Policy", err.Error())
		return
	}

	resp.Diagnostics.Append(r.policyToModel(ctx, updated, &plan, site)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *firewallPolicyResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var state firewallPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	site := state.Site.ValueString()
	if site == "" {
		site = r.client.Site
	}

	if state.Predefined.ValueBool() {
		// System-predefined policies can't be deleted; drop from state.
		return
	}

	if err := r.client.DeleteFirewallPolicy(ctx, site, state.ID.ValueString()); err != nil {
		if _, ok := err.(*unifi.NotFoundError); !ok {
			resp.Diagnostics.AddError("Error Deleting Firewall Policy", err.Error())
		}
	}
}

func (r *firewallPolicyResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *firewallPolicyResource) modelToPolicy(
	ctx context.Context,
	model *firewallPolicyResourceModel,
) (*unifi.FirewallPolicy, diag.Diagnostics) {
	var diags diag.Diagnostics

	policy := &unifi.FirewallPolicy{
		Name:                  model.Name.ValueString(),
		Action:                model.Action.ValueString(),
		Description:           model.Description.ValueString(),
		Enabled:               model.Enabled.ValueBool(),
		Protocol:              model.Protocol.ValueString(),
		Version:               model.IPVersion.ValueString(),
		Logging:               model.Logging.ValueBool(),
		MatchIPSec:            model.MatchIPSec.ValueBool(),
		MatchIPSecType:        model.MatchIPSecType.ValueString(),
		MatchOppositeProtocol: model.MatchOppositeProtocol.ValueBool(),
		ConnectionStateType:   model.ConnectionStateType.ValueString(),
		CreateAllowRespond:    model.CreateAllowRespond.ValueBool(),
		ICMPTypename:          model.ICMPTypename.ValueString(),
		ICMPV6Typename:        model.ICMPV6Typename.ValueString(),
		Index:                 int64PointerOrNil(model.Index),
		OriginID:              model.OriginID.ValueString(),
		OriginType:            model.OriginType.ValueString(),
	}

	if !model.ConnectionStates.IsNull() && !model.ConnectionStates.IsUnknown() {
		var states []string
		diags.Append(model.ConnectionStates.ElementsAs(ctx, &states, false)...)
		if !diags.HasError() {
			policy.ConnectionStates = states
		}
	}

	if !model.Source.IsNull() && !model.Source.IsUnknown() {
		src, d := endpointModelToSource(ctx, model.Source)
		diags.Append(d...)
		if !diags.HasError() {
			policy.Source = src
		}
	}
	if !model.Destination.IsNull() && !model.Destination.IsUnknown() {
		dst, d := endpointModelToDestination(ctx, model.Destination)
		diags.Append(d...)
		if !diags.HasError() {
			policy.Destination = dst
		}
	}
	if !model.Schedule.IsNull() && !model.Schedule.IsUnknown() {
		sched, d := scheduleModelToAPI(ctx, model.Schedule)
		diags.Append(d...)
		if !diags.HasError() {
			policy.Schedule = sched
		}
	}

	return policy, diags
}

func endpointModelToSource(
	ctx context.Context,
	obj types.Object,
) (*unifi.FirewallPolicySource, diag.Diagnostics) {
	var diags diag.Diagnostics
	var ep firewallPolicyEndpointModel
	diags.Append(obj.As(ctx, &ep, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil, diags
	}
	src := &unifi.FirewallPolicySource{
		ZoneID:                ep.ZoneID.ValueString(),
		MatchingTarget:        ep.MatchingTarget.ValueString(),
		MatchingTargetType:    ep.MatchingTargetType.ValueString(),
		IPGroupID:             ep.IPGroupID.ValueString(),
		MatchMAC:              ep.MatchMAC.ValueBool(),
		MatchOppositeIPs:      ep.MatchOppositeIPs.ValueBool(),
		MatchOppositeNetworks: ep.MatchOppositeNetworks.ValueBool(),
		MatchOppositePorts:    ep.MatchOppositePorts.ValueBool(),
		Port:                  int64PointerOrNil(ep.Port),
		PortGroupID:           ep.PortGroupID.ValueString(),
		PortMatchingType:      ep.PortMatchingType.ValueString(),
	}
	if !ep.IPs.IsNull() && !ep.IPs.IsUnknown() {
		var ips []string
		diags.Append(ep.IPs.ElementsAs(ctx, &ips, false)...)
		if !diags.HasError() {
			src.IPs = ips
		}
	}
	if !ep.NetworkIDs.IsNull() && !ep.NetworkIDs.IsUnknown() {
		var nids []string
		diags.Append(ep.NetworkIDs.ElementsAs(ctx, &nids, false)...)
		if !diags.HasError() {
			src.NetworkIDs = nids
		}
	}
	return src, diags
}

func endpointModelToDestination(
	ctx context.Context,
	obj types.Object,
) (*unifi.FirewallPolicyDestination, diag.Diagnostics) {
	var diags diag.Diagnostics
	var ep firewallPolicyEndpointModel
	diags.Append(obj.As(ctx, &ep, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil, diags
	}
	dst := &unifi.FirewallPolicyDestination{
		ZoneID:                ep.ZoneID.ValueString(),
		MatchingTarget:        ep.MatchingTarget.ValueString(),
		MatchingTargetType:    ep.MatchingTargetType.ValueString(),
		IPGroupID:             ep.IPGroupID.ValueString(),
		MatchMAC:              ep.MatchMAC.ValueBool(),
		MatchOppositeIPs:      ep.MatchOppositeIPs.ValueBool(),
		MatchOppositeNetworks: ep.MatchOppositeNetworks.ValueBool(),
		MatchOppositePorts:    ep.MatchOppositePorts.ValueBool(),
		Port:                  int64PointerOrNil(ep.Port),
		PortGroupID:           ep.PortGroupID.ValueString(),
		PortMatchingType:      ep.PortMatchingType.ValueString(),
	}
	if !ep.IPs.IsNull() && !ep.IPs.IsUnknown() {
		var ips []string
		diags.Append(ep.IPs.ElementsAs(ctx, &ips, false)...)
		if !diags.HasError() {
			dst.IPs = ips
		}
	}
	if !ep.NetworkIDs.IsNull() && !ep.NetworkIDs.IsUnknown() {
		var nids []string
		diags.Append(ep.NetworkIDs.ElementsAs(ctx, &nids, false)...)
		if !diags.HasError() {
			dst.NetworkIDs = nids
		}
	}
	return dst, diags
}

func scheduleModelToAPI(
	ctx context.Context,
	obj types.Object,
) (*unifi.FirewallPolicySchedule, diag.Diagnostics) {
	var diags diag.Diagnostics
	var s firewallPolicyScheduleModel
	diags.Append(obj.As(ctx, &s, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil, diags
	}
	sched := &unifi.FirewallPolicySchedule{
		Mode:           s.Mode.ValueString(),
		Date:           s.Date.ValueString(),
		DateStart:      s.DateStart.ValueString(),
		DateEnd:        s.DateEnd.ValueString(),
		TimeAllDay:     s.TimeAllDay.ValueBool(),
		TimeRangeStart: s.TimeRangeStart.ValueString(),
		TimeRangeEnd:   s.TimeRangeEnd.ValueString(),
	}
	if !s.RepeatOnDays.IsNull() && !s.RepeatOnDays.IsUnknown() {
		var days []string
		diags.Append(s.RepeatOnDays.ElementsAs(ctx, &days, false)...)
		if !diags.HasError() {
			sched.RepeatOnDays = days
		}
	}
	return sched, diags
}

func (r *firewallPolicyResource) policyToModel(
	ctx context.Context,
	policy *unifi.FirewallPolicy,
	model *firewallPolicyResourceModel,
	site string,
) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = types.StringValue(policy.ID)
	model.Site = types.StringValue(site)
	model.Name = types.StringValue(policy.Name)
	model.Description = stringOrNull(policy.Description)
	model.Enabled = types.BoolValue(policy.Enabled)
	model.Action = types.StringValue(policy.Action)
	model.Protocol = stringOrNull(policy.Protocol)
	model.IPVersion = stringOrNull(policy.Version)
	model.Index = types.Int64PointerValue(policy.Index)
	model.Logging = types.BoolValue(policy.Logging)
	model.MatchIPSec = types.BoolValue(policy.MatchIPSec)
	model.MatchOppositeProtocol = types.BoolValue(policy.MatchOppositeProtocol)
	model.MatchIPSecType = stringOrNull(policy.MatchIPSecType)
	model.ConnectionStateType = stringOrNull(policy.ConnectionStateType)
	model.CreateAllowRespond = types.BoolValue(policy.CreateAllowRespond)
	model.ICMPTypename = stringOrNull(policy.ICMPTypename)
	model.ICMPV6Typename = stringOrNull(policy.ICMPV6Typename)
	model.Predefined = types.BoolValue(policy.Predefined)
	model.OriginID = stringOrNull(policy.OriginID)
	model.OriginType = stringOrNull(policy.OriginType)

	if len(policy.ConnectionStates) == 0 {
		model.ConnectionStates = types.ListValueMust(types.StringType, nil)
	} else {
		vals := make([]attr.Value, len(policy.ConnectionStates))
		for i, s := range policy.ConnectionStates {
			vals[i] = types.StringValue(s)
		}
		listVal, d := types.ListValue(types.StringType, vals)
		diags.Append(d...)
		model.ConnectionStates = listVal
	}

	model.Source = endpointSourceToObject(ctx, policy.Source, &diags)
	model.Destination = endpointDestinationToObject(ctx, policy.Destination, &diags)
	model.Schedule = scheduleToObject(ctx, policy.Schedule, &diags)

	return diags
}

func endpointSourceToObject(
	ctx context.Context,
	src *unifi.FirewallPolicySource,
	diags *diag.Diagnostics,
) types.Object {
	if src == nil {
		return types.ObjectNull(firewallPolicyEndpointAttrTypes())
	}
	ep := firewallPolicyEndpointModel{
		ZoneID:                stringOrNull(src.ZoneID),
		MatchingTarget:        stringOrNull(src.MatchingTarget),
		MatchingTargetType:    stringOrNull(src.MatchingTargetType),
		IPGroupID:             stringOrNull(src.IPGroupID),
		MatchMAC:              types.BoolValue(src.MatchMAC),
		MatchOppositeIPs:      types.BoolValue(src.MatchOppositeIPs),
		MatchOppositeNetworks: types.BoolValue(src.MatchOppositeNetworks),
		MatchOppositePorts:    types.BoolValue(src.MatchOppositePorts),
		Port:                  types.Int64PointerValue(src.Port),
		PortGroupID:           stringOrNull(src.PortGroupID),
		PortMatchingType:      stringOrNull(src.PortMatchingType),
		IPs:                   stringList(ctx, src.IPs, diags),
		NetworkIDs:            stringList(ctx, src.NetworkIDs, diags),
	}
	obj, d := types.ObjectValueFrom(ctx, firewallPolicyEndpointAttrTypes(), ep)
	diags.Append(d...)
	return obj
}

func endpointDestinationToObject(
	ctx context.Context,
	dst *unifi.FirewallPolicyDestination,
	diags *diag.Diagnostics,
) types.Object {
	if dst == nil {
		return types.ObjectNull(firewallPolicyEndpointAttrTypes())
	}
	ep := firewallPolicyEndpointModel{
		ZoneID:                stringOrNull(dst.ZoneID),
		MatchingTarget:        stringOrNull(dst.MatchingTarget),
		MatchingTargetType:    stringOrNull(dst.MatchingTargetType),
		IPGroupID:             stringOrNull(dst.IPGroupID),
		MatchMAC:              types.BoolValue(dst.MatchMAC),
		MatchOppositeIPs:      types.BoolValue(dst.MatchOppositeIPs),
		MatchOppositeNetworks: types.BoolValue(dst.MatchOppositeNetworks),
		MatchOppositePorts:    types.BoolValue(dst.MatchOppositePorts),
		Port:                  types.Int64PointerValue(dst.Port),
		PortGroupID:           stringOrNull(dst.PortGroupID),
		PortMatchingType:      stringOrNull(dst.PortMatchingType),
		IPs:                   stringList(ctx, dst.IPs, diags),
		NetworkIDs:            stringList(ctx, dst.NetworkIDs, diags),
	}
	obj, d := types.ObjectValueFrom(ctx, firewallPolicyEndpointAttrTypes(), ep)
	diags.Append(d...)
	return obj
}

func scheduleToObject(
	ctx context.Context,
	s *unifi.FirewallPolicySchedule,
	diags *diag.Diagnostics,
) types.Object {
	if s == nil {
		return types.ObjectNull(firewallPolicyScheduleAttrTypes())
	}
	sm := firewallPolicyScheduleModel{
		Mode:           stringOrNull(s.Mode),
		Date:           stringOrNull(s.Date),
		DateStart:      stringOrNull(s.DateStart),
		DateEnd:        stringOrNull(s.DateEnd),
		TimeAllDay:     types.BoolValue(s.TimeAllDay),
		TimeRangeStart: stringOrNull(s.TimeRangeStart),
		TimeRangeEnd:   stringOrNull(s.TimeRangeEnd),
		RepeatOnDays:   stringList(ctx, s.RepeatOnDays, diags),
	}
	obj, d := types.ObjectValueFrom(ctx, firewallPolicyScheduleAttrTypes(), sm)
	diags.Append(d...)
	return obj
}

// stringList converts a slice to a list attribute, preserving an empty slice as
// an empty list (not null) so empty lists round-trip cleanly and can be declared
// explicitly in HCL.
func stringList(
	ctx context.Context,
	xs []string,
	diags *diag.Diagnostics,
) types.List {
	vals := make([]attr.Value, len(xs))
	for i, s := range xs {
		vals[i] = types.StringValue(s)
	}
	listVal, d := types.ListValue(types.StringType, vals)
	diags.Append(d...)
	_ = ctx
	return listVal
}
