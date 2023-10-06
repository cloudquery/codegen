package jsonschema

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/invopop/jsonschema"
)

func Sanitize(sc *jsonschema.Schema) error {
	data, err := json.MarshalIndent(sc, "", "  ")
	if err != nil {
		return err
	}

	refs := make(map[string]bool)
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		const refPfx = `"$ref": `
		if strings.HasPrefix(line, refPfx) {
			refs[unescapeRef(strings.Trim(strings.TrimSuffix(strings.TrimPrefix(line, refPfx), ","), `"`))] = true
		}
	}

	for key := range sc.Definitions {
		if _, ok := refs[key]; !ok {
			delete(sc.Definitions, key)
		}
	}

	for p := sc.Properties.Oldest(); p != nil; p = p.Next() {
		if err := Sanitize(p.Value); err != nil {
			return err
		}
	}
	for _, def := range sc.Definitions {
		if err := Sanitize(def); err != nil {
			return err
		}
	}

	return nil
}
func unescapeRef(ref string) string {
	ref = strings.TrimPrefix(ref, "#/$defs/")

	var err error
	ref, err = url.PathUnescape(ref)
	if err != nil {
		panic(err)
	}

	ref = strings.ReplaceAll(ref, "~1", "/")
	ref = strings.ReplaceAll(ref, "~0", "~")
	return ref
}
