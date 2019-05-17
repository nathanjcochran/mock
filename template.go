package main

var tmpl = `package {{ .Package }}
import (
	"sync"
	"testing"
	"github.com/google/go-cmp/cmp"
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

	{{- if gt (len .Results) 0 }}
	results := m.{{ .Name }}Returns[len(m.{{ .Name }}Calls) - 1]
	return {{ .Results.ReturnList "results" }}
	{{- end }}
}
{{ end -}}

func (m *{{ $.Name }}Mock) Assert(t *testing.T) {
	{{- range $method := .Methods }}
	if m.{{ .Name}}Expects != nil {
		if diff := cmp.Diff(m.{{ .Name }}Expects, m.{{ .Name }}Calls); diff != "" {
			t.Errorf("{{ $.Name }}.{{ .Name }} did not receive expected arguments:\n%s", diff)
		}
	}
	{{- end }}
}
`
