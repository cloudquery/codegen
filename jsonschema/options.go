package jsonschema

import "github.com/invopop/jsonschema"

type Option func(*jsonschema.Reflector)

func WithAddGoComments(base, path string) Option {
	return func(reflector *jsonschema.Reflector) {
		err := reflector.AddGoComments(base, path)
		if err != nil {
			panic(err)
		}
	}
}
