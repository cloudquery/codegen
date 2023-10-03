// Code generated by codegen; DO NOT EDIT.
package {{.PackageName}}

import (
    {{- range .Imports }}
    "{{ . }}"
    {{- end }}
)

{{ range .Clients }}
//go:generate mockgen -package=mocks -destination=../mocks/{{$.FileName}}.go -source={{$.FileName}}.go {{.ClientName}}
type {{.ClientName}} interface {
    {{- range $sig := .Signatures }}
    {{ $sig }}
    {{- end }}
}
{{ end }}