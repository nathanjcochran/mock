package example

import (
	"fmt"
	"html/template"
	. "os"
	renamed "text/template"
)

type MyInterfaceMock struct {
	NoParamsOrReturnStub               func()
	NoParamsOrReturnCalled             int
	UnnamedParamStub                   func(string)
	UnnamedParamCalled                 int
	UnnamedVariadicParamStub           func(...string)
	UnnamedVariadicParamCalled         int
	BlankParamStub                     func(_ string)
	BlankParamCalled                   int
	BlankVariadicParamStub             func(_ ...string)
	BlankVariadicParamCalled           int
	NamedParamStub                     func(str string)
	NamedParamCalled                   int
	NamedVariadicParamStub             func(strs ...string)
	NamedVariadicParamCalled           int
	SameTypeNamedParamsStub            func(str1 string, str2 string)
	SameTypeNamedParamsCalled          int
	ImportedParamStub                  func(tmpl template.Template)
	ImportedParamCalled                int
	ImportedVariadicParamStub          func(tmpl ...template.Template)
	ImportedVariadicParamCalled        int
	RenamedImportParamStub             func(tmpl renamed.Template)
	RenamedImportParamCalled           int
	RenamedImportVariadicParamStub     func(tmpls ...renamed.Template)
	RenamedImportVariadicParamCalled   int
	DotImportParamStub                 func(file File)
	DotImportParamCalled               int
	DotImportVariadicParamStub         func(files ...File)
	DotImportVariadicParamCalled       int
	SelfReferentialParamStub           func(intf MyInterface)
	SelfReferentialParamCalled         int
	SelfReferentialVariadicParamStub   func(intf ...MyInterface)
	SelfReferentialVariadicParamCalled int
	StructParamStub                    func(obj struct{ num int })
	StructParamCalled                  int
	StructVariadicParamStub            func(objs ...struct{ num int })
	StructVariadicParamCalled          int
	EmbeddedStructParamStub            func(obj struct{ int })
	EmbeddedStructParamCalled          int
	EmbeddedStructVariadicParamStub    func(objs ...struct{ int })
	EmbeddedStructVariadicParamCalled  int
	EmptyInterfaceParamStub            func(intf interface{})
	EmptyInterfaceParamCalled          int
	EmptyInterfaceVariadicParamStub    func(intf ...interface{})
	EmptyInterfaceVariadicParamCalled  int
	InterfaceParamStub                 func(intf interface {
		MyFunc(num int) error
	})
	InterfaceParamCalled       int
	InterfaceVariadicParamStub func(intf ...interface {
		MyFunc(num int) error
	})
	InterfaceVariadicParamCalled   int
	InterfaceVariadicFuncParamStub func(intf interface {
		MyFunc(nums ...int) error
	})
	InterfaceVariadicFuncParamCalled       int
	InterfaceVariadicFuncVariadicParamStub func(intf ...interface {
		MyFunc(nums ...int) error
	})
	InterfaceVariadicFuncVariadicParamCalled int
	EmbeddedInterfaceParamStub               func(intf interface {
		fmt.Stringer
	})
	EmbeddedInterfaceParamCalled int
	UnnamedReturnStub            func() error
	UnnamedReturnCalled          int
	MultipleUnnamedReturnStub    func() (int, error)
	MultipleUnnamedReturnCalled  int
	BlankReturnStub              func() (_ error)
	BlankReturnCalled            int
	NamedReturnStub              func() (err error)
	NamedReturnCalled            int
	SameTypeNamedReturnStub      func() (err1 error, err2 error)
	SameTypeNamedReturnCalled    int
	RenamedImportReturnStub      func() (tmpl renamed.Template)
	RenamedImportReturnCalled    int
	DotImportReturnStub          func() (file File)
	DotImportReturnCalled        int
	SelfReferentialReturnStub    func() (intf MyInterface)
	SelfReferentialReturnCalled  int
	StructReturnStub             func() (obj struct{ num int })
	StructReturnCalled           int
	EmbeddedStructReturnStub     func() (obj struct{ int })
	EmbeddedStructReturnCalled   int
	EmptyInterfaceReturnStub     func() (intf interface{})
	EmptyInterfaceReturnCalled   int
	InterfaceReturnStub          func() (intf interface {
		MyFunc(num int) error
	})
	InterfaceReturnCalled           int
	InterfaceVariadicFuncReturnStub func() (intf interface {
		MyFunc(nums ...int) error
	})
	InterfaceVariadicFuncReturnCalled int
	EmbeddedInterfaceReturnStub       func() (intf interface {
		fmt.Stringer
	})
	EmbeddedInterfaceReturnCalled int
}

