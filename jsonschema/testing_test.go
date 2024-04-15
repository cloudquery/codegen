package jsonschema

import (
	"testing"
)

type schemaTestCase struct {
	name   string
	schema string
	cases  []TestCase
}

func TestJSONSchemaCases(t *testing.T) {
	tests := []schemaTestCase{
		{
			name:   "simple",
			schema: `{ "type": "object", "properties": { "name": { "type": "string" } } }`,
			cases: []TestCase{
				{
					Name: "valid",
					Spec: `{ "name": "test" }`,
				},
				{
					Name: "invalid",
					Spec: `{ "name": 1 }`,
					Err:  true,
				},
				{
					Name:         "invalid with error message",
					Spec:         `{ "name": 1 }`,
					ErrorMessage: "expected string, but got number",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TestJSONSchema(t, tt.schema, tt.cases)
		})
	}
}
