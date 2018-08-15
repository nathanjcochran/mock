package example

import (
	"fmt"
	"html/template"
	. "os"
	"sync/atomic"
	renamed "text/template"
)

type MyInterfaceMock struct {
	NoParamsOrReturnStub                     func()
	NoParamsOrReturnCalled                   int32
	UnnamedParamStub                         func(string)
	UnnamedParamCalled                       int32
	UnnamedVariadicParamStub                 func(...string)
	UnnamedVariadicParamCalled               int32
	BlankParamStub                           func(_ string)
	BlankParamCalled                         int32
	BlankVariadicParamStub                   func(_ ...string)
	BlankVariadicParamCalled                 int32
	NamedParamStub                           func(str string)
	NamedParamCalled                         int32
	NamedVariadicParamStub                   func(strs ...string)
	NamedVariadicParamCalled                 int32
	SameTypeNamedParamsStub                  func(str1 string, str2 string)
	SameTypeNamedParamsCalled                int32
	ImportedParamStub                        func(tmpl template.Template)
	ImportedParamCalled                      int32
	ImportedVariadicParamStub                func(tmpl ...template.Template)
	ImportedVariadicParamCalled              int32
	RenamedImportParamStub                   func(tmpl renamed.Template)
	RenamedImportParamCalled                 int32
	RenamedImportVariadicParamStub           func(tmpls ...renamed.Template)
	RenamedImportVariadicParamCalled         int32
	DotImportParamStub                       func(file File)
	DotImportParamCalled                     int32
	DotImportVariadicParamStub               func(files ...File)
	DotImportVariadicParamCalled             int32
	SelfReferentialParamStub                 func(intf MyInterface)
	SelfReferentialParamCalled               int32
	SelfReferentialVariadicParamStub         func(intf ...MyInterface)
	SelfReferentialVariadicParamCalled       int32
	StructParamStub                          func(obj struct{ num int })
	StructParamCalled                        int32
	StructVariadicParamStub                  func(objs ...struct{ num int })
	StructVariadicParamCalled                int32
	EmbeddedStructParamStub                  func(obj struct{ int })
	EmbeddedStructParamCalled                int32
	EmbeddedStructVariadicParamStub          func(objs ...struct{ int })
	EmbeddedStructVariadicParamCalled        int32
	EmptyInterfaceParamStub                  func(intf interface{})
	EmptyInterfaceParamCalled                int32
	EmptyInterfaceVariadicParamStub          func(intf ...interface{})
	EmptyInterfaceVariadicParamCalled        int32
	InterfaceParamStub                       func(intf interface{ MyFunc(num int) error })
	InterfaceParamCalled                     int32
	InterfaceVariadicParamStub               func(intf ...interface{ MyFunc(num int) error })
	InterfaceVariadicParamCalled             int32
	InterfaceVariadicFuncParamStub           func(intf interface{ MyFunc(nums ...int) error })
	InterfaceVariadicFuncParamCalled         int32
	InterfaceVariadicFuncVariadicParamStub   func(intf ...interface{ MyFunc(nums ...int) error })
	InterfaceVariadicFuncVariadicParamCalled int32
	EmbeddedInterfaceParamStub               func(intf interface{ fmt.Stringer })
	EmbeddedInterfaceParamCalled             int32
	UnnamedReturnStub                        func() error
	UnnamedReturnCalled                      int32
	MultipleUnnamedReturnStub                func() (int, error)
	MultipleUnnamedReturnCalled              int32
	BlankReturnStub                          func() (_ error)
	BlankReturnCalled                        int32
	NamedReturnStub                          func() (err error)
	NamedReturnCalled                        int32
	SameTypeNamedReturnStub                  func() (err1 error, err2 error)
	SameTypeNamedReturnCalled                int32
	RenamedImportReturnStub                  func() (tmpl renamed.Template)
	RenamedImportReturnCalled                int32
	DotImportReturnStub                      func() (file File)
	DotImportReturnCalled                    int32
	SelfReferentialReturnStub                func() (intf MyInterface)
	SelfReferentialReturnCalled              int32
	StructReturnStub                         func() (obj struct{ num int })
	StructReturnCalled                       int32
	EmbeddedStructReturnStub                 func() (obj struct{ int })
	EmbeddedStructReturnCalled               int32
	EmptyInterfaceReturnStub                 func() (intf interface{})
	EmptyInterfaceReturnCalled               int32
	InterfaceReturnStub                      func() (intf interface{ MyFunc(num int) error })
	InterfaceReturnCalled                    int32
	InterfaceVariadicFuncReturnStub          func() (intf interface{ MyFunc(nums ...int) error })
	InterfaceVariadicFuncReturnCalled        int32
	EmbeddedInterfaceReturnStub              func() (intf interface{ fmt.Stringer })
	EmbeddedInterfaceReturnCalled            int32
}

