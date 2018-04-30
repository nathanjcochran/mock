package example

import (
	"text/template"
)

//go:generate mocker MyInterface
type MyInterface interface {
	MyMethod1(param1 string, param2 int) (template.Template, error)
	MyMethod2() error
}
