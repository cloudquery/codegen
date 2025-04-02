package docs

import (
	"encoding/json"
	"fmt"
	"math/big"
	"slices"
	"strconv"
	"strings"

	"github.com/invopop/jsonschema"
)

func generateDoc(root jsonschema.Schema, headerLevel int) (string, error) {
	buff := new(strings.Builder)
	toc, err := generate(&root, headerLevel, buff)
	return toc + "\n\n" + buff.String(), err
}

// GenerateFromSchema generates a markdown documentation from a jsonschema.Schema. During the writing process the `Comment` attribute can be overwritten
// To avoid this use the `Generate` function which will not modify the original schema
func GenerateFromSchema(schema jsonschema.Schema, headerLevel int) (string, error) {
	return generateDoc(schema, headerLevel)
}

func Generate(schema []byte, headerLevel int) (string, error) {
	var root jsonschema.Schema
	if err := json.Unmarshal(schema, &root); err != nil {
		return "", err
	}
	return generateDoc(root, headerLevel)
}

func generate(root *jsonschema.Schema, level int, buff *strings.Builder) (toc string, err error) {
	processed := make(map[refKey]struct{}, len(root.Definitions))
	references := make([]reference, 1, len(root.Definitions))
	references[0] = reference{
		key:         refKey{id: root.ID, key: root.Ref},
		level:       level + 1,
		definitions: root.Definitions,
	} // +1 as toc is on the level

	toc = strings.Repeat("#", level) + " Table of contents\n"
	var curr reference
	for len(references) > 0 {
		curr, references = references[0], references[1:]
		if _, ok := processed[curr.key]; ok {
			// skip already processed
			continue
		}
		if len(processed) > 0 {
			buff.WriteString("\n")
		}
		processed[curr.key] = struct{}{}

		currSchema, err := curr.schema()
		if err != nil {
			return toc, err
		}

		// we prepend references to make the docs more localized
		references = append(writeDefinition(curr, currSchema, buff), references...)
		toc += "\n" + strings.Repeat("  ", curr.level-level-1) + "* " + curr.key.link()
	}
	return toc, nil
}

func writeDefinition(ref reference, sc *jsonschema.Schema, buff *strings.Builder) []reference {
	buff.WriteString(ref.header())
	buff.WriteString("\n")

	if len(sc.Title) > 0 {
		buff.WriteString("\n")
		buff.WriteString(sc.Title)
		buff.WriteString("\n")
	}

	writeDescription(sc, buff)

	if sc.Properties.Len() == 0 {
		buff.WriteString("\n")
		return ref.newReferences(sc, writeInlineDefinition(sc, false, buff))
	}

	refs := make([]reference, 0, sc.Properties.Len()) // prealloc to some meaningful len
	for prop := sc.Properties.Oldest(); prop != nil; prop = prop.Next() {
		buff.WriteString("\n")
		refs = append(refs, ref.newReferences(sc, docProperty(prop.Key, prop.Value, slices.Contains(sc.Required, prop.Key), buff))...)
	}

	return refs
}

func writeInlineDefinition(sc *jsonschema.Schema, required bool, buff *strings.Builder) (refs []refKey) {
	return writeProperty(sc, required, buff)
}

func docProperty(key string, property *jsonschema.Schema, required bool, buff *strings.Builder) (refs []refKey) {
	buff.WriteString("* `" + key + "`")
	sc, _ := unwrapNullable(property)

	if len(sc.Title) > 0 {
		buff.WriteString(": ")
		buff.WriteString(sc.Title)
		buff.WriteString("\n  ")
	} else {
		// if no title is present we want the type definition inline
		buff.WriteString(" ")
	}

	return writeProperty(property, required, buff)
}

// writeProperty starts off with the type definition without any line breaks & prefixes
func writeProperty(property *jsonschema.Schema, required bool, buff *strings.Builder) (refs []refKey) {
	sc, nullable := unwrapNullable(property)
	propType, refs := propertyType(sc)
	buff.WriteString(propType)
	if nullable {
		buff.WriteString(" (nullable)")
	}

	if required {
		buff.WriteString(" (required)")
	}

	writeValueAnnotations(sc, buff)
	buff.WriteString("\n")

	writeDescription(sc, buff)

	return refs
}

func writeDescription(sc *jsonschema.Schema, buff *strings.Builder) {
	if len(sc.Description) == 0 || sc.Comments == "skip_description" {
		return
	}

	buff.WriteString("\n  ")
	buff.WriteString(strings.ReplaceAll(sc.Description, "\n", "\n  "))
	buff.WriteString("\n")

	sc.Comments = "skip_description"
}

func writeValueAnnotations(sc *jsonschema.Schema, buff *strings.Builder) {
	if sc.Type == "array" {
		// tricky, we will traverse the items first
		writeValueAnnotations(sc.Items, buff)
	}

	if len(sc.Format) > 0 {
		_, _ = fmt.Fprintf(buff, " ([format](https://json-schema.org/draft/2020-12/json-schema-validation#section-7): `%s`)", sc.Format)
	}

	if len(sc.Pattern) > 0 {
		pattern := strings.Trim(strconv.Quote(sc.Pattern), `"`)
		_, _ = fmt.Fprintf(buff, " ([pattern](https://json-schema.org/draft/2020-12/json-schema-validation#section-6.3.3): `%s`)", pattern)
	}

	if borders := valueBorders(sc); len(borders) > 0 {
		_, _ = fmt.Fprintf(buff, " (range: `%s`)", borders)
	}

	if len(sc.Enum) > 0 {
		buff.WriteString(" (possible values: ")
		for i, e := range sc.Enum {
			if i > 0 {
				buff.WriteString(", ")
			}
			_, _ = fmt.Fprintf(buff, "`%v`", anyValue(e))
		}
		buff.WriteString(")")
	}

	if sc.Default != nil {
		_, _ = fmt.Fprintf(buff, " (default: `%s`)", anyValue(sc.Default))
	}
}

