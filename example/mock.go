package example

import (
	"fmt"
	"html/template"
	. "os"
	renamed "text/template"
)

type MyInterfaceMock struct {
	BlankParamStub                           func(_ string)
	BlankParamCalled                         int
	BlankReturnStub                          func() (_ error)
	BlankReturnCalled                        int
	BlankVariadicParamStub                   func(_ ...string)
	BlankVariadicParamCalled                 int
	DotImportParamStub                       func(file File)
	DotImportParamCalled                     int
	DotImportReturnStub                      func() (file File)
	DotImportReturnCalled                    int
	DotImportVariadicParamStub               func(files ...File)
	DotImportVariadicParamCalled             int
	EmbeddedInterfaceParamStub               func(intf interface{ fmt.Stringer })
	EmbeddedInterfaceParamCalled             int
	EmbeddedInterfaceReturnStub              func() (intf interface{ fmt.Stringer })
	EmbeddedInterfaceReturnCalled            int
	EmbeddedStructParamStub                  func(obj struct{ int })
	EmbeddedStructParamCalled                int
	EmbeddedStructReturnStub                 func() (obj struct{ int })
	EmbeddedStructReturnCalled               int
	EmbeddedStructVariadicParamStub          func(objs ...struct{ int })
	EmbeddedStructVariadicParamCalled        int
	EmptyInterfaceParamStub                  func(intf interface{})
	EmptyInterfaceParamCalled                int
	EmptyInterfaceReturnStub                 func() (intf interface{})
	EmptyInterfaceReturnCalled               int
	EmptyInterfaceVariadicParamStub          func(intf ...interface{})
	EmptyInterfaceVariadicParamCalled        int
	ImportedParamStub                        func(tmpl template.Template)
	ImportedParamCalled                      int
	ImportedVariadicParamStub                func(tmpl ...template.Template)
	ImportedVariadicParamCalled              int
	InterfaceParamStub                       func(intf interface{ MyFunc(num int) error })
	InterfaceParamCalled                     int
	InterfaceReturnStub                      func() (intf interface{ MyFunc(num int) error })
	InterfaceReturnCalled                    int
	InterfaceVariadicFuncParamStub           func(intf interface{ MyFunc(nums ...int) error })
	InterfaceVariadicFuncParamCalled         int
	InterfaceVariadicFuncReturnStub          func() (intf interface{ MyFunc(nums ...int) error })
	InterfaceVariadicFuncReturnCalled        int
	InterfaceVariadicFuncVariadicParamStub   func(intf ...interface{ MyFunc(nums ...int) error })
	InterfaceVariadicFuncVariadicParamCalled int
	InterfaceVariadicParamStub               func(intf ...interface{ MyFunc(num int) error })
	InterfaceVariadicParamCalled             int
	MultipleUnnamedReturnStub                func() (int, error)
	MultipleUnnamedReturnCalled              int
	NamedParamStub                           func(str string)
	NamedParamCalled                         int
	NamedReturnStub                          func() (err error)
	NamedReturnCalled                        int
	NamedVariadicParamStub                   func(strs ...string)
	NamedVariadicParamCalled                 int
	NoParamsOrReturnStub                     func()
	NoParamsOrReturnCalled                   int
	RenamedImportParamStub                   func(tmpl renamed.Template)
	RenamedImportParamCalled                 int
	RenamedImportReturnStub                  func() (tmpl renamed.Template)
	RenamedImportReturnCalled                int
	RenamedImportVariadicParamStub           func(tmpls ...renamed.Template)
	RenamedImportVariadicParamCalled         int
	SameTypeNamedParamsStub                  func(str1 string, str2 string)
	SameTypeNamedParamsCalled                int
	SameTypeNamedReturnStub                  func() (err1 error, err2 error)
	SameTypeNamedReturnCalled                int
	SelfReferentialParamStub                 func(intf MyInterface)
	SelfReferentialParamCalled               int
	SelfReferentialReturnStub                func() (intf MyInterface)
	SelfReferentialReturnCalled              int
	SelfReferentialVariadicParamStub         func(intf ...MyInterface)
	SelfReferentialVariadicParamCalled       int
	StructParamStub                          func(obj struct{ num int })
	StructParamCalled                        int
	StructReturnStub                         func() (obj struct{ num int })
	StructReturnCalled                       int
	StructVariadicParamStub                  func(objs ...struct{ num int })
	StructVariadicParamCalled                int
	UnnamedParamStub                         func(string)
	UnnamedParamCalled                       int
	UnnamedReturnStub                        func() error
	UnnamedReturnCalled                      int
	UnnamedVariadicParamStub                 func(...string)
	UnnamedVariadicParamCalled               int
}

