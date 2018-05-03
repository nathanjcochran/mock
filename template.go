package main

var defaultTmpl = `package {{.PackageName}}_test
{{ if gt (len .Imports) 0 }}
imports (
{{ .ImportsString }}
)
{{- end }}

type {{ .InterfaceName }}Mock struct {
	{{- range $method := .Methods }}
	{{ $method.Name }}Stub func({{ $method.ParamsString }}) {{ $method.ResultsString }}
	{{ $method.Name }}Called int
	{{- end }}
}

var _ {{ .InterfaceName }} = &{{ .InterfaceName }}Mock{}
`
