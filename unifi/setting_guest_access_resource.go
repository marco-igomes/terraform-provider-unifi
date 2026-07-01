package unifi

// Generated from go-unifi/unifi/settings/guest_access.generated.go.
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
	_ resource.Resource                = &settingGuestAccessResource{}
	_ resource.ResourceWithImportState = &settingGuestAccessResource{}
)

func NewSettingGuestAccessResource() resource.Resource {
	return &settingGuestAccessResource{}
}

type settingGuestAccessResource struct {
	client *Client
}

type settingGuestAccessModel struct {
	ID                                     types.String `tfsdk:"id"`
	Site                                   types.String `tfsdk:"site"`
	AllowedSubnet                          types.String `tfsdk:"allowed_subnet_"`
	Auth                                   types.String `tfsdk:"auth"`
	AuthUrl                                types.String `tfsdk:"auth_url"`
	AuthorizeLoginid                       types.String `tfsdk:"x_authorize_loginid"`
	AuthorizeTransactionkey                types.String `tfsdk:"x_authorize_transactionkey"`
	AuthorizeUseSandbox                    types.Bool   `tfsdk:"authorize_use_sandbox"`
	CustomIP                               types.String `tfsdk:"custom_ip"`
	EcEnabled                              types.Bool   `tfsdk:"ec_enabled"`
	Expire                                 types.String `tfsdk:"expire"`
	ExpireNumber                           types.Int64  `tfsdk:"expire_number"`
	ExpireUnit                             types.Int64  `tfsdk:"expire_unit"`
	FacebookAppID                          types.String `tfsdk:"facebook_app_id"`
	FacebookAppSecret                      types.String `tfsdk:"x_facebook_app_secret"`
	FacebookEnabled                        types.Bool   `tfsdk:"facebook_enabled"`
	FacebookScopeEmail                     types.Bool   `tfsdk:"facebook_scope_email"`
	FacebookWifiBlockHttps                 types.Bool   `tfsdk:"facebook_wifi_block_https"`
	FacebookWifiGwID                       types.String `tfsdk:"facebook_wifi_gw_id"`
	FacebookWifiGwName                     types.String `tfsdk:"facebook_wifi_gw_name"`
	FacebookWifiGwSecret                   types.String `tfsdk:"x_facebook_wifi_gw_secret"`
	Gateway                                types.String `tfsdk:"gateway"`
	GoogleClientID                         types.String `tfsdk:"google_client_id"`
	GoogleClientSecret                     types.String `tfsdk:"x_google_client_secret"`
	GoogleDomain                           types.String `tfsdk:"google_domain"`
	GoogleEnabled                          types.Bool   `tfsdk:"google_enabled"`
	GoogleScopeEmail                       types.Bool   `tfsdk:"google_scope_email"`
	IPpayTerminalid                        types.String `tfsdk:"x_ippay_terminalid"`
	IPpayUseSandbox                        types.Bool   `tfsdk:"ippay_use_sandbox"`
	MerchantwarriorApikey                  types.String `tfsdk:"x_merchantwarrior_apikey"`
	MerchantwarriorApipassphrase           types.String `tfsdk:"x_merchantwarrior_apipassphrase"`
	MerchantwarriorMerchantuuid            types.String `tfsdk:"x_merchantwarrior_merchantuuid"`
	MerchantwarriorUseSandbox              types.Bool   `tfsdk:"merchantwarrior_use_sandbox"`
	Password                               types.String `tfsdk:"x_password"`
	PasswordEnabled                        types.Bool   `tfsdk:"password_enabled"`
	PaymentEnabled                         types.Bool   `tfsdk:"payment_enabled"`
	PaypalPassword                         types.String `tfsdk:"x_paypal_password"`
	PaypalSignature                        types.String `tfsdk:"x_paypal_signature"`
	PaypalUseSandbox                       types.Bool   `tfsdk:"paypal_use_sandbox"`
	PaypalUsername                         types.String `tfsdk:"x_paypal_username"`
	PortalCustomized                       types.Bool   `tfsdk:"portal_customized"`
	PortalCustomizedAuthenticationText     types.String `tfsdk:"portal_customized_authentication_text"`
	PortalCustomizedBgColor                types.String `tfsdk:"portal_customized_bg_color"`
	PortalCustomizedBgImageEnabled         types.Bool   `tfsdk:"portal_customized_bg_image_enabled"`
	PortalCustomizedBgImageFilename        types.String `tfsdk:"portal_customized_bg_image_filename"`
	PortalCustomizedBgImageTile            types.Bool   `tfsdk:"portal_customized_bg_image_tile"`
	PortalCustomizedBgType                 types.String `tfsdk:"portal_customized_bg_type"`
	PortalCustomizedBoxColor               types.String `tfsdk:"portal_customized_box_color"`
	PortalCustomizedBoxLinkColor           types.String `tfsdk:"portal_customized_box_link_color"`
	PortalCustomizedBoxOpacity             types.Int64  `tfsdk:"portal_customized_box_opacity"`
	PortalCustomizedBoxRADIUS              types.Int64  `tfsdk:"portal_customized_box_radius"`
	PortalCustomizedBoxTextColor           types.String `tfsdk:"portal_customized_box_text_color"`
	PortalCustomizedButtonColor            types.String `tfsdk:"portal_customized_button_color"`
	PortalCustomizedButtonText             types.String `tfsdk:"portal_customized_button_text"`
	PortalCustomizedButtonTextColor        types.String `tfsdk:"portal_customized_button_text_color"`
	PortalCustomizedLanguages              types.List   `tfsdk:"portal_customized_languages"`
	PortalCustomizedLinkColor              types.String `tfsdk:"portal_customized_link_color"`
	PortalCustomizedLogoEnabled            types.Bool   `tfsdk:"portal_customized_logo_enabled"`
	PortalCustomizedLogoFilename           types.String `tfsdk:"portal_customized_logo_filename"`
	PortalCustomizedLogoPosition           types.String `tfsdk:"portal_customized_logo_position"`
	PortalCustomizedLogoSize               types.Int64  `tfsdk:"portal_customized_logo_size"`
	PortalCustomizedSuccessText            types.String `tfsdk:"portal_customized_success_text"`
	PortalCustomizedTextColor              types.String `tfsdk:"portal_customized_text_color"`
	PortalCustomizedTitle                  types.String `tfsdk:"portal_customized_title"`
	PortalCustomizedTos                    types.String `tfsdk:"portal_customized_tos"`
	PortalCustomizedTosEnabled             types.Bool   `tfsdk:"portal_customized_tos_enabled"`
	PortalCustomizedUnsplashAuthorName     types.String `tfsdk:"portal_customized_unsplash_author_name"`
	PortalCustomizedUnsplashAuthorUsername types.String `tfsdk:"portal_customized_unsplash_author_username"`
	PortalCustomizedWelcomeText            types.String `tfsdk:"portal_customized_welcome_text"`
	PortalCustomizedWelcomeTextEnabled     types.Bool   `tfsdk:"portal_customized_welcome_text_enabled"`
	PortalCustomizedWelcomeTextPosition    types.String `tfsdk:"portal_customized_welcome_text_position"`
	PortalEnabled                          types.Bool   `tfsdk:"portal_enabled"`
	PortalHostname                         types.String `tfsdk:"portal_hostname"`
	PortalUseHostname                      types.Bool   `tfsdk:"portal_use_hostname"`
	QuickpayAgreementid                    types.String `tfsdk:"x_quickpay_agreementid"`
	QuickpayApikey                         types.String `tfsdk:"x_quickpay_apikey"`
	QuickpayMerchantid                     types.String `tfsdk:"x_quickpay_merchantid"`
	QuickpayTestmode                       types.Bool   `tfsdk:"quickpay_testmode"`
	RADIUSAuthType                         types.String `tfsdk:"radius_auth_type"`
	RADIUSDisconnectEnabled                types.Bool   `tfsdk:"radius_disconnect_enabled"`
	RADIUSDisconnectPort                   types.Int64  `tfsdk:"radius_disconnect_port"`
	RADIUSEnabled                          types.Bool   `tfsdk:"radius_enabled"`
	RADIUSProfileID                        types.String `tfsdk:"radiusprofile_id"`
	RedirectEnabled                        types.Bool   `tfsdk:"redirect_enabled"`
	RedirectHttps                          types.Bool   `tfsdk:"redirect_https"`
	RedirectToHttps                        types.Bool   `tfsdk:"redirect_to_https"`
	RedirectUrl                            types.String `tfsdk:"redirect_url"`
	RestrictedDNSEnabled                   types.Bool   `tfsdk:"restricted_dns_enabled"`
	RestrictedDNSServers                   types.List   `tfsdk:"restricted_dns_servers"`
	RestrictedSubnet                       types.String `tfsdk:"restricted_subnet_"`
	StripeApiKey                           types.String `tfsdk:"x_stripe_api_key"`
	TemplateEngine                         types.String `tfsdk:"template_engine"`
	VoucherCustomized                      types.Bool   `tfsdk:"voucher_customized"`
	VoucherEnabled                         types.Bool   `tfsdk:"voucher_enabled"`
	WechatAppID                            types.String `tfsdk:"wechat_app_id"`
	WechatAppSecret                        types.String `tfsdk:"x_wechat_app_secret"`
	WechatEnabled                          types.Bool   `tfsdk:"wechat_enabled"`
	WechatSecretKey                        types.String `tfsdk:"x_wechat_secret_key"`
	WechatShopID                           types.String `tfsdk:"wechat_shop_id"`
}

