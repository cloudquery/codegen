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
	buff.WriteString(strings.Repeat("#", headerLevel+1))
	buff.WriteString(" ")

	ref := unwrapRef(root.Ref)
	buff.WriteString(ref)
	buff.WriteString("\n")

	sc := root.Definitions[ref]
	for prop := sc.Properties.Oldest(); prop != nil; prop = prop.Next() {
		buff.WriteString("\n")
		buff.WriteString(propertyDoc(prop.Key, prop.Value, slices.Contains(sc.Required, prop.Key)))
	}
	return buff.String(), nil
}

func propertyDoc(key string, property *jsonschema.Schema, required bool) string {
	doc := "* `" + key + "` "

	sc, nullable := unwrapNullable(property)
	doc += propertyType(sc)
	if nullable {
		doc += " (nullable)"
	}

	if required {
		doc += " (required)"
	}

	if property.Default != nil {
		doc += fmt.Sprintf(" (default=`%v`)", property.Default)
	}

	return doc
}

func unwrapNullable(sc *jsonschema.Schema) (*jsonschema.Schema, bool) {
	if len(sc.OneOf) == 2 && sc.OneOf[1].Type == "null" {
		return sc.OneOf[0], true
	}
	return sc, false
}

func propertyType(sc *jsonschema.Schema) string {
	if ref := unwrapRef(sc.Ref); len(ref) > 0 {
		return "([`" + trimClashingSuffix(ref) + "`](#" + ref + "))"
	}

	if sc.Type != "array" {
		return "(`" + sc.Type + "`)"
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
	return pfx + propertyTypeNoSuffix(item)
}

func propertyTypeNoSuffix(sc *jsonschema.Schema) string {
	if ref := unwrapRef(sc.Ref); len(ref) > 0 {
		return trimClashingSuffix(ref) + "`](#" + ref + ")"
	}

	if sc.Type != "array" {
		return sc.Type + "`)"
	}

	// arrays are a bit tricky
	item, nullable := unwrapNullable(sc.Items)
	pfx := "[]"
	if nullable {
		pfx += "*"
	}
	return pfx + propertyTypeNoSuffix(item)
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
