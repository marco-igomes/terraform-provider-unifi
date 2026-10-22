package unifi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	ui "github.com/ubiquiti-community/go-unifi/unifi"
	"github.com/ubiquiti-community/go-unifi/unifi/settings"
)

func TestAccSettingResource_mgmt(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { preCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSettingConfig_mgmt(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"mgmt.auto_upgrade",
						"true",
					),
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"mgmt.ssh_enabled",
						"false",
					),
				),
			},
			{
				ResourceName:      "unifi_setting.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"mgmt.%",
					"mgmt.auto_upgrade",
					"mgmt.ssh_enabled",
				},
			},
		},
	})
}

func TestAccSettingResource_radius(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { preCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSettingConfig_radius(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"radius.accounting_enabled",
						"true",
					),
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"radius.auth_port",
						"1812",
					),
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"radius.acct_port",
						"1813",
					),
				),
			},
			{
				ResourceName:      "unifi_setting.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"radius.secret", // Secret is sensitive and won't be in state after import
					"radius.%",
					"radius.accounting_enabled",
					"radius.acct_port",
					"radius.auth_port",
					"radius.interim_update_interval",
				},
			},
		},
	})
}

func TestAccSettingResource_usg(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { preCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSettingConfig_usg(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"usg.ftp_module",
						"true",
					),
				),
			},
			{
				ResourceName:      "unifi_setting.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"usg.%",
					"usg.ftp_module",
				},
			},
		},
	})
}

func TestAccSettingResource_combined(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { preCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSettingConfig_combined(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"mgmt.auto_upgrade",
						"true",
					),
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"mgmt.ssh_enabled",
						"true",
					),
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"radius.accounting_enabled",
						"false",
					),
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"usg.ftp_module",
						"false",
					),
				),
			},
			{
				Config: testAccSettingConfig_combinedUpdate(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"mgmt.auto_upgrade",
						"false",
					),
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"mgmt.ssh_enabled",
						"false",
					),
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"radius.accounting_enabled",
						"true",
					),
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"usg.ftp_module",
						"true",
					),
				),
			},
		},
	})
}

func TestAccSettingResource_sshKeys(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { preCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSettingConfig_sshKeys(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"mgmt.ssh_enabled",
						"true",
					),
					resource.TestCheckResourceAttr("unifi_setting.test", "mgmt.ssh_keys.#", "1"),
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"mgmt.ssh_keys.0.name",
						"test-key",
					),
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"mgmt.ssh_keys.0.type",
						"ssh-rsa",
					),
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"mgmt.ssh_keys.0.comment",
						"Test SSH Key",
					),
				),
			},
			{
				Config: testAccSettingConfig_sshKeysUpdate(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("unifi_setting.test", "mgmt.ssh_keys.#", "2"),
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"mgmt.ssh_keys.0.name",
						"test-key",
					),
					resource.TestCheckResourceAttr(
						"unifi_setting.test",
						"mgmt.ssh_keys.1.name",
						"test-key-2",
					),
				),
			},
		},
	})
}

func testAccSettingConfig_mgmt() string {
	return `
resource "unifi_setting" "test" {
  mgmt = {
    auto_upgrade = true
    ssh_enabled  = false
  }
}
`
}

func testAccSettingConfig_radius() string {
	return `
resource "unifi_setting" "test" {
  radius = {
    accounting_enabled = true
    auth_port          = 1812
    acct_port          = 1813
    secret             = "test-secret-123"
  }
}
`
}

func testAccSettingConfig_usg() string {
	return `
resource "unifi_setting" "test" {
  usg = {
    ftp_module = true
  }
}
`
}

func testAccSettingConfig_combined() string {
	return `
resource "unifi_setting" "test" {
  mgmt = {
    auto_upgrade = true
    ssh_enabled  = true
  }

  radius = {
    accounting_enabled = false
  }

  usg = {
    ftp_module = false
  }
}
`
}

func testAccSettingConfig_combinedUpdate() string {
	return `
resource "unifi_setting" "test" {
  mgmt = {
    auto_upgrade = false
    ssh_enabled  = false
  }

  radius = {
    accounting_enabled = true
  }

  usg = {
    ftp_module = true
  }
}
`
}

func testAccSettingConfig_sshKeys() string {
	return `
resource "unifi_setting" "test" {
  mgmt = {
    ssh_enabled = true
    ssh_keys = [{
      name    = "test-key"
      type    = "ssh-rsa"
      key     = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDTest123"
      comment = "Test SSH Key"
    }]
  }
}
`
}

