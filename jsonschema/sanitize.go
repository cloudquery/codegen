package jsonschema

import (
	"net/url"
	"strings"

	"github.com/invopop/jsonschema"
	"golang.org/x/exp/maps"
)

func Sanitize(sc *jsonschema.Schema) {
	refs := collectRefs(sc)

	for key := range sc.Definitions {
		if _, ok := refs[key]; !ok {
			delete(sc.Definitions, key)
		}
	}

	for p := sc.Properties.Oldest(); p != nil; p = p.Next() {
		Sanitize(p.Value)
	}
}

func collectRefs(sc *jsonschema.Schema) map[string]bool {
	refs := make(map[string]bool)
	if len(sc.Ref) > 0 {
		refs[unescapeRef(sc.Ref)] = true
	}

	for p := sc.Properties.Oldest(); p != nil; p = p.Next() {
		maps.Copy(refs, collectRefs(p.Value))
	}

	return refs
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
