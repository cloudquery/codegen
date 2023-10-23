package docs

type typeReference struct {
	name string
	ref  *refKey
}

func (t typeReference) printable() string {
	name := backtick(t.name) // backticks for type name
	if t.ref == nil {
		return name
	}
	return `[` + name + `](#` + t.ref.anchor() + `)` // link
}

func (t typeReference) refs() []refKey {
	if t.ref == nil {
		return nil
	}
	return []refKey{*t.ref}
}

func backtick(s string) string {
	return "`" + s + "`"
}
