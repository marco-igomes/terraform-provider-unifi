package unifi

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &systemInfoDataSource{}

func NewSystemInfoDataSource() datasource.DataSource {
	return &systemInfoDataSource{}
}

type systemInfoDataSource struct {
	client *Client
}

type systemInfoDataSourceModel struct {
	ID              types.String `tfsdk:"id"`
	Site            types.String `tfsdk:"site"`
	Version         types.String `tfsdk:"version"`
	PreviousVersion types.String `tfsdk:"previous_version"`
	UDMVersion      types.String `tfsdk:"udm_version"`
	UBNTDeviceType  types.String `tfsdk:"ubnt_device_type"`
	Timezone        types.String `tfsdk:"timezone"`
}

func (d *systemInfoDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_system_info"
}

func (d *systemInfoDataSource) Schema(
	ctx context.Context,
	req datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Data source for the UniFi controller's sysinfo, mainly the controller/application version.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The site this sysinfo was read from.",
				Computed:            true,
			},
			"site": schema.StringAttribute{
				MarkdownDescription: "The name of the site to read sysinfo from. Defaults to the provider's site.",
				Optional:            true,
				Computed:            true,
			},
			"version": schema.StringAttribute{
				MarkdownDescription: "The controller/application version.",
				Computed:            true,
			},
			"previous_version": schema.StringAttribute{
				MarkdownDescription: "The controller/application version before the last upgrade.",
				Computed:            true,
			},
			"udm_version": schema.StringAttribute{
				MarkdownDescription: "The UniFi OS (UDM) firmware version, when running on a UDM/UDM-Pro.",
				Computed:            true,
			},
			"ubnt_device_type": schema.StringAttribute{
				MarkdownDescription: "The controller's hardware device type (e.g. UDMPRO).",
				Computed:            true,
			},
			"timezone": schema.StringAttribute{
				MarkdownDescription: "The controller's configured timezone.",
				Computed:            true,
			},
		},
	}
}

func (d *systemInfoDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	resp *datasource.ConfigureResponse,
) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf(
				"Expected *Client, got: %T. Please report this issue to the provider developers.",
				req.ProviderData,
			),
		)
		return
	}

	d.client = client
}

func (d *systemInfoDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var data systemInfoDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	site := data.Site.ValueString()
	if site == "" {
		site = d.client.Site
	}

	info, err := d.client.Sysinfo(ctx, site)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Sysinfo",
			"Could not read controller sysinfo: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(site)
	data.Site = types.StringValue(site)
	data.Version = types.StringValue(info.Version)
	data.PreviousVersion = types.StringValue(info.PreviousVersion)
	data.UDMVersion = types.StringValue(info.UDMVersion)
	data.UBNTDeviceType = types.StringValue(info.UBNTDeviceType)
	data.Timezone = types.StringValue(info.Timezone)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
