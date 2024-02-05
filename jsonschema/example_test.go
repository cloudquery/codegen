package jsonschema_test

import (
	"fmt"
	"log"
	"path"
	"runtime"

	"github.com/cloudquery/codegen/jsonschema"
)

func Example_basicSchema() {
	type basic struct {
		A string `json:"a" jsonschema:"minLength=2,required"`
	}

	data, err := jsonschema.Generate(new(basic))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
	// Output:
	// {
	//   "$schema": "https://json-schema.org/draft/2020-12/schema",
	//   "$id": "https://github.com/cloudquery/codegen/jsonschema_test/basic",
	//   "$ref": "#/$defs/basic",
	//   "$defs": {
	//     "basic": {
	//       "properties": {
	//         "a": {
	//           "type": "string",
	//           "minLength": 2
	//         }
	//       },
	//       "additionalProperties": false,
	//       "type": "object",
	//       "required": [
	//         "a"
	//       ]
	//     }
	//   }
	// }
}

type HTMLStruct struct {
	// A description with some <html> tags
	A string `json:"a" jsonschema:"minLength=2,required"`
}

func Example_basicSchemaWithHTML() {
	data, err := jsonschema.Generate(new(HTMLStruct), jsonschema.WithAddGoComments("github.com/cloudquery/codegen/jsonschema_test", currDir()))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
	// Output:
	// {
	//   "$schema": "https://json-schema.org/draft/2020-12/schema",
	//   "$id": "https://github.com/cloudquery/codegen/jsonschema_test/html-struct",
	//   "$ref": "#/$defs/HTMLStruct",
	//   "$defs": {
	//     "HTMLStruct": {
	//       "properties": {
	//         "a": {
	//           "type": "string",
	//           "minLength": 2,
	//           "description": "A description with some <html> tags"
	//         }
	//       },
	//       "additionalProperties": false,
	//       "type": "object",
	//       "required": [
	//         "a"
	//       ]
	//     }
	//   }
	// }
}

func currDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	return path.Dir(filename)
}
