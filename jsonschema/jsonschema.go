package jsonschema

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"reflect"

	"github.com/invopop/jsonschema"
)

// Generate returns a formatted JSON schema for the input struct, according to the tags
// defined by https://github.com/invopop/jsonschema
func Generate(a any) ([]byte, error) {
	sc := (&jsonschema.Reflector{RequiredFromJSONSchemaTags: true}).ReflectFromType(reflect.TypeOf(a))
	return json.MarshalIndent(sc, "", "  ")
}

func GenerateIntoFile(a any, filePath string) {
	data, err := Generate(a)
	if err != nil {
		log.Fatalf("failed to generate JSON schema for %T", a)
	}

	ensureDir(filePath)
	if err = os.WriteFile(filePath, append(data, '\n'), 0o644); err != nil {
		log.Fatalf("failed to write file %s: %s", filePath, err.Error())
	}
}

func ensureDir(filePath string) {
	if err := os.MkdirAll(path.Dir(filePath), os.ModePerm); err != nil {
		log.Fatalf("failed to create dir for file %s: %v", filePath, err)
	}
}
