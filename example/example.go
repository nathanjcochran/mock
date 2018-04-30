package example

import (
	. "go/ast"
	"html/template"
	cheese "text/template"
)

//go:generate mocker MyInterface
type MyInterface interface {
	MyMethod1(param1 string, param2 int) (cheese.Template, error)
	MyMethod2(param1 template.Template) struct{ MyField int }
	MyMethod3(File) struct{ cheese.Template }
}