var _ MyInterface = &MyInterfaceMock{}

func (m *MyInterfaceMock) NoParamsOrReturn() {
	m.NoParamsOrReturnCalled++
	m.NoParamsOrReturnStub()
}

func (m *MyInterfaceMock) UnnamedParam(param1 string) {
	m.UnnamedParamCalled++
	m.UnnamedParamStub(param1)
}

func (m *MyInterfaceMock) UnnamedVariadicParam(param1 ...string) {
	m.UnnamedVariadicParamCalled++
	m.UnnamedVariadicParamStub(param1...)
}

func (m *MyInterfaceMock) BlankParam(param1 string) {
	m.BlankParamCalled++
	m.BlankParamStub(param1)
}

func (m *MyInterfaceMock) BlankVariadicParam(param1 ...string) {
	m.BlankVariadicParamCalled++
	m.BlankVariadicParamStub(param1...)
}

func (m *MyInterfaceMock) NamedParam(str string) {
	m.NamedParamCalled++
	m.NamedParamStub(str)
}

func (m *MyInterfaceMock) NamedVariadicParam(strs ...string) {
	m.NamedVariadicParamCalled++
	m.NamedVariadicParamStub(strs...)
}

func (m *MyInterfaceMock) SameTypeNamedParams(str1 string, str2 string) {
	m.SameTypeNamedParamsCalled++
	m.SameTypeNamedParamsStub(str1, str2)
}

func (m *MyInterfaceMock) ImportedParam(tmpl template.Template) {
	m.ImportedParamCalled++
	m.ImportedParamStub(tmpl)
}

func (m *MyInterfaceMock) ImportedVariadicParam(tmpl ...template.Template) {
	m.ImportedVariadicParamCalled++
	m.ImportedVariadicParamStub(tmpl...)
}

func (m *MyInterfaceMock) RenamedImportParam(tmpl renamed.Template) {
	m.RenamedImportParamCalled++
	m.RenamedImportParamStub(tmpl)
}

func (m *MyInterfaceMock) RenamedImportVariadicParam(tmpls ...renamed.Template) {
	m.RenamedImportVariadicParamCalled++
	m.RenamedImportVariadicParamStub(tmpls...)
}

func (m *MyInterfaceMock) DotImportParam(file File) {
	m.DotImportParamCalled++
	m.DotImportParamStub(file)
}

func (m *MyInterfaceMock) DotImportVariadicParam(files ...File) {
	m.DotImportVariadicParamCalled++
	m.DotImportVariadicParamStub(files...)
}

func (m *MyInterfaceMock) SelfReferentialParam(intf MyInterface) {
	m.SelfReferentialParamCalled++
	m.SelfReferentialParamStub(intf)
}

func (m *MyInterfaceMock) SelfReferentialVariadicParam(intf ...MyInterface) {
	m.SelfReferentialVariadicParamCalled++
	m.SelfReferentialVariadicParamStub(intf...)
}

func (m *MyInterfaceMock) StructParam(obj struct{ num int }) {
	m.StructParamCalled++
	m.StructParamStub(obj)
}

func (m *MyInterfaceMock) StructVariadicParam(objs ...struct{ num int }) {
	m.StructVariadicParamCalled++
	m.StructVariadicParamStub(objs...)
}

func (m *MyInterfaceMock) EmbeddedStructParam(obj struct{ int }) {
	m.EmbeddedStructParamCalled++
	m.EmbeddedStructParamStub(obj)
}

func (m *MyInterfaceMock) EmbeddedStructVariadicParam(objs ...struct{ int }) {
	m.EmbeddedStructVariadicParamCalled++
	m.EmbeddedStructVariadicParamStub(objs...)
}

