package example

import (
	"html/template"
	. "os"
	renamed "text/template"
)

type MyInterfaceMock struct {
	MyMethod1Stub   func(name string, _ int) (err1 error, err2 error)
	MyMethod1Called int
	MyMethod2Stub   func(templ template.Template) (out struct{ OutTemplate renamed.Template })
	MyMethod2Called int
	MyMethod3Stub   func(File) struct{ MyInterface }
	MyMethod3Called int
}

var _ MyInterface = &MyInterfaceMock{}

func (m *MyInterfaceMock) MyMethod1(name string, param2 int) (err1 error, err2 error) {
	m.MyMethod1Called++
	return m.MyMethod1Stub(name, param2)
}

func (m *MyInterfaceMock) MyMethod2(templ template.Template) (out struct{ OutTemplate renamed.Template }) {
	m.MyMethod2Called++
	return m.MyMethod2Stub(templ)
}

func (m *MyInterfaceMock) MyMethod3(param1 File) struct{ MyInterface } {
	m.MyMethod3Called++
	return m.MyMethod3Stub(param1)
}
