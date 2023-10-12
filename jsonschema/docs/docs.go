package docs

import (
	"encoding/json"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/invopop/jsonschema"
)

func Generate(schema []byte, headerLevel int) (string, error) {
	var root jsonschema.Schema
	if err := json.Unmarshal(schema, &root); err != nil {
		return "", err
	}

	buff := new(strings.Builder)
	err := generate(root.Definitions, []reference{{key: unwrapRef(root.Ref), level: headerLevel}}, buff)
	return buff.String(), err
}

type reference struct {
	key   string
	level int
}

func generate(definitions jsonschema.Definitions, references []reference, buff *strings.Builder) error {
	processed := make(map[string]struct{}, len(definitions))
	var curr reference
	for len(references) > 0 {
		curr, references = references[0], references[1:]
		if _, ok := processed[curr.key]; ok {
			// skip already processed
			continue
		}
		if len(processed) > 0 {
			buff.WriteString("\n\n")
		}
		processed[curr.key] = struct{}{}

		currSchema, ok := definitions[curr.key]
		if !ok {
			return fmt.Errorf("missing definition for key %q, possibly incomplete schema", curr.key)
		}

		// we prepend references to make the docs more localized
		references = append(writeDefinition(curr, currSchema, buff), references...)
	}
	return nil
}

func writeDefinition(ref reference, sc *jsonschema.Schema, buff *strings.Builder) []reference {
	buff.WriteString(strings.Repeat("#", ref.level))
	buff.WriteString(` <a name="` + ref.key + `"></a>`) // add anchor
	buff.WriteString(unwrapRef(ref.key))

	refs := make([]reference, 0, sc.Properties.Len()) // prealloc to some meaningful len
	for prop := sc.Properties.Oldest(); prop != nil; prop = prop.Next() {
		buff.WriteString("\n")
		newRef := docProperty(prop.Key, prop.Value, slices.Contains(sc.Required, prop.Key), buff)
		if len(newRef) > 0 {
			refs = append(refs, reference{key: newRef, level: ref.level + 1})
		}
	}

	return refs
}

func docProperty(key string, property *jsonschema.Schema, required bool, buff *strings.Builder) (ref string) {
	buff.WriteString("* `" + key + "` ")

	sc, nullable := unwrapNullable(property)
	propType, ref := propertyType(sc)
	buff.WriteString(propType)
	if nullable {
		buff.WriteString(" (nullable)")
	}

	if required {
		buff.WriteString(" (required)")
	}

	if property.Default != nil {
		buff.WriteString(fmt.Sprintf(" (default=`%v`)", property.Default))
	}

	return ref
}

func unwrapNullable(sc *jsonschema.Schema) (*jsonschema.Schema, bool) {
	if len(sc.OneOf) == 2 && sc.OneOf[1].Type == "null" {
		return sc.OneOf[0], true
	}
	return sc, false
}

func propertyType(sc *jsonschema.Schema) (_type string, ref string) {
	if ref = unwrapRef(sc.Ref); len(ref) > 0 {
		_type, ref = propertyTypeNoSuffix(sc)
		return "([`" + _type + `)`, ref
	}

	if sc.Type != "array" {
		_type, ref = propertyTypeNoSuffix(sc)
		return "(`" + _type + `)`, ref
	}

	// arrays are a bit tricky
	item, nullable := unwrapNullable(sc.Items)
	pfx := "(`[]"
	if len(item.Ref) > 0 {
		pfx = "([`[]"
	}
	if nullable {
		pfx += "*"
	}
	_type, ref = propertyTypeNoSuffix(item)
	return pfx + _type + `)`, ref
}

func propertyTypeNoSuffix(sc *jsonschema.Schema) (_type string, ref string) {
	sc, _ = unwrapNullable(sc)

	if ref = unwrapRef(sc.Ref); len(ref) > 0 {
		return trimClashingSuffix(ref) + "`](#" + ref + ")", ref
	}

	if _type, ref, ok := mapType(sc); ok {
		return _type, ref
	}

	if sc.Type != "array" {
		return sc.Type + "`", ""
	}

	// arrays are a bit tricky
	item, nullable := unwrapNullable(sc.Items)
	pfx := "[]"
	if nullable {
		pfx += "*"
	}
	_type, ref = propertyTypeNoSuffix(item)
	return pfx + _type, ref
}

func mapType(sc *jsonschema.Schema) (_type string, ref string, ok bool) {
	if sc.Type != "object" || sc.AdditionalProperties == nil {
		return "", "", false
	}
	pfx := `map[string]`
	_type, ref = propertyTypeNoSuffix(sc.AdditionalProperties)
	return pfx + _type, ref, true
}

func unwrapRef(ref string) string {
	return strings.TrimPrefix(ref, "#/$defs/")
}

func trimClashingSuffix(ref string) string {
	clashingRef := regexp.MustCompile(`^(.+)_\d+$`)
	if !clashingRef.MatchString(ref) {
		return ref
	}

	return clashingRef.FindStringSubmatch(ref)[1]
}