func testAccSettingConfig_sshKeysUpdate() string {
	return `
resource "unifi_setting" "test" {
  mgmt = {
    ssh_enabled = true
    ssh_keys = [
      {
        name    = "test-key"
        type    = "ssh-rsa"
        key     = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDTest123"
        comment = "Test SSH Key"
      },
      {
        name    = "test-key-2"
        type    = "ssh-ed25519"
        key     = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAITest456"
        comment = "Second Test Key"
      }
    ]
  }
}
`
}

func TestSettingMgmtChangedFields(t *testing.T) {
	ctx := context.Background()
	current := &settings.Mgmt{}
	if err := json.Unmarshal([]byte(`{
		"key": "mgmt",
		"auto_upgrade": true,
		"auto_upgrade_hour": 5,
		"led_enabled": true,
		"x_ssh_auth_password_enabled": true,
		"x_ssh_enabled": true,
		"x_ssh_username": "admin",
		"x_ssh_keys": [{"name": "laptop", "type": "ssh-rsa", "key": "AAAA", "comment": "laptop",
			"date": "2024-01-09T00:55:35Z", "fingerprint": "6f:5f:f4:52"}]
	}`), current); err != nil {
		t.Fatal(err)
	}
	sshKeys, diags := types.ListValueFrom(ctx, types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"type":    types.StringType,
			"key":     types.StringType,
			"comment": types.StringType,
		},
	}, []sshKeyModel{{
		Name:    types.StringValue("laptop"),
		Type:    types.StringValue("ssh-rsa"),
		Key:     types.StringValue("AAAA"),
		Comment: types.StringValue("laptop"),
	}})
	if diags.HasError() {
		t.Fatal(diags)
	}
	model := &settingMgmtModel{
		AutoUpgrade: types.BoolValue(false),
		SSHEnabled:  types.BoolUnknown(),
		SSHKeys:     sshKeys,
	}

	changed, err := changedFields(current, (&settingResource{}).mgmtModelToSetting(ctx, model, current))
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]any{"auto_upgrade": false}
	if !reflect.DeepEqual(changed, want) {
		t.Errorf("changed = %v, want %v", changed, want)
	}
}

func TestSettingUSGChangedFields(t *testing.T) {
	ctx := context.Background()
	current := &settings.Usg{}
	if err := json.Unmarshal([]byte(`{
		"key": "usg",
		"dhcp_relay_hop_count": "",
		"dhcp_relay_port": "",
		"ftp_module": true,
		"lldp_enable_all": true,
		"upnp_enabled": true,
		"dns_verification": {"domain": "ui.com", "primary_dns_server": "1.1.1.1",
			"secondary_dns_server": "8.8.8.8", "setting_preference": "auto"}
	}`), current); err != nil {
		t.Fatal(err)
	}
	dnsVerification, diags := types.ObjectValueFrom(ctx, map[string]attr.Type{
		"domain":               types.StringType,
		"primary_dns_server":   types.StringType,
		"secondary_dns_server": types.StringType,
		"setting_preference":   types.StringType,
	}, dnsVerificationModel{
		Domain:             types.StringValue("ui.com"),
		PrimaryDNSServer:   types.StringValue("9.9.9.9"),
		SecondaryDNSServer: types.StringUnknown(),
		SettingPreference:  types.StringNull(),
	})
	if diags.HasError() {
		t.Fatal(diags)
	}
	model := &settingUSGModel{
		FtpModule:       types.BoolValue(false),
		UPnPEnabled:     types.BoolUnknown(),
		DNSVerification: dnsVerification,
	}

	changed, err := changedFields(current, (&settingResource{}).usgModelToSetting(ctx, model, current))
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]any{
		"ftp_module": false,
		"dns_verification": map[string]any{
			"domain":               "ui.com",
			"primary_dns_server":   "9.9.9.9",
			"secondary_dns_server": "8.8.8.8",
			"setting_preference":   "auto",
		},
	}
	if !reflect.DeepEqual(changed, want) {
		t.Errorf("changed = %v, want %v", changed, want)
	}
}

func TestSettingRadiusChangedFields(t *testing.T) {
	ctx := context.Background()
	current := &settings.Radius{}
	if err := json.Unmarshal([]byte(`{
		"key": "radius",
		"accounting_enabled": false,
		"auth_port": 1812,
		"acct_port": 1813,
		"x_secret": "secret"
	}`), current); err != nil {
		t.Fatal(err)
	}
	model := &settingRadiusModel{
		AccountingEnabled: types.BoolValue(false),
		AuthPort:          types.Int64Value(1812),
		AcctPort:          types.Int64Value(1814),
	}

	changed, err := changedFields(current, (&settingResource{}).radiusModelToSetting(ctx, model, current))
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]any{"acct_port": float64(1814)}
	if !reflect.DeepEqual(changed, want) {
		t.Errorf("changed = %v, want %v", changed, want)
	}
}