var _ MyInterface = &MyInterfaceMock{}

func (m *MyInterfaceMock) NoParamsOrReturn() {
	atomic.AddInt32(&m.NoParamsOrReturnCalled, 1)
	m.NoParamsOrReturnStub()
}

func (m *MyInterfaceMock) UnnamedParam(param1 string) {
	atomic.AddInt32(&m.UnnamedParamCalled, 1)
	m.UnnamedParamStub(param1)
}

func (m *MyInterfaceMock) UnnamedVariadicParam(param1 ...string) {
	atomic.AddInt32(&m.UnnamedVariadicParamCalled, 1)
	m.UnnamedVariadicParamStub(param1...)
}

func (m *MyInterfaceMock) BlankParam(param1 string) {
	atomic.AddInt32(&m.BlankParamCalled, 1)
	m.BlankParamStub(param1)
}

func (m *MyInterfaceMock) BlankVariadicParam(param1 ...string) {
	atomic.AddInt32(&m.BlankVariadicParamCalled, 1)
	m.BlankVariadicParamStub(param1...)
}

func (m *MyInterfaceMock) NamedParam(str string) {
	atomic.AddInt32(&m.NamedParamCalled, 1)
	m.NamedParamStub(str)
}

func (m *MyInterfaceMock) NamedVariadicParam(strs ...string) {
	atomic.AddInt32(&m.NamedVariadicParamCalled, 1)
	m.NamedVariadicParamStub(strs...)
}

func (m *MyInterfaceMock) SameTypeNamedParams(str1 string, str2 string) {
	atomic.AddInt32(&m.SameTypeNamedParamsCalled, 1)
	m.SameTypeNamedParamsStub(str1, str2)
}

func (m *MyInterfaceMock) ImportedParam(tmpl template.Template) {
	atomic.AddInt32(&m.ImportedParamCalled, 1)
	m.ImportedParamStub(tmpl)
}

func (m *MyInterfaceMock) ImportedVariadicParam(tmpl ...template.Template) {
	atomic.AddInt32(&m.ImportedVariadicParamCalled, 1)
	m.ImportedVariadicParamStub(tmpl...)
}

func (m *MyInterfaceMock) RenamedImportParam(tmpl renamed.Template) {
	atomic.AddInt32(&m.RenamedImportParamCalled, 1)
	m.RenamedImportParamStub(tmpl)
}

func (m *MyInterfaceMock) RenamedImportVariadicParam(tmpls ...renamed.Template) {
	atomic.AddInt32(&m.RenamedImportVariadicParamCalled, 1)
	m.RenamedImportVariadicParamStub(tmpls...)
}

func (m *MyInterfaceMock) DotImportParam(file File) {
	atomic.AddInt32(&m.DotImportParamCalled, 1)
	m.DotImportParamStub(file)
}

func (m *MyInterfaceMock) DotImportVariadicParam(files ...File) {
	atomic.AddInt32(&m.DotImportVariadicParamCalled, 1)
	m.DotImportVariadicParamStub(files...)
}

func (m *MyInterfaceMock) SelfReferentialParam(intf MyInterface) {
	atomic.AddInt32(&m.SelfReferentialParamCalled, 1)
	m.SelfReferentialParamStub(intf)
}

func (m *MyInterfaceMock) SelfReferentialVariadicParam(intf ...MyInterface) {
	atomic.AddInt32(&m.SelfReferentialVariadicParamCalled, 1)
	m.SelfReferentialVariadicParamStub(intf...)
}

func (m *MyInterfaceMock) StructParam(obj struct{ num int }) {
	atomic.AddInt32(&m.StructParamCalled, 1)
	m.StructParamStub(obj)
}

func (m *MyInterfaceMock) StructVariadicParam(objs ...struct{ num int }) {
	atomic.AddInt32(&m.StructVariadicParamCalled, 1)
	m.StructVariadicParamStub(objs...)
}

func (m *MyInterfaceMock) EmbeddedStructParam(obj struct{ int }) {
	atomic.AddInt32(&m.EmbeddedStructParamCalled, 1)
	m.EmbeddedStructParamStub(obj)
}

func (m *MyInterfaceMock) EmbeddedStructVariadicParam(objs ...struct{ int }) {
	atomic.AddInt32(&m.EmbeddedStructVariadicParamCalled, 1)
	m.EmbeddedStructVariadicParamStub(objs...)
}

func (m *MyInterfaceMock) EmptyInterfaceParam(intf interface{}) {
	atomic.AddInt32(&m.EmptyInterfaceParamCalled, 1)
	m.EmptyInterfaceParamStub(intf)
}