var _ MyInterface = &MyInterfaceMock{}

func (m *MyInterfaceMock) BlankParam(param1 string) {
	m.BlankParamCalled++
	m.BlankParamStub(param1)
}

func (m *MyInterfaceMock) BlankReturn() (_ error) {
	m.BlankReturnCalled++
	return m.BlankReturnStub()
}

func (m *MyInterfaceMock) BlankVariadicParam(param1 ...string) {
	m.BlankVariadicParamCalled++
	m.BlankVariadicParamStub(param1...)
}

func (m *MyInterfaceMock) DotImportParam(file File) {
	m.DotImportParamCalled++
	m.DotImportParamStub(file)
}

func (m *MyInterfaceMock) DotImportReturn() (file File) {
	m.DotImportReturnCalled++
	return m.DotImportReturnStub()
}

func (m *MyInterfaceMock) DotImportVariadicParam(files ...File) {
	m.DotImportVariadicParamCalled++
	m.DotImportVariadicParamStub(files...)
}

func (m *MyInterfaceMock) EmbeddedInterfaceParam(intf interface{ fmt.Stringer }) {
	m.EmbeddedInterfaceParamCalled++
	m.EmbeddedInterfaceParamStub(intf)
}

func (m *MyInterfaceMock) EmbeddedInterfaceReturn() (intf interface{ fmt.Stringer }) {
	m.EmbeddedInterfaceReturnCalled++
	return m.EmbeddedInterfaceReturnStub()
}

func (m *MyInterfaceMock) EmbeddedStructParam(obj struct{ int }) {
	m.EmbeddedStructParamCalled++
	m.EmbeddedStructParamStub(obj)
}

func (m *MyInterfaceMock) EmbeddedStructReturn() (obj struct{ int }) {
	m.EmbeddedStructReturnCalled++
	return m.EmbeddedStructReturnStub()
}

func (m *MyInterfaceMock) EmbeddedStructVariadicParam(objs ...struct{ int }) {
	m.EmbeddedStructVariadicParamCalled++
	m.EmbeddedStructVariadicParamStub(objs...)
}

func (m *MyInterfaceMock) EmptyInterfaceParam(intf interface{}) {
	m.EmptyInterfaceParamCalled++
	m.EmptyInterfaceParamStub(intf)
}

func (m *MyInterfaceMock) EmptyInterfaceReturn() (intf interface{}) {
	m.EmptyInterfaceReturnCalled++
	return m.EmptyInterfaceReturnStub()
}

func (m *MyInterfaceMock) EmptyInterfaceVariadicParam(intf ...interface{}) {
	m.EmptyInterfaceVariadicParamCalled++
	m.EmptyInterfaceVariadicParamStub(intf...)
}

func (m *MyInterfaceMock) ImportedParam(tmpl template.Template) {
	m.ImportedParamCalled++
	m.ImportedParamStub(tmpl)
}

func (m *MyInterfaceMock) ImportedVariadicParam(tmpl ...template.Template) {
	m.ImportedVariadicParamCalled++
	m.ImportedVariadicParamStub(tmpl...)
}

func (m *MyInterfaceMock) InterfaceParam(intf interface{ MyFunc(num int) error }) {
	m.InterfaceParamCalled++
	m.InterfaceParamStub(intf)
}

func (m *MyInterfaceMock) InterfaceReturn() (intf interface{ MyFunc(num int) error }) {
	m.InterfaceReturnCalled++
	return m.InterfaceReturnStub()
}

func (m *MyInterfaceMock) InterfaceVariadicFuncParam(intf interface{ MyFunc(nums ...int) error }) {
	m.InterfaceVariadicFuncParamCalled++
	m.InterfaceVariadicFuncParamStub(intf)
}

