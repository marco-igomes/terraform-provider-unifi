package unifi

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/ubiquiti-community/go-unifi/unifi"
)

var (
	_ resource.Resource                = &userGroupResource{}
	_ resource.ResourceWithImportState = &userGroupResource{}
)

func NewUserGroupResource() resource.Resource {
	return &userGroupResource{}
}

type userGroupResource struct {
	client *Client
}

type userGroupResourceModel struct {
	ID             types.String `tfsdk:"id"`
	Site           types.String `tfsdk:"site"`
	Name           types.String `tfsdk:"name"`
	QOSRateMaxDown types.Int64  `tfsdk:"qos_rate_max_down"`
	QOSRateMaxUp   types.Int64  `tfsdk:"qos_rate_max_up"`
	NoDelete       types.Bool   `tfsdk:"no_delete"`
}

func (r *userGroupResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_usergroup"
}

func (r *userGroupResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	// -1 disables the limit; otherwise the controller accepts 2..100000 kbps.
	rateValidators := []validator.Int64{
		int64validator.Any(
			int64validator.OneOf(-1),
			int64validator.Between(2, 100000),
		),
	}

	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a UniFi user (client) group. Groups carry per-client QoS rate limits applied via `unifi_client.user_group_id`.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "ID of the user group.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"site": schema.StringAttribute{
				MarkdownDescription: "Site the group lives in.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Display name (1-128 chars).",
				Required:            true,
			},
			"qos_rate_max_down": schema.Int64Attribute{
				MarkdownDescription: "Max download rate in kbps. `-1` = unlimited.",
				Optional:            true,
				Computed:            true,
				Validators:          rateValidators,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"qos_rate_max_up": schema.Int64Attribute{
				MarkdownDescription: "Max upload rate in kbps. `-1` = unlimited.",
				Optional:            true,
				Computed:            true,
				Validators:          rateValidators,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"no_delete": schema.BoolAttribute{
				MarkdownDescription: "True for the system `Default` group (cannot be deleted; Terraform Delete drops it from state only).",
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *userGroupResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *userGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan userGroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	site := r.siteOrDefault(plan.Site)

	created, err := r.client.CreateClientGroup(ctx, site, r.modelToGroup(&plan))
	if err != nil {
		resp.Diagnostics.AddError("Error Creating User Group", err.Error())
		return
	}
	r.groupToModel(created, &plan, site)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *userGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state userGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	site := r.siteOrDefault(state.Site)

	g, err := r.client.GetClientGroup(ctx, site, state.ID.ValueString())
	if err != nil {
		if _, ok := err.(*unifi.NotFoundError); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading User Group", err.Error())
		return
	}
	r.groupToModel(g, &state, site)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *userGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state userGroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	site := r.siteOrDefault(state.Site)

	g := r.modelToGroup(&plan)
	g.ID = state.ID.ValueString()

	updated, err := r.client.UpdateClientGroup(ctx, site, g)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating User Group", err.Error())
		return
	}
	r.groupToModel(updated, &plan, site)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *userGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state userGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if state.NoDelete.ValueBool() {
		// System default group is non-deletable on the controller.
		return
	}
	site := r.siteOrDefault(state.Site)
	if err := r.client.DeleteClientGroup(ctx, site, state.ID.ValueString()); err != nil {
		if _, ok := err.(*unifi.NotFoundError); !ok {
			resp.Diagnostics.AddError("Error Deleting User Group", err.Error())
		}
	}
}

func (r *userGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *userGroupResource) siteOrDefault(site types.String) string {
	if s := site.ValueString(); s != "" {
		return s
	}
	return r.client.Site
}

func (r *userGroupResource) modelToGroup(m *userGroupResourceModel) *unifi.ClientGroup {
	return &unifi.ClientGroup{
		Name:           m.Name.ValueString(),
		QOSRateMaxDown: int64PointerOrNil(m.QOSRateMaxDown),
		QOSRateMaxUp:   int64PointerOrNil(m.QOSRateMaxUp),
	}
}

func (r *userGroupResource) groupToModel(g *unifi.ClientGroup, m *userGroupResourceModel, site string) {
	m.ID = types.StringValue(g.ID)
	m.Site = types.StringValue(site)
	m.Name = types.StringValue(g.Name)
	m.QOSRateMaxDown = types.Int64PointerValue(g.QOSRateMaxDown)
	m.QOSRateMaxUp = types.Int64PointerValue(g.QOSRateMaxUp)
	m.NoDelete = types.BoolValue(g.NoDelete)
}
