package main

var tmpl = `package {{ .Package }}
import (
	"sync/atomic"
	{{- range .Imports }}
	{{ . }}
	{{- end }}
)

type {{ .Name }}Mock struct {
	{{- range .Methods }}
	{{ .Name }}Stub func({{ .Params }}) {{ .Results }}
	{{ .Name }}Called int32
	{{- end }}
}

var _ {{ .Name }} = &{{ .Name }}Mock{}

{{- range .Methods }}

func (m *{{ $.Name }}Mock) {{ .Name }}({{ .Params.NamedString }}) {{ .Results }}{
	atomic.AddInt32(&m.{{ .Name }}Called, 1) 
	{{- if gt (len .Results) 0 }}
	return m.{{ .Name }}Stub({{ .Params.ArgsString }})
	{{- else }}
	m.{{ .Name }}Stub({{ .Params.ArgsString }})
	{{- end }}
}
{{- end -}}
`
