package jsonschema_test

import (
	"fmt"
	"log"

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
	//   "$ref": "#/$defs/github.com~1cloudquery~1codegen~1jsonschema_test~1basic",
	//   "$defs": {
	//     "github.com/cloudquery/codegen/jsonschema_test/basic": {
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
