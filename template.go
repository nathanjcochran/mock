package main

var tmpl = `package {{ .Package }}
import (
	"sync"
	{{- range .Imports }}
	{{ . }}
	{{- end }}
)

{{ range $method := .Methods -}}
type {{ $.Name }}{{ .Name }}Args struct {
	{{ .Params.NamedFields }}
}
type {{ $.Name }}{{ .Name }}Results struct {
	{{ .Results.NamedFields }}
}
{{ end }}

type {{ .Name }}Mock struct {
	T *testing.T
	{{- range .Methods }}
	{{ .Name }}Expects []{{ $.Name }}{{ .Name }}Args
	{{ .Name }}Returns []{{ $.Name }}{{ .Name}}Results
	{{ .Name }}Calls []{{ $.Name }}{{ .Name }}Args
	{{ .Name }}Lock sync.Mutex
	{{- end }}
}

var _ {{ .Name }} = &{{ .Name }}Mock{}

{{- range $method := .Methods }}

func (m *{{ $.Name }}Mock) {{ .Name }}({{ .Params.NamedString }}) {{ .Results }}{
	m.{{ .Name }}Lock.Lock()
	defer m.{{ .Name }}Lock.Unlock()

	m.{{ .Name }}Calls = append(m.{{ .Name }}Calls, {{ $.Name }}{{ .Name }}Args{ {{ .Params.ArgsString }} })

	if m.T != nil && m.{{ .Name }}Expects != nil {
		if len(m.{{ .Name }}Calls) > len(m.{{ .Name }}Expects) {
			m.T.Errorf("Unexpected call to {{ $.Name }}.{{ .Name }}. Expected calls: %d. Actual: %d",
				len(m.{{ .Name }}Expects), len(m.{{ .Name }}Calls))
		} else {
			expected := m.{{ .Name }}Expects[len(m.{{ .Name }}Calls) - 1]
			{{- range $index, $element := .Params }}
			if !reflect.DeepEqual(expected.{{ .FieldName $index }}, {{ .Named $index }}){
				m.T.Errorf("{{ $.Name }}.{{ $method.Name }} expected {{ .Named $index }}: %v. Actual: %v", expected.{{ .FieldName $index}}, {{ .Named $index }})
			}
			{{- end }}
		}
	}

	{{- if gt (len .Results) 0 }}
	results := m.{{ .Name }}Returns[len(m.{{ .Name }}Calls) - 1]
	return {{ .Results.ReturnList "results" }}
	{{- end }}
}
{{- end -}}
`
