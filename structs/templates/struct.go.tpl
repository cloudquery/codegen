// {{ .HeaderComment }}

package {{ .PackageName }}

// {{ .OutputName }} is an optimized version of {{ .SourceName }}
// where deeply nested struct fields are replaced with map[string]any to avoid
// the decode -> allocate -> re-encode cycle.
type {{ .OutputName }} struct {
{{- range .Fields }}
	{{ .Name }} {{ .Type }} `json:"{{ .JSONTag }}"`
{{- end }}
}
