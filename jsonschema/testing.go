package jsonschema

import (
	"encoding/json"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/stretchr/testify/require"
)

type TestCase struct {
	Name string
	Spec string
	Err  bool
}

func TestJSONSchema(t *testing.T, schema string, cases []TestCase) {
	t.Helper()

	validator, err := plugin.JSONSchemaValidator(schema)
	require.NoError(t, err)

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			var v any
			require.NoErrorf(t, json.Unmarshal([]byte(tc.Spec), &v), "failed input:\n%s\n", tc.Spec)
			err := validator.Validate(v)
			if tc.Err {
				require.Errorf(t, err, "failed input:\n%s\n", tc.Spec)
			} else {
				require.NoErrorf(t, err, "failed input:\n%s\n", tc.Spec)
			}
		})
	}
}

func WithRemovedKeys(t *testing.T, val any, keys ...string) string {
	data, err := json.Marshal(val)
	require.NoError(t, err)

	var m any
	require.NoError(t, json.Unmarshal(data, &m))

	switch m := m.(type) {
	case map[string]any:
		for _, k := range keys {
			delete(m, k)
		}
	default:
		t.Fatalf("failed to remove JSON keys from value of type %T", m)
	}

	data, err = json.MarshalIndent(m, "", "  ")
	require.NoError(t, err)
	return string(data)
}