func (m *MyInterfaceMock) EmptyInterfaceVariadicParam(intf ...interface{}) {
	atomic.AddInt32(&m.EmptyInterfaceVariadicParamCalled, 1)
	m.EmptyInterfaceVariadicParamStub(intf...)
}

func (m *MyInterfaceMock) InterfaceParam(intf interface{ MyFunc(num int) error }) {
	atomic.AddInt32(&m.InterfaceParamCalled, 1)
	m.InterfaceParamStub(intf)
}

func (m *MyInterfaceMock) InterfaceVariadicParam(intf ...interface{ MyFunc(num int) error }) {
	atomic.AddInt32(&m.InterfaceVariadicParamCalled, 1)
	m.InterfaceVariadicParamStub(intf...)
}

func (m *MyInterfaceMock) InterfaceVariadicFuncParam(intf interface{ MyFunc(nums ...int) error }) {
	atomic.AddInt32(&m.InterfaceVariadicFuncParamCalled, 1)
	m.InterfaceVariadicFuncParamStub(intf)
}

func (m *MyInterfaceMock) InterfaceVariadicFuncVariadicParam(intf ...interface{ MyFunc(nums ...int) error }) {
	atomic.AddInt32(&m.InterfaceVariadicFuncVariadicParamCalled, 1)
	m.InterfaceVariadicFuncVariadicParamStub(intf...)
}

func (m *MyInterfaceMock) EmbeddedInterfaceParam(intf interface{ fmt.Stringer }) {
	atomic.AddInt32(&m.EmbeddedInterfaceParamCalled, 1)
	m.EmbeddedInterfaceParamStub(intf)
}

func (m *MyInterfaceMock) UnnamedReturn() error {
	atomic.AddInt32(&m.UnnamedReturnCalled, 1)
	return m.UnnamedReturnStub()
}

func (m *MyInterfaceMock) MultipleUnnamedReturn() (int, error) {
	atomic.AddInt32(&m.MultipleUnnamedReturnCalled, 1)
	return m.MultipleUnnamedReturnStub()
}

func (m *MyInterfaceMock) BlankReturn() (_ error) {
	atomic.AddInt32(&m.BlankReturnCalled, 1)
	return m.BlankReturnStub()
}

func (m *MyInterfaceMock) NamedReturn() (err error) {
	atomic.AddInt32(&m.NamedReturnCalled, 1)
	return m.NamedReturnStub()
}

func (m *MyInterfaceMock) SameTypeNamedReturn() (err1 error, err2 error) {
	atomic.AddInt32(&m.SameTypeNamedReturnCalled, 1)
	return m.SameTypeNamedReturnStub()
}

func (m *MyInterfaceMock) RenamedImportReturn() (tmpl renamed.Template) {
	atomic.AddInt32(&m.RenamedImportReturnCalled, 1)
	return m.RenamedImportReturnStub()
}

func (m *MyInterfaceMock) DotImportReturn() (file File) {
	atomic.AddInt32(&m.DotImportReturnCalled, 1)
	return m.DotImportReturnStub()
}

func (m *MyInterfaceMock) SelfReferentialReturn() (intf MyInterface) {
	atomic.AddInt32(&m.SelfReferentialReturnCalled, 1)
	return m.SelfReferentialReturnStub()
}

func (m *MyInterfaceMock) StructReturn() (obj struct{ num int }) {
	atomic.AddInt32(&m.StructReturnCalled, 1)
	return m.StructReturnStub()
}

func (m *MyInterfaceMock) EmbeddedStructReturn() (obj struct{ int }) {
	atomic.AddInt32(&m.EmbeddedStructReturnCalled, 1)
	return m.EmbeddedStructReturnStub()
}

func (m *MyInterfaceMock) EmptyInterfaceReturn() (intf interface{}) {
	atomic.AddInt32(&m.EmptyInterfaceReturnCalled, 1)
	return m.EmptyInterfaceReturnStub()
}

func (m *MyInterfaceMock) InterfaceReturn() (intf interface{ MyFunc(num int) error }) {
	atomic.AddInt32(&m.InterfaceReturnCalled, 1)
	return m.InterfaceReturnStub()
}

func (m *MyInterfaceMock) InterfaceVariadicFuncReturn() (intf interface{ MyFunc(nums ...int) error }) {
	atomic.AddInt32(&m.InterfaceVariadicFuncReturnCalled, 1)
	return m.InterfaceVariadicFuncReturnStub()
}

func (m *MyInterfaceMock) EmbeddedInterfaceReturn() (intf interface{ fmt.Stringer }) {
	atomic.AddInt32(&m.EmbeddedInterfaceReturnCalled, 1)
	return m.EmbeddedInterfaceReturnStub()
}
