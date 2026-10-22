package unifi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/ubiquiti-community/go-unifi/unifi"
)

func TestAccDeviceFramework_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { preCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeviceFrameworkConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("unifi_device.test", "id"),
					resource.TestCheckResourceAttr("unifi_device.test", "name", "Test Device"),
					resource.TestCheckResourceAttr("unifi_device.test", "adopted", "true"),
				),
			},
			{
				ResourceName:            "unifi_device.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"allow_adoption", "forget_on_destroy"},
			},
		},
	})
}

func testAccDeviceFrameworkConfig_basic() string {
	return `
resource "unifi_device" "test" {
	mac  = "00:27:22:00:00:01"
	name = "Test Device"
	allow_adoption = true
	forget_on_destroy = false
}
`
}

func TestDeviceUpdateWaitsForNewValues(t *testing.T) {
	ctx := t.Context()
	site := "default"
	mac := "00:27:22:00:00:01"
	device := func(ledOverride string) string {
		return fmt.Sprintf(`{"_id": "000000000000000000000001", "mac": %q,
			"name": "USW-24-PoE", "model": "USL24P", "type": "usw", "adopted": true,
			"state": 1, "version": "7.5.15.17146", "led_override": %q,
			"stp_priority": "4096", "config_network": {"type": "dhcp"},
			"port_overrides": [{"port_idx": 1, "name": "Uplink", "op_mode": "switch"}]}`,
			mac, ledOverride)
	}
	// After a PUT, stat/device shows the old led_override for two more reads,
	// while the device is connected.
	var stored, shown string
	var staleReads int
	var sent []map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPut:
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Error(err)
			}
			sent = append(sent, body)
			stored, _ = body["led_override"].(string)
			staleReads = 2
			fmt.Fprint(w, `{"meta": {"rc": "ok"}, "data": []}`)
		case r.Method == http.MethodGet && path.Base(r.URL.Path) == mac:
			if staleReads > 0 {
				staleReads--
			} else {
				shown = stored
			}
			fmt.Fprintf(w, `{"meta": {"rc": "ok"}, "data": [%s]}`, device(shown))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()
	apiClient, err := unifi.New(ctx, &unifi.Config{BaseURL: server.URL, APIKey: "test"})
	if err != nil {
		t.Fatal(err)
	}
	r := &deviceResource{client: &Client{ApiClient: apiClient, Site: site}}

	// Each model is the state that a read of the device gives, with at most one attribute changed.
	tests := []struct {
		name            string
		change          func(model *deviceResourceModel)
		wantSent        int
		wantLedOverride string
	}{
		{
			name: "led_override changed",
			change: func(model *deviceResourceModel) {
				model.LedOverride = types.StringValue("on")
			},
			wantSent:        1,
			wantLedOverride: "on",
		},
		{
			name:            "nothing changed",
			change:          func(model *deviceResourceModel) {},
			wantLedOverride: "default",
		},
		{
			// Create sends no config_network when the configuration leaves it out.
			name: "config_network null",
			change: func(model *deviceResourceModel) {
				model.ConfigNetwork = types.ObjectNull(configNetworkAttrTypes())
			},
			wantLedOverride: "default",
		},
	}
	for _, test := range tests {
		stored, shown, staleReads, sent = "default", "default", 0, nil
		var before unifi.Device
		if err := json.Unmarshal([]byte(device("default")), &before); err != nil {
			t.Fatal(err)
		}
		var model deviceResourceModel
		var diags diag.Diagnostics
		r.setResourceData(ctx, &diags, &before, &model, site)
		test.change(&model)

		diags.Append(r.updateDevice(ctx, &model)...)
		if diags.HasError() {
			t.Errorf("%s: %v", test.name, diags)
		}
		if len(sent) != test.wantSent {
			t.Errorf("%s: sent %d PUTs, want %d", test.name, len(sent), test.wantSent)
		}
		if got := model.LedOverride.ValueString(); got != test.wantLedOverride {
			t.Errorf("%s: led_override read back as %q, want %q",
				test.name, got, test.wantLedOverride)
		}
	}
}
