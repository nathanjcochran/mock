package main

var tmpl = `package {{ .Package }}
import (
	"sync"
	{{- range .Imports }}
	{{ . }}
	{{- end }}
)

type {{ .Name }}Mock struct {
	T *testing.T
	{{- range .Methods }}
	{{ .Name }}Func struct{
		Expect []struct{
			{{ .Params.NamedFields }}
		}
		Return []struct{
		    {{ .Results.NamedFields }}
		}
		Calls []struct{
			{{ .Params.NamedFields }}
		}
		lock sync.Mutex
	}
	{{- end }}
}

var _ {{ .Name }} = &{{ .Name }}Mock{}

{{- range $method := .Methods }}

func (m *{{ $.Name }}Mock) {{ .Name }}({{ .Params.NamedString }}) {{ .Results }}{
	m.{{ .Name }}Func.lock.Lock()
	defer m.{{ .Name }}Func.lock.Unlock()

	m.{{ .Name }}Func.Calls = append(m.{{ .Name }}Func.Calls, struct{
		{{ .Params.NamedFields }}
	}{ {{ .Params.ArgsString }} })

	if m.T != nil && m.{{ .Name }}Func.Expect != nil {
		if len(m.{{ .Name }}Func.Calls) > len(m.{{ .Name }}Func.Expect) {
			m.T.Errorf("Unexpected call to {{ $.Name }}.{{ .Name }}. Expected calls: %d. Actual: %d",
				len(m.{{ .Name }}Func.Expect), len(m.{{ .Name }}Func.Calls))
		} else {
			expected := m.{{ .Name }}Func.Expect[len(m.{{ .Name }}Func.Calls) - 1]
			{{ range $index, $element := .Params }}
			if !reflect.DeepEqual(expected.{{ .FieldName $index }}, {{ .Named $index }}){
				m.T.Errorf("{{ $method.Name }} expected {{ .Named $index }}: %v. Actual: %v", expected.{{ .FieldName $index}}, {{ .Named $index }})
			}
			{{ end }}
		}
	}

	{{- if gt (len .Results) 0 }}
	results := m.{{ .Name }}Func.Return[len(m.{{ .Name }}Func.Calls) - 1]
	return {{ .Results.ReturnList "results" }}
	{{- end }}
}
{{- end -}}
`