func anyValue(a any) string {
	switch a := a.(type) {
	case string:
		if len(a) == 0 {
			// Markdown needs at least 1 space to represent empty string
			return " "
		}
	case float32:
		if float32(int64(a)) == a {
			return fmt.Sprintf("%d", int64(a))
		}
	case float64:
		if float64(int64(a)) == a {
			return fmt.Sprintf("%d", int64(a))
		}
	case []any:
		elems := make([]string, len(a))
		for i, v := range a {
			elems[i] = anyValue(v)
		}
		return fmt.Sprintf("%+v", elems)
	}
	return fmt.Sprintf("%v", a)
}

func valueBorders(sc *jsonschema.Schema) string {
	minimum, _ := new(big.Rat).SetString(string(sc.Minimum))
	exclMinimum, _ := new(big.Rat).SetString(string(sc.ExclusiveMinimum))
	var l string
	switch {
	case minimum == nil:
		if exclMinimum != nil {
			l = "(" + exclMinimum.RatString()
		}
	case exclMinimum == nil:
		if minimum != nil {
			l = "[" + minimum.RatString()
		}
	default:
		if minimum.Cmp(exclMinimum) <= 0 {
			l = "(" + exclMinimum.RatString()
		} else {
			l = "[" + minimum.RatString()
		}
	}

	maximum, _ := new(big.Rat).SetString(string(sc.Maximum))
	exclMaximum, _ := new(big.Rat).SetString(string(sc.ExclusiveMaximum))
	var r string
	switch {
	case maximum == nil:
		if exclMaximum != nil {
			r = exclMaximum.RatString() + ")"
		}
	case exclMaximum == nil:
		if maximum != nil {
			r = maximum.RatString() + "]"
		}
	default:
		if maximum.Cmp(exclMaximum) <= 0 {
			r = exclMaximum.RatString() + ")"
		} else {
			r = maximum.RatString() + "]"
		}
	}

	switch {
	case len(l) == 0:
		if len(r) > 0 {
			return "(-∞," + r
		}
	case len(r) == 0:
		if len(l) > 0 {
			return l + ",+∞)"
		}
	default:
		return l + "," + r
	}

	return ""
}

func unwrapNullable(sc *jsonschema.Schema) (*jsonschema.Schema, bool) {
	if len(sc.OneOf) == 2 && sc.OneOf[1].Type == "null" {
		return sc.OneOf[0], true
	}
	return sc, false
}

func propertyType(sc *jsonschema.Schema) (_type string, refs []refKey) {
	types := propertyTypeNoSuffix(sc)

	if len(types) == 1 {
		t := types[0]
		return "(" + t.printable() + ")", t.refs()
	}

	parts := make([]string, len(types)) // >1 part ~ oneOf/anyOf
	for i, t := range types {
		parts[i] = t.printable()
		refs = append(refs, t.refs()...)
	}

	return "(" + strings.Join(parts[:len(parts)-1], ", ") + " or " + parts[len(parts)-1] + ")", refs
}

func propertyTypeNoSuffix(sc *jsonschema.Schema) []typeReference {
	sc, _ = unwrapNullable(sc)

	if isAnything(sc) {
		return []typeReference{{name: "anything"}}
	}

	if len(sc.Ref) > 0 {
		ref := refKey{key: sc.Ref}
		return []typeReference{{name: ref.name(), ref: &ref}}
	}

	if _types, ok := mapType(sc); ok {
		return _types
	}

	if _types, ok := arrayType(sc); ok {
		return _types
	}

	if _types, ok := anyOfType(sc); ok {
		return _types
	}

	if _types, ok := oneOfType(sc); ok {
		return _types
	}

	// default case
	return []typeReference{{name: sc.Type}}
}

func isAnything(sc *jsonschema.Schema) bool {
	data, err := json.Marshal(sc)
	if err != nil {
		panic(err)
	}
	return string(data) == "true"
}

func isNothing(sc *jsonschema.Schema) bool {
	data, err := json.Marshal(sc)
	if err != nil {
		panic(err)
	}
	return string(data) == "false"
}

func mapType(sc *jsonschema.Schema) (refs []typeReference, ok bool) {
	if sc.Type != "object" || sc.AdditionalProperties == nil || isNothing(sc.AdditionalProperties) {
		return nil, false
	}
	pfx := `map[string]`
	refs = propertyTypeNoSuffix(sc.AdditionalProperties)
	for i := range refs {
		refs[i].name = pfx + refs[i].name
	}
	return refs, true
}

func arrayType(sc *jsonschema.Schema) (refs []typeReference, ok bool) {
	if sc.Type != "array" {
		return nil, false
	}
	item, nullable := unwrapNullable(sc.Items)
	pfx := "[]"
	if nullable {
		pfx += "*"
	}
	refs = propertyTypeNoSuffix(item)
	for i := range refs {
		refs[i].name = pfx + refs[i].name
	}
	return refs, true
}

func oneOfType(sc *jsonschema.Schema) (refs []typeReference, ok bool) {
	if len(sc.OneOf) == 0 {
		return nil, false
	}

	return ofTypes(sc.OneOf), true
}

func anyOfType(sc *jsonschema.Schema) (refs []typeReference, ok bool) {
	if len(sc.AnyOf) == 0 {
		return nil, false
	}

	return ofTypes(sc.AnyOf), true
}

func ofTypes(types []*jsonschema.Schema) []typeReference {
	refs := make([]typeReference, 0, len(types))
	for _, t := range types {
		refs = append(refs, propertyTypeNoSuffix(t)...)
	}
	return refs
}
