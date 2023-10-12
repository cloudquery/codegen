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

	if property.Default != nil {
		_, _ = fmt.Fprintf(buff, " (default: `%v`)", property.Default)
	}

	if len(property.Pattern) > 0 {
		_, _ = fmt.Fprintf(buff, " (pattern: `%s`)", property.Pattern)
	}

	if len(property.Enum) > 0 {
		buff.WriteString(" (possible values: ")
		for i, e := range property.Enum {
			if i > 0 {
				buff.WriteString(", ")
			}
			_, _ = fmt.Fprintf(buff, "`%v`", e)
		}
		buff.WriteString(")")
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
	_type, ref = propertyTypeNoSuffix(sc)
	_type = "`" + _type + "`" // backticks for type name
	if len(ref) > 0 {
		_type = `[` + _type + `](#` + ref + `)` // link
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
	clashingRef := regexp.MustCompile(`^(.+)_\d+$`)
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
