package main

var defaultTmpl = `package {{ .PackageName }}
{{ if gt (len .Imports) 0 }}
import (
{{ .ImportsString }}
)
{{- end }}

type {{ .InterfaceName }}Mock struct {
	{{- range .Methods }}
	{{ .Name }}Stub func({{ .ParamsString }}) {{ .ResultsString }}
	{{ .Name }}Called int
	{{- end }}
}

var _ {{ .InterfaceName }} = &{{ .InterfaceName }}Mock{}

{{- range .Methods }}

func (m *{{ $.InterfaceName }}Mock) {{ .Name }}({{ .NamedParamsString }}) {{ .ResultsString }}{
	m.{{ .Name }}Called ++
	{{- if gt (len .Results) 0 }}
	return m.{{ .Name }}Stub({{ .ParamNamesString }})
	{{- else }}
	m.{{ .Name }}Stub({{ .ParamNamesString }})
	{{- end }}
}
{{- end -}}
`
