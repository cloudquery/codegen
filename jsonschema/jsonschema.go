package jsonschema

import (
	"encoding/json"
	"reflect"

	"github.com/invopop/jsonschema"
)

// Generate returns a formatted JSON schema for the input struct, according to the tags
// defined by https://github.com/invopop/jsonschema
func Generate(a any) ([]byte, error) {
	sc := (&jsonschema.Reflector{RequiredFromJSONSchemaTags: true}).ReflectFromType(reflect.TypeOf(a))
	return json.MarshalIndent(sc, "", "  ")
}