func (m *MyInterfaceMock) EmptyInterfaceParam(intf interface{}) {
	m.EmptyInterfaceParamCalled++
	m.EmptyInterfaceParamStub(intf)
}

func (m *MyInterfaceMock) EmptyInterfaceVariadicParam(intf ...interface{}) {
	m.EmptyInterfaceVariadicParamCalled++
	m.EmptyInterfaceVariadicParamStub(intf...)
}

func (m *MyInterfaceMock) InterfaceParam(intf interface {
	MyFunc(num int) error
}) {
	m.InterfaceParamCalled++
	m.InterfaceParamStub(intf)
}

func (m *MyInterfaceMock) InterfaceVariadicParam(intf ...interface {
	MyFunc(num int) error
}) {
	m.InterfaceVariadicParamCalled++
	m.InterfaceVariadicParamStub(intf...)
}

func (m *MyInterfaceMock) InterfaceVariadicFuncParam(intf interface {
	MyFunc(nums ...int) error
}) {
	m.InterfaceVariadicFuncParamCalled++
	m.InterfaceVariadicFuncParamStub(intf)
}

func (m *MyInterfaceMock) InterfaceVariadicFuncVariadicParam(intf ...interface {
	MyFunc(nums ...int) error
}) {
	m.InterfaceVariadicFuncVariadicParamCalled++
	m.InterfaceVariadicFuncVariadicParamStub(intf...)
}

func (m *MyInterfaceMock) EmbeddedInterfaceParam(intf interface {
	fmt.Stringer
}) {
	m.EmbeddedInterfaceParamCalled++
	m.EmbeddedInterfaceParamStub(intf)
}

func (m *MyInterfaceMock) UnnamedReturn() error {
	m.UnnamedReturnCalled++
	return m.UnnamedReturnStub()
}

func (m *MyInterfaceMock) MultipleUnnamedReturn() (int, error) {
	m.MultipleUnnamedReturnCalled++
	return m.MultipleUnnamedReturnStub()
}

func (m *MyInterfaceMock) BlankReturn() (_ error) {
	m.BlankReturnCalled++
	return m.BlankReturnStub()
}

func (m *MyInterfaceMock) NamedReturn() (err error) {
	m.NamedReturnCalled++
	return m.NamedReturnStub()
}

func (m *MyInterfaceMock) SameTypeNamedReturn() (err1 error, err2 error) {
	m.SameTypeNamedReturnCalled++
	return m.SameTypeNamedReturnStub()
}

func (m *MyInterfaceMock) RenamedImportReturn() (tmpl renamed.Template) {
	m.RenamedImportReturnCalled++
	return m.RenamedImportReturnStub()
}

func (m *MyInterfaceMock) DotImportReturn() (file File) {
	m.DotImportReturnCalled++
	return m.DotImportReturnStub()
}

func (m *MyInterfaceMock) SelfReferentialReturn() (intf MyInterface) {
	m.SelfReferentialReturnCalled++
	return m.SelfReferentialReturnStub()
}

func (m *MyInterfaceMock) StructReturn() (obj struct{ num int }) {
	m.StructReturnCalled++
	return m.StructReturnStub()
}

func (m *MyInterfaceMock) EmbeddedStructReturn() (obj struct{ int }) {
	m.EmbeddedStructReturnCalled++
	return m.EmbeddedStructReturnStub()
}

func (m *MyInterfaceMock) EmptyInterfaceReturn() (intf interface{}) {
	m.EmptyInterfaceReturnCalled++
	return m.EmptyInterfaceReturnStub()
}

func (m *MyInterfaceMock) InterfaceReturn() (intf interface {
	MyFunc(num int) error
}) {
	m.InterfaceReturnCalled++
	return m.InterfaceReturnStub()
}

func (m *MyInterfaceMock) InterfaceVariadicFuncReturn() (intf interface {
	MyFunc(nums ...int) error
}) {
	m.InterfaceVariadicFuncReturnCalled++
	return m.InterfaceVariadicFuncReturnStub()
}

func (m *MyInterfaceMock) EmbeddedInterfaceReturn() (intf interface {
	fmt.Stringer
}) {
	m.EmbeddedInterfaceReturnCalled++
	return m.EmbeddedInterfaceReturnStub()
}