func (r *settingGuestAccessResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_setting_guest_access"
}

func (r *settingGuestAccessResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Captive portal / guest hotspot configuration.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"site": schema.StringAttribute{
				Optional: true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace(), stringplanmodifier.UseStateForUnknown()},
			},
			"allowed_subnet_": schema.StringAttribute{
				MarkdownDescription: "allowed_subnet_ field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"auth": schema.StringAttribute{
				MarkdownDescription: "none|hotspot|facebook_wifi|custom",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"auth_url": schema.StringAttribute{
				MarkdownDescription: "auth_url field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"x_authorize_loginid": schema.StringAttribute{
				MarkdownDescription: "x_authorize_loginid field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"x_authorize_transactionkey": schema.StringAttribute{
				MarkdownDescription: "x_authorize_transactionkey field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"authorize_use_sandbox": schema.BoolAttribute{
				MarkdownDescription: "authorize_use_sandbox field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"custom_ip": schema.StringAttribute{
				MarkdownDescription: "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ec_enabled": schema.BoolAttribute{
				MarkdownDescription: "ec_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"expire": schema.StringAttribute{
				MarkdownDescription: "[\\d]+|custom",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"expire_number": schema.Int64Attribute{
				MarkdownDescription: "^[1-9][0-9]{0,5}|1000000$",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"expire_unit": schema.Int64Attribute{
				MarkdownDescription: "1|60|1440",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"facebook_app_id": schema.StringAttribute{
				MarkdownDescription: "facebook_app_id field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"x_facebook_app_secret": schema.StringAttribute{
				MarkdownDescription: "x_facebook_app_secret field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"facebook_enabled": schema.BoolAttribute{
				MarkdownDescription: "facebook_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"facebook_scope_email": schema.BoolAttribute{
				MarkdownDescription: "facebook_scope_email field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"facebook_wifi_block_https": schema.BoolAttribute{
				MarkdownDescription: "facebook_wifi_block_https field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"facebook_wifi_gw_id": schema.StringAttribute{
				MarkdownDescription: "facebook_wifi_gw_id field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"facebook_wifi_gw_name": schema.StringAttribute{
				MarkdownDescription: "facebook_wifi_gw_name field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"x_facebook_wifi_gw_secret": schema.StringAttribute{
				MarkdownDescription: "x_facebook_wifi_gw_secret field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"gateway": schema.StringAttribute{
				MarkdownDescription: "paypal|stripe|authorize|quickpay|merchantwarrior|ippay",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"google_client_id": schema.StringAttribute{
				MarkdownDescription: "google_client_id field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"x_google_client_secret": schema.StringAttribute{
				MarkdownDescription: "x_google_client_secret field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"google_domain": schema.StringAttribute{
				MarkdownDescription: "google_domain field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"google_enabled": schema.BoolAttribute{
				MarkdownDescription: "google_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"google_scope_email": schema.BoolAttribute{
				MarkdownDescription: "google_scope_email field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"x_ippay_terminalid": schema.StringAttribute{
				MarkdownDescription: "x_ippay_terminalid field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ippay_use_sandbox": schema.BoolAttribute{
				MarkdownDescription: "ippay_use_sandbox field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"x_merchantwarrior_apikey": schema.StringAttribute{
				MarkdownDescription: "x_merchantwarrior_apikey field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"x_merchantwarrior_apipassphrase": schema.StringAttribute{
				MarkdownDescription: "x_merchantwarrior_apipassphrase field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"x_merchantwarrior_merchantuuid": schema.StringAttribute{
				MarkdownDescription: "x_merchantwarrior_merchantuuid field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"merchantwarrior_use_sandbox": schema.BoolAttribute{
				MarkdownDescription: "merchantwarrior_use_sandbox field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"x_password": schema.StringAttribute{
				MarkdownDescription: "x_password field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"password_enabled": schema.BoolAttribute{
				MarkdownDescription: "password_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"payment_enabled": schema.BoolAttribute{
				MarkdownDescription: "payment_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"x_paypal_password": schema.StringAttribute{
				MarkdownDescription: "x_paypal_password field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"x_paypal_signature": schema.StringAttribute{
				MarkdownDescription: "x_paypal_signature field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"paypal_use_sandbox": schema.BoolAttribute{
				MarkdownDescription: "paypal_use_sandbox field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"x_paypal_username": schema.StringAttribute{
				MarkdownDescription: "x_paypal_username field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized": schema.BoolAttribute{
				MarkdownDescription: "portal_customized field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_authentication_text": schema.StringAttribute{
				MarkdownDescription: "portal_customized_authentication_text field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_bg_color": schema.StringAttribute{
				MarkdownDescription: "^#[a-zA-Z0-9]{6}$|^#[a-zA-Z0-9]{3}$|^$",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_bg_image_enabled": schema.BoolAttribute{
				MarkdownDescription: "portal_customized_bg_image_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_bg_image_filename": schema.StringAttribute{
				MarkdownDescription: "portal_customized_bg_image_filename field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_bg_image_tile": schema.BoolAttribute{
				MarkdownDescription: "portal_customized_bg_image_tile field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_bg_type": schema.StringAttribute{
				MarkdownDescription: "color|image|gallery",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_box_color": schema.StringAttribute{
				MarkdownDescription: "^#[a-zA-Z0-9]{6}$|^#[a-zA-Z0-9]{3}$|^$",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_box_link_color": schema.StringAttribute{
				MarkdownDescription: "^#[a-zA-Z0-9]{6}$|^#[a-zA-Z0-9]{3}$|^$",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_box_opacity": schema.Int64Attribute{
				MarkdownDescription: "^[1-9][0-9]?$|^100$|^$",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"portal_customized_box_radius": schema.Int64Attribute{
				MarkdownDescription: "[0-9]|[1-4][0-9]|50",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"portal_customized_box_text_color": schema.StringAttribute{
				MarkdownDescription: "^#[a-zA-Z0-9]{6}$|^#[a-zA-Z0-9]{3}$|^$",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_button_color": schema.StringAttribute{
				MarkdownDescription: "^#[a-zA-Z0-9]{6}$|^#[a-zA-Z0-9]{3}$|^$",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_button_text": schema.StringAttribute{
				MarkdownDescription: "portal_customized_button_text field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_button_text_color": schema.StringAttribute{
				MarkdownDescription: "^#[a-zA-Z0-9]{6}$|^#[a-zA-Z0-9]{3}$|^$",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_languages": schema.ListAttribute{
				MarkdownDescription: "^[a-z]{2}([_-][a-zA-Z]{2,4})*$",
				Optional:            true, Computed: true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_link_color": schema.StringAttribute{
				MarkdownDescription: "^#[a-zA-Z0-9]{6}$|^#[a-zA-Z0-9]{3}$|^$",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_logo_enabled": schema.BoolAttribute{
				MarkdownDescription: "portal_customized_logo_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_logo_filename": schema.StringAttribute{
				MarkdownDescription: "portal_customized_logo_filename field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_logo_position": schema.StringAttribute{
				MarkdownDescription: "left|center|right",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_logo_size": schema.Int64Attribute{
				MarkdownDescription: "6[4-9]|[7-9][0-9]|1[0-8][0-9]|19[0-2]",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"portal_customized_success_text": schema.StringAttribute{
				MarkdownDescription: "portal_customized_success_text field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_text_color": schema.StringAttribute{
				MarkdownDescription: "^#[a-zA-Z0-9]{6}$|^#[a-zA-Z0-9]{3}$|^$",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_title": schema.StringAttribute{
				MarkdownDescription: "portal_customized_title field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_tos": schema.StringAttribute{
				MarkdownDescription: "portal_customized_tos field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_tos_enabled": schema.BoolAttribute{
				MarkdownDescription: "portal_customized_tos_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_unsplash_author_name": schema.StringAttribute{
				MarkdownDescription: "portal_customized_unsplash_author_name field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_unsplash_author_username": schema.StringAttribute{
				MarkdownDescription: "portal_customized_unsplash_author_username field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_welcome_text": schema.StringAttribute{
				MarkdownDescription: "portal_customized_welcome_text field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_welcome_text_enabled": schema.BoolAttribute{
				MarkdownDescription: "portal_customized_welcome_text_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"portal_customized_welcome_text_position": schema.StringAttribute{
				MarkdownDescription: "under_logo|above_boxes",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_enabled": schema.BoolAttribute{
				MarkdownDescription: "portal_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"portal_hostname": schema.StringAttribute{
				MarkdownDescription: "^[a-zA-Z0-9.-]+$|^$",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"portal_use_hostname": schema.BoolAttribute{
				MarkdownDescription: "portal_use_hostname field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"x_quickpay_agreementid": schema.StringAttribute{
				MarkdownDescription: "x_quickpay_agreementid field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"x_quickpay_apikey": schema.StringAttribute{
				MarkdownDescription: "x_quickpay_apikey field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"x_quickpay_merchantid": schema.StringAttribute{
				MarkdownDescription: "x_quickpay_merchantid field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"quickpay_testmode": schema.BoolAttribute{
				MarkdownDescription: "quickpay_testmode field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"radius_auth_type": schema.StringAttribute{
				MarkdownDescription: "chap|mschapv2",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"radius_disconnect_enabled": schema.BoolAttribute{
				MarkdownDescription: "radius_disconnect_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"radius_disconnect_port": schema.Int64Attribute{
				MarkdownDescription: "[1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5]",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"radius_enabled": schema.BoolAttribute{
				MarkdownDescription: "radius_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"radiusprofile_id": schema.StringAttribute{
				MarkdownDescription: "radiusprofile_id field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"redirect_enabled": schema.BoolAttribute{
				MarkdownDescription: "redirect_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"redirect_https": schema.BoolAttribute{
				MarkdownDescription: "redirect_https field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"redirect_to_https": schema.BoolAttribute{
				MarkdownDescription: "redirect_to_https field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"redirect_url": schema.StringAttribute{
				MarkdownDescription: "redirect_url field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"restricted_dns_enabled": schema.BoolAttribute{
				MarkdownDescription: "restricted_dns_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"restricted_dns_servers": schema.ListAttribute{
				MarkdownDescription: "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$",
				Optional:            true, Computed: true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"restricted_subnet_": schema.StringAttribute{
				MarkdownDescription: "restricted_subnet_ field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"x_stripe_api_key": schema.StringAttribute{
				MarkdownDescription: "x_stripe_api_key field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"template_engine": schema.StringAttribute{
				MarkdownDescription: "jsp|angular",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"voucher_customized": schema.BoolAttribute{
				MarkdownDescription: "voucher_customized field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"voucher_enabled": schema.BoolAttribute{
				MarkdownDescription: "voucher_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"wechat_app_id": schema.StringAttribute{
				MarkdownDescription: "wechat_app_id field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"x_wechat_app_secret": schema.StringAttribute{
				MarkdownDescription: "x_wechat_app_secret field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"wechat_enabled": schema.BoolAttribute{
				MarkdownDescription: "wechat_enabled field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"x_wechat_secret_key": schema.StringAttribute{
				MarkdownDescription: "x_wechat_secret_key field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"wechat_shop_id": schema.StringAttribute{
				MarkdownDescription: "wechat_shop_id field",
				Optional:            true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *settingGuestAccessResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *settingGuestAccessResource) siteOrDefault(site types.String) string {
	if s := site.ValueString(); s != "" {
		return s
	}
	return r.client.Site
}

func (r *settingGuestAccessResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan settingGuestAccessModel
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

func (r *settingGuestAccessResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state settingGuestAccessModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	site := r.siteOrDefault(state.Site)

	meta, current, err := ui.GetSetting[*settings.GuestAccess](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); ok {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading GuestAccess Setting", err.Error())
		return
	}
	resp.Diagnostics.Append(r.settingToModel(ctx, meta, current, &state, site)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *settingGuestAccessResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state settingGuestAccessModel
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

func (r *settingGuestAccessResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Settings cannot be deleted from the controller; dropping from state only.
}

func (r *settingGuestAccessResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *settingGuestAccessResource) writeAndRefresh(ctx context.Context, site string, m *settingGuestAccessModel) diag.Diagnostics {
	var diags diag.Diagnostics

	// Read current so unmodified fields aren't dropped to zero values.
	_, current, err := ui.GetSetting[*settings.GuestAccess](r.client.ApiClient, ctx, site)
	if err != nil {
		if _, ok := err.(*ui.NotFoundError); !ok {
			diags.AddError("Error Reading GuestAccess Setting", err.Error())
			return diags
		}
		current = &settings.GuestAccess{}
	}

	r.applyModelToSetting(ctx, m, current, &diags)
	if diags.HasError() {
		return diags
	}

	if err := r.client.UpdateSetting(ctx, site, current); err != nil {
		diags.AddError("Error Updating GuestAccess Setting", err.Error())
		return diags
	}

	meta, refreshed, err := ui.GetSetting[*settings.GuestAccess](r.client.ApiClient, ctx, site)
	if err != nil {
		diags.AddError("Error Re-reading GuestAccess Setting", err.Error())
		return diags
	}
	diags.Append(r.settingToModel(ctx, meta, refreshed, m, site)...)
	return diags
}

func (r *settingGuestAccessResource) settingToModel(ctx context.Context, meta *ui.Setting, s *settings.GuestAccess, m *settingGuestAccessModel, site string) diag.Diagnostics {
	var diags diag.Diagnostics
	if meta != nil {
		m.ID = types.StringValue(meta.Id)
	}
	m.Site = types.StringValue(site)
	_ = ctx
	_ = diags
	m.AllowedSubnet = stringOrNull(s.AllowedSubnet)
	m.Auth = stringOrNull(s.Auth)
	m.AuthUrl = stringOrNull(s.AuthUrl)
	m.AuthorizeLoginid = stringOrNull(s.AuthorizeLoginid)
	m.AuthorizeTransactionkey = stringOrNull(s.AuthorizeTransactionkey)
	m.AuthorizeUseSandbox = types.BoolValue(s.AuthorizeUseSandbox)
	m.CustomIP = stringOrNull(s.CustomIP)
	m.EcEnabled = types.BoolValue(s.EcEnabled)
	m.Expire = stringOrNull(s.Expire)
	m.ExpireNumber = types.Int64PointerValue(s.ExpireNumber)
	m.ExpireUnit = types.Int64PointerValue(s.ExpireUnit)
	m.FacebookAppID = stringOrNull(s.FacebookAppID)
	m.FacebookAppSecret = stringOrNull(s.FacebookAppSecret)
	m.FacebookEnabled = types.BoolValue(s.FacebookEnabled)
	m.FacebookScopeEmail = types.BoolValue(s.FacebookScopeEmail)
	m.FacebookWifiBlockHttps = types.BoolValue(s.FacebookWifiBlockHttps)
	m.FacebookWifiGwID = stringOrNull(s.FacebookWifiGwID)
	m.FacebookWifiGwName = stringOrNull(s.FacebookWifiGwName)
	m.FacebookWifiGwSecret = stringOrNull(s.FacebookWifiGwSecret)
	m.Gateway = stringOrNull(s.Gateway)
	m.GoogleClientID = stringOrNull(s.GoogleClientID)
	m.GoogleClientSecret = stringOrNull(s.GoogleClientSecret)
	m.GoogleDomain = stringOrNull(s.GoogleDomain)
	m.GoogleEnabled = types.BoolValue(s.GoogleEnabled)
	m.GoogleScopeEmail = types.BoolValue(s.GoogleScopeEmail)
	m.IPpayTerminalid = stringOrNull(s.IPpayTerminalid)
	m.IPpayUseSandbox = types.BoolValue(s.IPpayUseSandbox)
	m.MerchantwarriorApikey = stringOrNull(s.MerchantwarriorApikey)
	m.MerchantwarriorApipassphrase = stringOrNull(s.MerchantwarriorApipassphrase)
	m.MerchantwarriorMerchantuuid = stringOrNull(s.MerchantwarriorMerchantuuid)
	m.MerchantwarriorUseSandbox = types.BoolValue(s.MerchantwarriorUseSandbox)
	m.Password = stringOrNull(s.Password)
	m.PasswordEnabled = types.BoolValue(s.PasswordEnabled)
	m.PaymentEnabled = types.BoolValue(s.PaymentEnabled)
	m.PaypalPassword = stringOrNull(s.PaypalPassword)
	m.PaypalSignature = stringOrNull(s.PaypalSignature)
	m.PaypalUseSandbox = types.BoolValue(s.PaypalUseSandbox)
	m.PaypalUsername = stringOrNull(s.PaypalUsername)
	m.PortalCustomized = types.BoolValue(s.PortalCustomized)
	m.PortalCustomizedAuthenticationText = stringOrNull(s.PortalCustomizedAuthenticationText)
	m.PortalCustomizedBgColor = stringOrNull(s.PortalCustomizedBgColor)
	m.PortalCustomizedBgImageEnabled = types.BoolValue(s.PortalCustomizedBgImageEnabled)
	m.PortalCustomizedBgImageFilename = stringOrNull(s.PortalCustomizedBgImageFilename)
	m.PortalCustomizedBgImageTile = types.BoolValue(s.PortalCustomizedBgImageTile)
	m.PortalCustomizedBgType = stringOrNull(s.PortalCustomizedBgType)
	m.PortalCustomizedBoxColor = stringOrNull(s.PortalCustomizedBoxColor)
	m.PortalCustomizedBoxLinkColor = stringOrNull(s.PortalCustomizedBoxLinkColor)
	m.PortalCustomizedBoxOpacity = types.Int64PointerValue(s.PortalCustomizedBoxOpacity)
	m.PortalCustomizedBoxRADIUS = types.Int64PointerValue(s.PortalCustomizedBoxRADIUS)
	m.PortalCustomizedBoxTextColor = stringOrNull(s.PortalCustomizedBoxTextColor)
	m.PortalCustomizedButtonColor = stringOrNull(s.PortalCustomizedButtonColor)
	m.PortalCustomizedButtonText = stringOrNull(s.PortalCustomizedButtonText)
	m.PortalCustomizedButtonTextColor = stringOrNull(s.PortalCustomizedButtonTextColor)
	m.PortalCustomizedLanguages = stringList(ctx, s.PortalCustomizedLanguages, &diags)
	m.PortalCustomizedLinkColor = stringOrNull(s.PortalCustomizedLinkColor)
	m.PortalCustomizedLogoEnabled = types.BoolValue(s.PortalCustomizedLogoEnabled)
	m.PortalCustomizedLogoFilename = stringOrNull(s.PortalCustomizedLogoFilename)
	m.PortalCustomizedLogoPosition = stringOrNull(s.PortalCustomizedLogoPosition)
	m.PortalCustomizedLogoSize = types.Int64PointerValue(s.PortalCustomizedLogoSize)
	m.PortalCustomizedSuccessText = stringOrNull(s.PortalCustomizedSuccessText)
	m.PortalCustomizedTextColor = stringOrNull(s.PortalCustomizedTextColor)
	m.PortalCustomizedTitle = stringOrNull(s.PortalCustomizedTitle)
	m.PortalCustomizedTos = stringOrNull(s.PortalCustomizedTos)
	m.PortalCustomizedTosEnabled = types.BoolValue(s.PortalCustomizedTosEnabled)
	m.PortalCustomizedUnsplashAuthorName = stringOrNull(s.PortalCustomizedUnsplashAuthorName)
	m.PortalCustomizedUnsplashAuthorUsername = stringOrNull(s.PortalCustomizedUnsplashAuthorUsername)
	m.PortalCustomizedWelcomeText = stringOrNull(s.PortalCustomizedWelcomeText)
	m.PortalCustomizedWelcomeTextEnabled = types.BoolValue(s.PortalCustomizedWelcomeTextEnabled)
	m.PortalCustomizedWelcomeTextPosition = stringOrNull(s.PortalCustomizedWelcomeTextPosition)
	m.PortalEnabled = types.BoolValue(s.PortalEnabled)
	m.PortalHostname = stringOrNull(s.PortalHostname)
	m.PortalUseHostname = types.BoolValue(s.PortalUseHostname)
	m.QuickpayAgreementid = stringOrNull(s.QuickpayAgreementid)
	m.QuickpayApikey = stringOrNull(s.QuickpayApikey)
	m.QuickpayMerchantid = stringOrNull(s.QuickpayMerchantid)
	m.QuickpayTestmode = types.BoolValue(s.QuickpayTestmode)
	m.RADIUSAuthType = stringOrNull(s.RADIUSAuthType)
	m.RADIUSDisconnectEnabled = types.BoolValue(s.RADIUSDisconnectEnabled)
	m.RADIUSDisconnectPort = types.Int64PointerValue(s.RADIUSDisconnectPort)
	m.RADIUSEnabled = types.BoolValue(s.RADIUSEnabled)
	m.RADIUSProfileID = stringOrNull(s.RADIUSProfileID)
	m.RedirectEnabled = types.BoolValue(s.RedirectEnabled)
	m.RedirectHttps = types.BoolValue(s.RedirectHttps)
	m.RedirectToHttps = types.BoolValue(s.RedirectToHttps)
	m.RedirectUrl = stringOrNull(s.RedirectUrl)
	m.RestrictedDNSEnabled = types.BoolValue(s.RestrictedDNSEnabled)
	m.RestrictedDNSServers = stringList(ctx, s.RestrictedDNSServers, &diags)
	m.RestrictedSubnet = stringOrNull(s.RestrictedSubnet)
	m.StripeApiKey = stringOrNull(s.StripeApiKey)
	m.TemplateEngine = stringOrNull(s.TemplateEngine)
	m.VoucherCustomized = types.BoolValue(s.VoucherCustomized)
	m.VoucherEnabled = types.BoolValue(s.VoucherEnabled)
	m.WechatAppID = stringOrNull(s.WechatAppID)
	m.WechatAppSecret = stringOrNull(s.WechatAppSecret)
	m.WechatEnabled = types.BoolValue(s.WechatEnabled)
	m.WechatSecretKey = stringOrNull(s.WechatSecretKey)
	m.WechatShopID = stringOrNull(s.WechatShopID)
	return diags
}

func (r *settingGuestAccessResource) applyModelToSetting(ctx context.Context, m *settingGuestAccessModel, s *settings.GuestAccess, diags *diag.Diagnostics) {
	_ = ctx
	_ = diags
	if !m.AllowedSubnet.IsNull() && !m.AllowedSubnet.IsUnknown() {
		s.AllowedSubnet = m.AllowedSubnet.ValueString()
	}
	if !m.Auth.IsNull() && !m.Auth.IsUnknown() {
		s.Auth = m.Auth.ValueString()
	}
	if !m.AuthUrl.IsNull() && !m.AuthUrl.IsUnknown() {
		s.AuthUrl = m.AuthUrl.ValueString()
	}
	if !m.AuthorizeLoginid.IsNull() && !m.AuthorizeLoginid.IsUnknown() {
		s.AuthorizeLoginid = m.AuthorizeLoginid.ValueString()
	}
	if !m.AuthorizeTransactionkey.IsNull() && !m.AuthorizeTransactionkey.IsUnknown() {
		s.AuthorizeTransactionkey = m.AuthorizeTransactionkey.ValueString()
	}
	if !m.AuthorizeUseSandbox.IsNull() && !m.AuthorizeUseSandbox.IsUnknown() {
		s.AuthorizeUseSandbox = m.AuthorizeUseSandbox.ValueBool()
	}
	if !m.CustomIP.IsNull() && !m.CustomIP.IsUnknown() {
		s.CustomIP = m.CustomIP.ValueString()
	}
	if !m.EcEnabled.IsNull() && !m.EcEnabled.IsUnknown() {
		s.EcEnabled = m.EcEnabled.ValueBool()
	}
	if !m.Expire.IsNull() && !m.Expire.IsUnknown() {
		s.Expire = m.Expire.ValueString()
	}
	if !m.ExpireNumber.IsNull() && !m.ExpireNumber.IsUnknown() {
		s.ExpireNumber = int64PointerOrNil(m.ExpireNumber)
	}
	if !m.ExpireUnit.IsNull() && !m.ExpireUnit.IsUnknown() {
		s.ExpireUnit = int64PointerOrNil(m.ExpireUnit)
	}
	if !m.FacebookAppID.IsNull() && !m.FacebookAppID.IsUnknown() {
		s.FacebookAppID = m.FacebookAppID.ValueString()
	}
	if !m.FacebookAppSecret.IsNull() && !m.FacebookAppSecret.IsUnknown() {
		s.FacebookAppSecret = m.FacebookAppSecret.ValueString()
	}
	if !m.FacebookEnabled.IsNull() && !m.FacebookEnabled.IsUnknown() {
		s.FacebookEnabled = m.FacebookEnabled.ValueBool()
	}
	if !m.FacebookScopeEmail.IsNull() && !m.FacebookScopeEmail.IsUnknown() {
		s.FacebookScopeEmail = m.FacebookScopeEmail.ValueBool()
	}
	if !m.FacebookWifiBlockHttps.IsNull() && !m.FacebookWifiBlockHttps.IsUnknown() {
		s.FacebookWifiBlockHttps = m.FacebookWifiBlockHttps.ValueBool()
	}
	if !m.FacebookWifiGwID.IsNull() && !m.FacebookWifiGwID.IsUnknown() {
		s.FacebookWifiGwID = m.FacebookWifiGwID.ValueString()
	}
	if !m.FacebookWifiGwName.IsNull() && !m.FacebookWifiGwName.IsUnknown() {
		s.FacebookWifiGwName = m.FacebookWifiGwName.ValueString()
	}
	if !m.FacebookWifiGwSecret.IsNull() && !m.FacebookWifiGwSecret.IsUnknown() {
		s.FacebookWifiGwSecret = m.FacebookWifiGwSecret.ValueString()
	}
	if !m.Gateway.IsNull() && !m.Gateway.IsUnknown() {
		s.Gateway = m.Gateway.ValueString()
	}
	if !m.GoogleClientID.IsNull() && !m.GoogleClientID.IsUnknown() {
		s.GoogleClientID = m.GoogleClientID.ValueString()
	}
	if !m.GoogleClientSecret.IsNull() && !m.GoogleClientSecret.IsUnknown() {
		s.GoogleClientSecret = m.GoogleClientSecret.ValueString()
	}
	if !m.GoogleDomain.IsNull() && !m.GoogleDomain.IsUnknown() {
		s.GoogleDomain = m.GoogleDomain.ValueString()
	}
	if !m.GoogleEnabled.IsNull() && !m.GoogleEnabled.IsUnknown() {
		s.GoogleEnabled = m.GoogleEnabled.ValueBool()
	}
	if !m.GoogleScopeEmail.IsNull() && !m.GoogleScopeEmail.IsUnknown() {
		s.GoogleScopeEmail = m.GoogleScopeEmail.ValueBool()
	}
	if !m.IPpayTerminalid.IsNull() && !m.IPpayTerminalid.IsUnknown() {
		s.IPpayTerminalid = m.IPpayTerminalid.ValueString()
	}
	if !m.IPpayUseSandbox.IsNull() && !m.IPpayUseSandbox.IsUnknown() {
		s.IPpayUseSandbox = m.IPpayUseSandbox.ValueBool()
	}
	if !m.MerchantwarriorApikey.IsNull() && !m.MerchantwarriorApikey.IsUnknown() {
		s.MerchantwarriorApikey = m.MerchantwarriorApikey.ValueString()
	}
	if !m.MerchantwarriorApipassphrase.IsNull() && !m.MerchantwarriorApipassphrase.IsUnknown() {
		s.MerchantwarriorApipassphrase = m.MerchantwarriorApipassphrase.ValueString()
	}
	if !m.MerchantwarriorMerchantuuid.IsNull() && !m.MerchantwarriorMerchantuuid.IsUnknown() {
		s.MerchantwarriorMerchantuuid = m.MerchantwarriorMerchantuuid.ValueString()
	}
	if !m.MerchantwarriorUseSandbox.IsNull() && !m.MerchantwarriorUseSandbox.IsUnknown() {
		s.MerchantwarriorUseSandbox = m.MerchantwarriorUseSandbox.ValueBool()
	}
	if !m.Password.IsNull() && !m.Password.IsUnknown() {
		s.Password = m.Password.ValueString()
	}
	if !m.PasswordEnabled.IsNull() && !m.PasswordEnabled.IsUnknown() {
		s.PasswordEnabled = m.PasswordEnabled.ValueBool()
	}
	if !m.PaymentEnabled.IsNull() && !m.PaymentEnabled.IsUnknown() {
		s.PaymentEnabled = m.PaymentEnabled.ValueBool()
	}
	if !m.PaypalPassword.IsNull() && !m.PaypalPassword.IsUnknown() {
		s.PaypalPassword = m.PaypalPassword.ValueString()
	}
	if !m.PaypalSignature.IsNull() && !m.PaypalSignature.IsUnknown() {
		s.PaypalSignature = m.PaypalSignature.ValueString()
	}
	if !m.PaypalUseSandbox.IsNull() && !m.PaypalUseSandbox.IsUnknown() {
		s.PaypalUseSandbox = m.PaypalUseSandbox.ValueBool()
	}
	if !m.PaypalUsername.IsNull() && !m.PaypalUsername.IsUnknown() {
		s.PaypalUsername = m.PaypalUsername.ValueString()
	}
	if !m.PortalCustomized.IsNull() && !m.PortalCustomized.IsUnknown() {
		s.PortalCustomized = m.PortalCustomized.ValueBool()
	}
	if !m.PortalCustomizedAuthenticationText.IsNull() && !m.PortalCustomizedAuthenticationText.IsUnknown() {
		s.PortalCustomizedAuthenticationText = m.PortalCustomizedAuthenticationText.ValueString()
	}
	if !m.PortalCustomizedBgColor.IsNull() && !m.PortalCustomizedBgColor.IsUnknown() {
		s.PortalCustomizedBgColor = m.PortalCustomizedBgColor.ValueString()
	}
	if !m.PortalCustomizedBgImageEnabled.IsNull() && !m.PortalCustomizedBgImageEnabled.IsUnknown() {
		s.PortalCustomizedBgImageEnabled = m.PortalCustomizedBgImageEnabled.ValueBool()
	}
	if !m.PortalCustomizedBgImageFilename.IsNull() && !m.PortalCustomizedBgImageFilename.IsUnknown() {
		s.PortalCustomizedBgImageFilename = m.PortalCustomizedBgImageFilename.ValueString()
	}
	if !m.PortalCustomizedBgImageTile.IsNull() && !m.PortalCustomizedBgImageTile.IsUnknown() {
		s.PortalCustomizedBgImageTile = m.PortalCustomizedBgImageTile.ValueBool()
	}
	if !m.PortalCustomizedBgType.IsNull() && !m.PortalCustomizedBgType.IsUnknown() {
		s.PortalCustomizedBgType = m.PortalCustomizedBgType.ValueString()
	}
	if !m.PortalCustomizedBoxColor.IsNull() && !m.PortalCustomizedBoxColor.IsUnknown() {
		s.PortalCustomizedBoxColor = m.PortalCustomizedBoxColor.ValueString()
	}
	if !m.PortalCustomizedBoxLinkColor.IsNull() && !m.PortalCustomizedBoxLinkColor.IsUnknown() {
		s.PortalCustomizedBoxLinkColor = m.PortalCustomizedBoxLinkColor.ValueString()
	}
	if !m.PortalCustomizedBoxOpacity.IsNull() && !m.PortalCustomizedBoxOpacity.IsUnknown() {
		s.PortalCustomizedBoxOpacity = int64PointerOrNil(m.PortalCustomizedBoxOpacity)
	}
	if !m.PortalCustomizedBoxRADIUS.IsNull() && !m.PortalCustomizedBoxRADIUS.IsUnknown() {
		s.PortalCustomizedBoxRADIUS = int64PointerOrNil(m.PortalCustomizedBoxRADIUS)
	}
	if !m.PortalCustomizedBoxTextColor.IsNull() && !m.PortalCustomizedBoxTextColor.IsUnknown() {
		s.PortalCustomizedBoxTextColor = m.PortalCustomizedBoxTextColor.ValueString()
	}
	if !m.PortalCustomizedButtonColor.IsNull() && !m.PortalCustomizedButtonColor.IsUnknown() {
		s.PortalCustomizedButtonColor = m.PortalCustomizedButtonColor.ValueString()
	}
	if !m.PortalCustomizedButtonText.IsNull() && !m.PortalCustomizedButtonText.IsUnknown() {
		s.PortalCustomizedButtonText = m.PortalCustomizedButtonText.ValueString()
	}
	if !m.PortalCustomizedButtonTextColor.IsNull() && !m.PortalCustomizedButtonTextColor.IsUnknown() {
		s.PortalCustomizedButtonTextColor = m.PortalCustomizedButtonTextColor.ValueString()
	}
	if !m.PortalCustomizedLanguages.IsNull() && !m.PortalCustomizedLanguages.IsUnknown() {
		var v []string
		diags.Append(m.PortalCustomizedLanguages.ElementsAs(ctx, &v, false)...)
		s.PortalCustomizedLanguages = v
	}
	if !m.PortalCustomizedLinkColor.IsNull() && !m.PortalCustomizedLinkColor.IsUnknown() {
		s.PortalCustomizedLinkColor = m.PortalCustomizedLinkColor.ValueString()
	}
	if !m.PortalCustomizedLogoEnabled.IsNull() && !m.PortalCustomizedLogoEnabled.IsUnknown() {
		s.PortalCustomizedLogoEnabled = m.PortalCustomizedLogoEnabled.ValueBool()
	}
	if !m.PortalCustomizedLogoFilename.IsNull() && !m.PortalCustomizedLogoFilename.IsUnknown() {
		s.PortalCustomizedLogoFilename = m.PortalCustomizedLogoFilename.ValueString()
	}
	if !m.PortalCustomizedLogoPosition.IsNull() && !m.PortalCustomizedLogoPosition.IsUnknown() {
		s.PortalCustomizedLogoPosition = m.PortalCustomizedLogoPosition.ValueString()
	}
	if !m.PortalCustomizedLogoSize.IsNull() && !m.PortalCustomizedLogoSize.IsUnknown() {
		s.PortalCustomizedLogoSize = int64PointerOrNil(m.PortalCustomizedLogoSize)
	}
	if !m.PortalCustomizedSuccessText.IsNull() && !m.PortalCustomizedSuccessText.IsUnknown() {
		s.PortalCustomizedSuccessText = m.PortalCustomizedSuccessText.ValueString()
	}
	if !m.PortalCustomizedTextColor.IsNull() && !m.PortalCustomizedTextColor.IsUnknown() {
		s.PortalCustomizedTextColor = m.PortalCustomizedTextColor.ValueString()
	}
	if !m.PortalCustomizedTitle.IsNull() && !m.PortalCustomizedTitle.IsUnknown() {
		s.PortalCustomizedTitle = m.PortalCustomizedTitle.ValueString()
	}
	if !m.PortalCustomizedTos.IsNull() && !m.PortalCustomizedTos.IsUnknown() {
		s.PortalCustomizedTos = m.PortalCustomizedTos.ValueString()
	}
	if !m.PortalCustomizedTosEnabled.IsNull() && !m.PortalCustomizedTosEnabled.IsUnknown() {
		s.PortalCustomizedTosEnabled = m.PortalCustomizedTosEnabled.ValueBool()
	}
	if !m.PortalCustomizedUnsplashAuthorName.IsNull() && !m.PortalCustomizedUnsplashAuthorName.IsUnknown() {
		s.PortalCustomizedUnsplashAuthorName = m.PortalCustomizedUnsplashAuthorName.ValueString()
	}
	if !m.PortalCustomizedUnsplashAuthorUsername.IsNull() && !m.PortalCustomizedUnsplashAuthorUsername.IsUnknown() {
		s.PortalCustomizedUnsplashAuthorUsername = m.PortalCustomizedUnsplashAuthorUsername.ValueString()
	}
	if !m.PortalCustomizedWelcomeText.IsNull() && !m.PortalCustomizedWelcomeText.IsUnknown() {
		s.PortalCustomizedWelcomeText = m.PortalCustomizedWelcomeText.ValueString()
	}
	if !m.PortalCustomizedWelcomeTextEnabled.IsNull() && !m.PortalCustomizedWelcomeTextEnabled.IsUnknown() {
		s.PortalCustomizedWelcomeTextEnabled = m.PortalCustomizedWelcomeTextEnabled.ValueBool()
	}
	if !m.PortalCustomizedWelcomeTextPosition.IsNull() && !m.PortalCustomizedWelcomeTextPosition.IsUnknown() {
		s.PortalCustomizedWelcomeTextPosition = m.PortalCustomizedWelcomeTextPosition.ValueString()
	}
	if !m.PortalEnabled.IsNull() && !m.PortalEnabled.IsUnknown() {
		s.PortalEnabled = m.PortalEnabled.ValueBool()
	}
	if !m.PortalHostname.IsNull() && !m.PortalHostname.IsUnknown() {
		s.PortalHostname = m.PortalHostname.ValueString()
	}
	if !m.PortalUseHostname.IsNull() && !m.PortalUseHostname.IsUnknown() {
		s.PortalUseHostname = m.PortalUseHostname.ValueBool()
	}
	if !m.QuickpayAgreementid.IsNull() && !m.QuickpayAgreementid.IsUnknown() {
		s.QuickpayAgreementid = m.QuickpayAgreementid.ValueString()
	}
	if !m.QuickpayApikey.IsNull() && !m.QuickpayApikey.IsUnknown() {
		s.QuickpayApikey = m.QuickpayApikey.ValueString()
	}
	if !m.QuickpayMerchantid.IsNull() && !m.QuickpayMerchantid.IsUnknown() {
		s.QuickpayMerchantid = m.QuickpayMerchantid.ValueString()
	}
	if !m.QuickpayTestmode.IsNull() && !m.QuickpayTestmode.IsUnknown() {
		s.QuickpayTestmode = m.QuickpayTestmode.ValueBool()
	}
	if !m.RADIUSAuthType.IsNull() && !m.RADIUSAuthType.IsUnknown() {
		s.RADIUSAuthType = m.RADIUSAuthType.ValueString()
	}
	if !m.RADIUSDisconnectEnabled.IsNull() && !m.RADIUSDisconnectEnabled.IsUnknown() {
		s.RADIUSDisconnectEnabled = m.RADIUSDisconnectEnabled.ValueBool()
	}
	if !m.RADIUSDisconnectPort.IsNull() && !m.RADIUSDisconnectPort.IsUnknown() {
		s.RADIUSDisconnectPort = int64PointerOrNil(m.RADIUSDisconnectPort)
	}
	if !m.RADIUSEnabled.IsNull() && !m.RADIUSEnabled.IsUnknown() {
		s.RADIUSEnabled = m.RADIUSEnabled.ValueBool()
	}
	if !m.RADIUSProfileID.IsNull() && !m.RADIUSProfileID.IsUnknown() {
		s.RADIUSProfileID = m.RADIUSProfileID.ValueString()
	}
	if !m.RedirectEnabled.IsNull() && !m.RedirectEnabled.IsUnknown() {
		s.RedirectEnabled = m.RedirectEnabled.ValueBool()
	}
	if !m.RedirectHttps.IsNull() && !m.RedirectHttps.IsUnknown() {
		s.RedirectHttps = m.RedirectHttps.ValueBool()
	}
	if !m.RedirectToHttps.IsNull() && !m.RedirectToHttps.IsUnknown() {
		s.RedirectToHttps = m.RedirectToHttps.ValueBool()
	}
	if !m.RedirectUrl.IsNull() && !m.RedirectUrl.IsUnknown() {
		s.RedirectUrl = m.RedirectUrl.ValueString()
	}
	if !m.RestrictedDNSEnabled.IsNull() && !m.RestrictedDNSEnabled.IsUnknown() {
		s.RestrictedDNSEnabled = m.RestrictedDNSEnabled.ValueBool()
	}
	if !m.RestrictedDNSServers.IsNull() && !m.RestrictedDNSServers.IsUnknown() {
		var v []string
		diags.Append(m.RestrictedDNSServers.ElementsAs(ctx, &v, false)...)
		s.RestrictedDNSServers = v
	}
	if !m.RestrictedSubnet.IsNull() && !m.RestrictedSubnet.IsUnknown() {
		s.RestrictedSubnet = m.RestrictedSubnet.ValueString()
	}
	if !m.StripeApiKey.IsNull() && !m.StripeApiKey.IsUnknown() {
		s.StripeApiKey = m.StripeApiKey.ValueString()
	}
	if !m.TemplateEngine.IsNull() && !m.TemplateEngine.IsUnknown() {
		s.TemplateEngine = m.TemplateEngine.ValueString()
	}
	if !m.VoucherCustomized.IsNull() && !m.VoucherCustomized.IsUnknown() {
		s.VoucherCustomized = m.VoucherCustomized.ValueBool()
	}
	if !m.VoucherEnabled.IsNull() && !m.VoucherEnabled.IsUnknown() {
		s.VoucherEnabled = m.VoucherEnabled.ValueBool()
	}
	if !m.WechatAppID.IsNull() && !m.WechatAppID.IsUnknown() {
		s.WechatAppID = m.WechatAppID.ValueString()
	}
	if !m.WechatAppSecret.IsNull() && !m.WechatAppSecret.IsUnknown() {
		s.WechatAppSecret = m.WechatAppSecret.ValueString()
	}
	if !m.WechatEnabled.IsNull() && !m.WechatEnabled.IsUnknown() {
		s.WechatEnabled = m.WechatEnabled.ValueBool()
	}
	if !m.WechatSecretKey.IsNull() && !m.WechatSecretKey.IsUnknown() {
		s.WechatSecretKey = m.WechatSecretKey.ValueString()
	}
	if !m.WechatShopID.IsNull() && !m.WechatShopID.IsUnknown() {
		s.WechatShopID = m.WechatShopID.ValueString()
	}
}
