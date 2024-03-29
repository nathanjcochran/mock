package main

var tmpl = `package {{ .Package }}
import (
	"sync/atomic"
	{{- range .Imports }}
	{{ . }}
	{{- end }}
)

// {{ .Name }}Mock is a mock implementation of the {{ .Name }}
// interface.
type {{ .Name }}Mock struct {
	T *testing.T
	{{- range .Methods }}
	{{ .Name }}Stub func({{ .Params }}) {{ .Results }}
	{{ .Name }}Called int32
	{{- end }}
}

var _ {{ .Name }} = &{{ .Name }}Mock{}

{{- range .Methods }}

// {{ .Name}} is a stub for the {{ $.Name }}.{{ .Name }}
// method that records the number of times it has been called.
func (m *{{ $.Name }}Mock) {{ .Name }}({{ .Params.NamedString }}) {{ .Results }}{
	atomic.AddInt32(&m.{{ .Name }}Called, 1) 
	if m.{{ .Name }}Stub == nil {
		if m.T != nil {
			m.T.Error("{{ .Name }}Stub is nil")
		}
		panic("{{ .Name }} unimplemented")
	}
	{{- if gt (len .Results) 0 }}
	return m.{{ .Name }}Stub({{ .Params.ArgsString }})
	{{- else }}
	m.{{ .Name }}Stub({{ .Params.ArgsString }})
	{{- end }}
}
{{- end -}}
`
