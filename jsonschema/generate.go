package jsonschema

import (
	"encoding/json"
	"log"
	"os"
	"path"

	"github.com/invopop/jsonschema"
)

// Generate returns a formatted JSON schema for the input struct, according to the tags
// defined by https://github.com/invopop/jsonschema
func Generate(a any, options ...Option) ([]byte, error) {
	reflector := &jsonschema.Reflector{
		RequiredFromJSONSchemaTags: true,
		NullableFromType:           true,
	}
	for _, opt := range options {
		opt(reflector)
	}
	sc := reflector.Reflect(a)
	if err := Sanitize(sc); err != nil {
		return nil, err
	}

	return json.MarshalIndent(sc, "", "  ")
}

func GenerateIntoFile(a any, filePath string, options ...Option) {
	data, err := Generate(a, options...)
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
