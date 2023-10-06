package jsonschema

import (
	"testing"

	"github.com/invopop/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestSanitize(t *testing.T) {
	sc := &jsonschema.Schema{
		Properties: jsonschema.NewProperties(),
	}

	sc.Definitions = jsonschema.Definitions{
		"key": new(jsonschema.Schema),
	}

	require.NoError(t, Sanitize(sc))
	require.Empty(t, sc.Definitions)
}
