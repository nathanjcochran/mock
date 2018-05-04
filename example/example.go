package example

import (
	"html/template"
	. "os"
	renamed "text/template"
)

//go:generate mocker -o mock.go MyInterface
type MyInterface interface {
	MyMethod1(name string, _ int) (err1, err2 error)
	MyMethod2(templ template.Template) (out struct{ OutTemplate renamed.Template })
	MyMethod3(File) struct{ MyInterface }
}