func (m *MyInterfaceMock) InterfaceVariadicFuncReturn() (intf interface{ MyFunc(nums ...int) error }) {
	m.InterfaceVariadicFuncReturnCalled++
	return m.InterfaceVariadicFuncReturnStub()
}

func (m *MyInterfaceMock) InterfaceVariadicFuncVariadicParam(intf ...interface{ MyFunc(nums ...int) error }) {
	m.InterfaceVariadicFuncVariadicParamCalled++
	m.InterfaceVariadicFuncVariadicParamStub(intf...)
}

func (m *MyInterfaceMock) InterfaceVariadicParam(intf ...interface{ MyFunc(num int) error }) {
	m.InterfaceVariadicParamCalled++
	m.InterfaceVariadicParamStub(intf...)
}

func (m *MyInterfaceMock) MultipleUnnamedReturn() (int, error) {
	m.MultipleUnnamedReturnCalled++
	return m.MultipleUnnamedReturnStub()
}

func (m *MyInterfaceMock) NamedParam(str string) {
	m.NamedParamCalled++
	m.NamedParamStub(str)
}

func (m *MyInterfaceMock) NamedReturn() (err error) {
	m.NamedReturnCalled++
	return m.NamedReturnStub()
}

func (m *MyInterfaceMock) NamedVariadicParam(strs ...string) {
	m.NamedVariadicParamCalled++
	m.NamedVariadicParamStub(strs...)
}

func (m *MyInterfaceMock) NoParamsOrReturn() {
	m.NoParamsOrReturnCalled++
	m.NoParamsOrReturnStub()
}

func (m *MyInterfaceMock) RenamedImportParam(tmpl renamed.Template) {
	m.RenamedImportParamCalled++
	m.RenamedImportParamStub(tmpl)
}

func (m *MyInterfaceMock) RenamedImportReturn() (tmpl renamed.Template) {
	m.RenamedImportReturnCalled++
	return m.RenamedImportReturnStub()
}

func (m *MyInterfaceMock) RenamedImportVariadicParam(tmpls ...renamed.Template) {
	m.RenamedImportVariadicParamCalled++
	m.RenamedImportVariadicParamStub(tmpls...)
}

func (m *MyInterfaceMock) SameTypeNamedParams(str1 string, str2 string) {
	m.SameTypeNamedParamsCalled++
	m.SameTypeNamedParamsStub(str1, str2)
}

func (m *MyInterfaceMock) SameTypeNamedReturn() (err1 error, err2 error) {
	m.SameTypeNamedReturnCalled++
	return m.SameTypeNamedReturnStub()
}

func (m *MyInterfaceMock) SelfReferentialParam(intf MyInterface) {
	m.SelfReferentialParamCalled++
	m.SelfReferentialParamStub(intf)
}

func (m *MyInterfaceMock) SelfReferentialReturn() (intf MyInterface) {
	m.SelfReferentialReturnCalled++
	return m.SelfReferentialReturnStub()
}

func (m *MyInterfaceMock) SelfReferentialVariadicParam(intf ...MyInterface) {
	m.SelfReferentialVariadicParamCalled++
	m.SelfReferentialVariadicParamStub(intf...)
}

func (m *MyInterfaceMock) StructParam(obj struct{ num int }) {
	m.StructParamCalled++
	m.StructParamStub(obj)
}

func (m *MyInterfaceMock) StructReturn() (obj struct{ num int }) {
	m.StructReturnCalled++
	return m.StructReturnStub()
}

func (m *MyInterfaceMock) StructVariadicParam(objs ...struct{ num int }) {
	m.StructVariadicParamCalled++
	m.StructVariadicParamStub(objs...)
}

func (m *MyInterfaceMock) UnnamedParam(param1 string) {
	m.UnnamedParamCalled++
	m.UnnamedParamStub(param1)
}

func (m *MyInterfaceMock) UnnamedReturn() error {
	m.UnnamedReturnCalled++
	return m.UnnamedReturnStub()
}

func (m *MyInterfaceMock) UnnamedVariadicParam(param1 ...string) {
	m.UnnamedVariadicParamCalled++
	m.UnnamedVariadicParamStub(param1...)
}
