package docs

import (
	"encoding/json"
	"fmt"
	"math/big"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/invopop/jsonschema"
)

func Generate(schema []byte, headerLevel int) (string, error) {
	var root jsonschema.Schema
	if err := json.Unmarshal(schema, &root); err != nil {
		return "", err
	}

	buff := new(strings.Builder)
	toc, err := generate(root.Definitions, unwrapRef(root.Ref), headerLevel, buff)
	return toc + "\n\n" + buff.String(), err
}

type reference struct {
	key   string
	level int
}

func generate(definitions jsonschema.Definitions, ref string, level int, buff *strings.Builder) (toc string, err error) {
	processed := make(map[string]struct{}, len(definitions))
	references := make([]reference, 1, len(definitions))
	references[0] = reference{key: ref, level: level + 1} // +1 as toc is on the level

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

		currSchema, ok := definitions[curr.key]
		if !ok {
			return toc, fmt.Errorf("missing definition for key %q, possibly incomplete schema", curr.key)
		}

		// we prepend references to make the docs more localized
		references = append(writeDefinition(curr, currSchema, buff), references...)
		toc += "\n" + strings.Repeat("  ", curr.level-level-1) + "* " + linkTo(curr.key)
	}
	return toc, nil
}

func writeDefinition(ref reference, sc *jsonschema.Schema, buff *strings.Builder) []reference {
	buff.WriteString(header(ref))
	buff.WriteString("\n")

	if len(sc.Title) > 0 {
		buff.WriteString("\n")
		buff.WriteString(sc.Title)
		buff.WriteString("\n")
	}

	writeDescription(sc, buff)

	if sc.Properties.Len() == 0 {
		buff.WriteString("\n")
		newRef := writeInlineDefinition(sc, slices.Contains(sc.Required, ref.key), buff)
		if len(newRef) > 0 {
			return []reference{{key: newRef, level: ref.level + 1}}
		}
		return nil
	}

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

func writeInlineDefinition(sc *jsonschema.Schema, required bool, buff *strings.Builder) (ref string) {
	return writeProperty(sc, required, buff)
}

func header(ref reference) string {
	return strings.Repeat("#", min(ref.level, 6)) + ` <a name="` + anchorValue(ref.key) + `"></a>` + trimClashingSuffix(ref.key)
}

func docProperty(key string, property *jsonschema.Schema, required bool, buff *strings.Builder) (ref string) {
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
func writeProperty(property *jsonschema.Schema, required bool, buff *strings.Builder) (ref string) {
	sc, nullable := unwrapNullable(property)
	propType, ref := propertyType(sc)
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

	return ref
}

func writeDescription(sc *jsonschema.Schema, buff *strings.Builder) {
	if len(sc.Description) == 0 {
		return
	}

	buff.WriteString("\n  ")
	buff.WriteString(strings.ReplaceAll(sc.Description, "\n", "\n  "))
	buff.WriteString("\n")
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
			_, _ = fmt.Fprintf(buff, "`%v`", e)
		}
		buff.WriteString(")")
	}

	if sc.Default != nil {
		_, _ = fmt.Fprintf(buff, " (default: `%s`)", anyValue(sc.Default))
	}
}

func anyValue(a any) string {
	switch a := a.(type) {
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

func propertyType(sc *jsonschema.Schema) (_type string, ref string) {
	_type, ref = propertyTypeNoSuffix(sc)
	_type = "`" + _type + "`" // backticks for type name
	if len(ref) > 0 {
		_type = `[` + _type + `](#` + anchorValue(ref) + `)` // link
	}
	_type = `(` + _type + `)` // wrap in brackets
	return _type, ref
}

func propertyTypeNoSuffix(sc *jsonschema.Schema) (_type string, ref string) {
	sc, _ = unwrapNullable(sc)

	if isAnything(sc) {
		return "anything", ""
	}

	if ref = unwrapRef(sc.Ref); len(ref) > 0 {
		return trimClashingSuffix(ref), ref
	}

	if _type, ref, ok := mapType(sc); ok {
		return _type, ref
	}

	if sc.Type != "array" {
		return sc.Type, ""
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

func isAnything(sc *jsonschema.Schema) bool {
	data, err := json.Marshal(sc)
	if err != nil {
		panic(err)
	}
	return string(data) == "true"
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
	clashingRef := regexp.MustCompile(`^(.+)[_-]\d+$`)
	if !clashingRef.MatchString(ref) {
		return ref
	}

	return clashingRef.FindStringSubmatch(ref)[1]
}

func linkTo(key string) string {
	return "[`" + trimClashingSuffix(key) + "`](#" + anchorValue(key) + ")"
}

func anchorValue(key string) string {
	return strings.ReplaceAll(key, "_", "-")
}
