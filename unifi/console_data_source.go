package unifi

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &consoleDataSource{}

func NewConsoleDataSource() datasource.DataSource {
	return &consoleDataSource{}
}

type consoleDataSource struct {
	client *Client
}

type consoleDataSourceModel struct {
	ID              types.String `tfsdk:"id"`
	FirmwareVersion types.String `tfsdk:"firmware_version"`
	UcoreVersion    types.String `tfsdk:"ucore_version"`
	AppVersions     types.Map    `tfsdk:"app_versions"`
}

func (d *consoleDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_console"
}

func (d *consoleDataSource) Schema(
	ctx context.Context,
	req datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Data source for the UniFi OS console (`/api/system`): the console firmware version and the versions of every installed application (Network, Protect, Access, Talk, ...). Unlike `unifi_system_info`, which only exposes the Network app version, this surfaces all installed UniFi OS apps.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The console name.",
				Computed:            true,
			},
			"firmware_version": schema.StringAttribute{
				MarkdownDescription: "The UniFi OS console firmware version (e.g. `5.1.19`).",
				Computed:            true,
			},
			"ucore_version": schema.StringAttribute{
				MarkdownDescription: "The console uCore version.",
				Computed:            true,
			},
			"app_versions": schema.MapAttribute{
				MarkdownDescription: "Installed applications keyed by name (e.g. `network`, `protect`, `access`) to their version.",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}

func (d *consoleDataSource) Configure(
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

func (d *consoleDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var data consoleDataSourceModel

	info, err := d.client.ConsoleSystem(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Console System",
			"Could not read the UniFi OS /api/system: "+err.Error(),
		)
		return
	}

	// Expose only installed apps, keyed by name → version.
	appVersions := map[string]string{}
	for _, app := range info.Apps.Controllers {
		if app.Name != "" && app.InstallState == "installed" {
			appVersions[app.Name] = app.Version
		}
	}
	appVersionsMap, diags := types.MapValueFrom(ctx, types.StringType, appVersions)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ID = types.StringValue(info.Name)
	data.FirmwareVersion = types.StringValue(info.Hardware.FirmwareVersion)
	data.UcoreVersion = types.StringValue(info.UcoreVersion)
	data.AppVersions = appVersionsMap

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
