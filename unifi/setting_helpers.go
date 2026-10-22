package unifi

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/ubiquiti-community/go-unifi/unifi/settings"
)

// jsonStringFrom encodes a non-primitive SDK setting field (slice of structs,
// nested struct pointer, etc.) as a JSON string types.String, so it round-trips
// through Terraform state without lossy schema flattening.
func jsonStringFrom(v any) types.String {
	if v == nil {
		return types.StringNull()
	}
	// Detect zero-length slices and nil pointers via reflection-free shortcut:
	// json.Marshal returns "null" for nil interface and "[]"/"{}"/"null" for
	// empty containers — treat all of those as null state to avoid noisy diffs.
	b, err := json.Marshal(v)
	if err != nil || len(b) == 0 {
		return types.StringNull()
	}
	s := string(b)
	if s == "null" || s == "[]" || s == "{}" {
		return types.StringNull()
	}
	return types.StringValue(s)
}

// updateChangedFields sends only the fields that differ between current and setting.
// The controller keeps every field the request leaves out.
func (c *Client) updateChangedFields(
	ctx context.Context,
	site string,
	current settings.Setting,
	setting settings.Setting,
) error {
	changed, err := changedFields(current, setting)
	if err != nil {
		return err
	}
	if len(changed) == 0 {
		return nil
	}
	key, err := settings.GetSettingKey(setting)
	if err != nil {
		return err
	}
	return c.UpdateSetting(ctx, site, &settings.RawSetting{
		BaseSetting: settings.BaseSetting{Key: key},
		Data:        changed,
	})
}

// changedFields returns the JSON fields of setting whose value differs from current.
func changedFields(current, setting settings.Setting) (map[string]any, error) {
	currentFields, err := jsonFields(current)
	if err != nil {
		return nil, err
	}
	settingFields, err := jsonFields(setting)
	if err != nil {
		return nil, err
	}
	changed := map[string]any{}
	for name, value := range settingFields {
		if !reflect.DeepEqual(currentFields[name], value) {
			changed[name] = value
		}
	}
	return changed, nil
}

func jsonFields(setting settings.Setting) (map[string]any, error) {
	encoded, err := json.Marshal(setting)
	if err != nil {
		return nil, err
	}
	var fields map[string]any
	err = json.Unmarshal(encoded, &fields)
	return fields, err
}