func TestSettingGeneratedChangedFields(t *testing.T) {
	ctx := context.Background()
	site := "default"
	live := map[string]string{
		"country":      `{"key": "country", "code": "620"}`,
		"guest_access": `{"key": "guest_access", "auth": "none", "expire": 1440}`,
		"radio_ai": `{"key": "radio_ai", "enabled": false,
			"channels_ng": ["1", "6", "11"], "ht_modes_ng": ["20"]}`,
		"rsyslogd": `{"key": "rsyslogd", "enabled": true, "debug": false,
			"port": "514", "netconsole_port": "514"}`,
		"connectivity": `{"key": "connectivity", "enabled": true, "uplink_type": "gateway"}`,
		"ips": `{"key": "ips", "ips_mode": "ips", "dns_filters": [{"name": "", "description": "",
			"filter": "none", "network_id": "lan", "version": "v4",
			"allowed_sites": [], "blocked_sites": [], "blocked_tld": []}]}`,
	}
	var sent []map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Error(err)
			}
			sent = append(sent, body)
		}
		fmt.Fprintf(w, `{"meta": {"rc": "ok"}, "data": [%s]}`, live[path.Base(r.URL.Path)])
	}))
	defer server.Close()
	apiClient, err := ui.New(ctx, &ui.Config{BaseURL: server.URL, APIKey: "test"})
	if err != nil {
		t.Fatal(err)
	}
	client := &Client{ApiClient: apiClient, Site: site}

	// Each plan is the state that a read of live gives, with at most one field changed.
	tests := []struct {
		name  string
		write func() diag.Diagnostics
		want  []map[string]any
	}{
		{
			name: "country unchanged",
			write: func() diag.Diagnostics {
				r := &settingCountryResource{client: client}
				var plan settingCountryModel
				current := decodeSetting[settings.Country](t, live["country"])
				r.settingToModel(ctx, nil, current, &plan, site)
				return r.writeAndRefresh(ctx, site, &plan)
			},
		},
		{
			name: "guest_access unchanged",
			write: func() diag.Diagnostics {
				r := &settingGuestAccessResource{client: client}
				var plan settingGuestAccessModel
				current := decodeSetting[settings.GuestAccess](t, live["guest_access"])
				r.settingToModel(ctx, nil, current, &plan, site)
				return r.writeAndRefresh(ctx, site, &plan)
			},
		},
		{
			name: "radio_ai unchanged",
			write: func() diag.Diagnostics {
				r := &settingRadioAiResource{client: client}
				var plan settingRadioAiModel
				current := decodeSetting[settings.RadioAi](t, live["radio_ai"])
				r.settingToModel(ctx, nil, current, &plan, site)
				return r.writeAndRefresh(ctx, site, &plan)
			},
		},
		{
			name: "rsyslogd debug changed",
			write: func() diag.Diagnostics {
				r := &settingRsyslogdResource{client: client}
				var plan settingRsyslogdModel
				current := decodeSetting[settings.Rsyslogd](t, live["rsyslogd"])
				r.settingToModel(ctx, nil, current, &plan, site)
				plan.Debug = types.BoolValue(true)
				return r.writeAndRefresh(ctx, site, &plan)
			},
			want: []map[string]any{{"key": "rsyslogd", "debug": true}},
		},
		{
			name: "connectivity enabled changed",
			write: func() diag.Diagnostics {
				r := &settingConnectivityResource{client: client}
				var plan settingConnectivityModel
				current := decodeSetting[settings.Connectivity](t, live["connectivity"])
				r.settingToModel(ctx, nil, current, &plan, site)
				plan.Enabled = types.BoolValue(false)
				return r.writeAndRefresh(ctx, site, &plan)
			},
			want: []map[string]any{{"key": "connectivity", "enabled": false}},
		},
		{
			name: "ips ips_mode changed",
			write: func() diag.Diagnostics {
				r := &settingIpsResource{client: client}
				var plan settingIpsModel
				current := decodeSetting[settings.Ips](t, live["ips"])
				r.settingToModel(ctx, nil, current, &plan, site)
				plan.IPsMode = types.StringValue("ids")
				return r.writeAndRefresh(ctx, site, &plan)
			},
			want: []map[string]any{{"key": "ips", "ips_mode": "ids"}},
		},
	}
	for _, test := range tests {
		sent = nil
		if diags := test.write(); diags.HasError() {
			t.Errorf("%s: %v", test.name, diags)
		}
		if !reflect.DeepEqual(sent, test.want) {
			t.Errorf("%s: sent %v, want %v", test.name, sent, test.want)
		}
	}
}

func decodeSetting[T any](t *testing.T, raw string) *T {
	t.Helper()
	var setting T
	if err := json.Unmarshal([]byte(raw), &setting); err != nil {
		t.Fatal(err)
	}
	return &setting
}
