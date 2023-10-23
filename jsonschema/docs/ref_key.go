package docs

import (
	"regexp"
	"strings"

	"github.com/invopop/jsonschema"
)

type refKey struct {
	id  jsonschema.ID // $id of the schema, differs for nested schemas
	key string        // key in definitions map
}

func (r refKey) unwrap() string {
	return strings.TrimPrefix(r.key, "#/$defs/")
}

func (r refKey) name() string {
	clashingRef := regexp.MustCompile(`^(.+)[_-]\d+$`)
	key := r.unwrap()

	match := clashingRef.FindStringSubmatch(key)
	if len(match) > 1 {
		return match[1]
	}
	return key
}

func (r refKey) anchor() string {
	return strings.ReplaceAll(r.unwrap(), "_", "-")
}

func (r refKey) link() string {
	return "[`" + r.name() + "`](#" + r.anchor() + ")"
}
