package jsonschema

import (
	"encoding/json"
	"reflect"

	"github.com/invopop/jsonschema"
)

func Generate(a any) ([]byte, error) {
	sc := (&jsonschema.Reflector{RequiredFromJSONSchemaTags: true}).ReflectFromType(reflect.TypeOf(a))
	return json.MarshalIndent(sc, "", "  ")
}
