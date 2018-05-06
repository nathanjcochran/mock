package main

var defaultTmpl = `package {{ .Package }}
{{ if gt (len .Imports) 0 }}
import (
	{{- range .Imports }}
	{{ . }}
	{{- end }}
)
{{- end }}

type {{ .Name }}Mock struct {
	{{- range .Methods }}
	{{ .Name }}Stub func({{ .Params }}) {{ .Results }}
	{{ .Name }}Called int
	{{- end }}
}

var _ {{ .Name }} = &{{ .Name }}Mock{}

{{- range .Methods }}

func (m *{{ $.Name }}Mock) {{ .Name }}({{ .Params.NamedString }}) {{ .Results }}{
	m.{{ .Name }}Called ++
	{{- if gt (len .Results) 0 }}
	return m.{{ .Name }}Stub({{ .Params.ArgsString }})
	{{- else }}
	m.{{ .Name }}Stub({{ .Params.ArgsString }})
	{{- end }}
}
{{- end -}}
`
