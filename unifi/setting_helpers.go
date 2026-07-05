package unifi

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/types"
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
