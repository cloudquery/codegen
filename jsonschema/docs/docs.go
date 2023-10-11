package docs

import (
	"encoding/json"
	"strings"

	"github.com/invopop/jsonschema"
)

func Generate(schema []byte, headerLevel int) (string, error) {
	var sc jsonschema.Schema
	if err := json.Unmarshal(schema, &sc); err != nil {
		return "", err
	}

	buff := new(strings.Builder)
	buff.WriteString(strings.Repeat("#", headerLevel+1))
	buff.WriteString(" ")

	ref := strings.TrimPrefix(sc.Ref, "#/$defs/")
	buff.WriteString(ref)

	return buff.String(), nil
}
