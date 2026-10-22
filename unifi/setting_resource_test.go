package unifi

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
