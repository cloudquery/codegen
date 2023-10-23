package docs

import (
	"fmt"
	"strings"

	"github.com/invopop/jsonschema"
)

type reference struct {
	key         refKey
	level       int                    // header & nesting level
	definitions jsonschema.Definitions // currently used definitions
}

func (r *reference) schema() (*jsonschema.Schema, error) {
	key := r.key.unwrap()

	sc, ok := r.definitions[key]
	if !ok {
		return nil, fmt.Errorf("missing definition for key %q, possibly incomplete schema", key)
	}

	return sc, nil
}

func (r *reference) header() string {
	return strings.Repeat("#", min(r.level, 6)) + ` <a name="` + r.key.anchor() + `"></a>` + r.key.name()
}

func (r *reference) newReferences(sc *jsonschema.Schema, keys []refKey) []reference {
	if len(keys) == 0 {
		return nil
	}

	refs := make([]reference, len(keys))
	for i, key := range keys {
		refs[i] = r.newRef(sc, key)
	}
	return refs
}

func (r *reference) newRef(sc *jsonschema.Schema, key refKey) reference {
	var defs jsonschema.Definitions
	if len(key.id) == 0 {
		key.id = sc.ID
		defs = sc.Definitions
	}
	if len(key.id) == 0 {
		key.id = r.key.id
		defs = r.definitions
	}

	return reference{
		key:         key,
		level:       r.level + 1,
		definitions: defs,
	}
}
