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
			buff.WriteString("\n\n")
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
	buff.WriteString("* `" + key + "` ")
	return writeProperty(property, required, buff)
}

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

	return ref
}

func writeValueAnnotations(sc *jsonschema.Schema, buff *strings.Builder) {
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
		_, _ = fmt.Fprintf(buff, " (default: `%v`)", sc.Default)
	}
}

func valueBorders(sc *jsonschema.Schema) string {
	var l string
	switch {
	case len(sc.Minimum) == 0:
		if len(sc.ExclusiveMinimum) > 0 {
			l = "(" + string(sc.ExclusiveMinimum)
		}
	case len(sc.ExclusiveMinimum) == 0:
		if len(sc.Minimum) > 0 {
			l = "[" + string(sc.Minimum)
		}
	default:
		lVal, _, err := big.ParseFloat(string(sc.Minimum), 10, 1000, big.ToZero)
		if err != nil {
			panic(fmt.Sprintf("minimum=%q is not a correct number: %s", sc.Minimum, err.Error()))
		}
		lValExcl, _, err := big.ParseFloat(string(sc.Minimum), 10, 1000, big.ToZero)
		if err != nil {
			panic(fmt.Sprintf("eclusiveMinimum=%q is not a correct number: %s", sc.Minimum, err.Error()))
		}
		if lVal.Cmp(lValExcl) <= 0 {
			l = "(" + string(sc.ExclusiveMinimum)
		} else {
			l = "[" + string(sc.Minimum)
		}
	}

	var r string
	switch {
	case len(sc.Maximum) == 0:
		if len(sc.ExclusiveMaximum) > 0 {
			r = string(sc.ExclusiveMaximum) + ")"
		}
	case len(sc.ExclusiveMaximum) == 0:
		if len(sc.Maximum) > 0 {
			r = string(sc.Maximum) + "]"
		}
	default:
		rVal, _, err := big.ParseFloat(string(sc.Maximum), 10, 1000, big.ToZero)
		if err != nil {
			panic(fmt.Sprintf("maximum=%q is not a correct number: %s", sc.Minimum, err.Error()))
		}
		rValExcl, _, err := big.ParseFloat(string(sc.ExclusiveMaximum), 10, 1000, big.ToZero)
		if err != nil {
			panic(fmt.Sprintf("eclusiveMaximum=%q is not a correct number: %s", sc.Minimum, err.Error()))
		}
		if rVal.Cmp(rValExcl) >= 0 {
			r = string(sc.ExclusiveMaximum) + ")"
		} else {
			r = string(sc.Maximum) + "]"
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
