package example

import (
	"fmt"
	"html/template"
	. "os"
	"sync/atomic"
	"testing"
	renamed "text/template"

	"github.com/nathanjcochran/mock/example/internal"
)

// ExampleMock is a mock implementation of the Example
// interface.
type ExampleMock struct {
	T                                        *testing.T
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
	InternalTypeParamStub                    func(internal internal.Internal)
	InternalTypeParamCalled                  int32
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
	SelfReferentialParamStub                 func(intf Example)
	SelfReferentialParamCalled               int32
	SelfReferentialVariadicParamStub         func(intf ...Example)
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
	SelfReferentialReturnStub                func() (intf Example)
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

// Verify that *ExampleMock implements Example.
var _ Example = &ExampleMock{}

// NoParamsOrReturn is a stub for the Example.NoParamsOrReturn
// method that records the number of times it has been called.
func (m *ExampleMock) NoParamsOrReturn() {
	atomic.AddInt32(&m.NoParamsOrReturnCalled, 1)
	if m.NoParamsOrReturnStub == nil {
		if m.T != nil {
			m.T.Error("NoParamsOrReturnStub is nil")
		}
		panic("NoParamsOrReturn unimplemented")
	}
	m.NoParamsOrReturnStub()
}

// UnnamedParam is a stub for the Example.UnnamedParam
// method that records the number of times it has been called.
func (m *ExampleMock) UnnamedParam(param1 string) {
	atomic.AddInt32(&m.UnnamedParamCalled, 1)
	if m.UnnamedParamStub == nil {
		if m.T != nil {
			m.T.Error("UnnamedParamStub is nil")
		}
		panic("UnnamedParam unimplemented")
	}
	m.UnnamedParamStub(param1)
}

// UnnamedVariadicParam is a stub for the Example.UnnamedVariadicParam
// method that records the number of times it has been called.
func (m *ExampleMock) UnnamedVariadicParam(param1 ...string) {
	atomic.AddInt32(&m.UnnamedVariadicParamCalled, 1)
	if m.UnnamedVariadicParamStub == nil {
		if m.T != nil {
			m.T.Error("UnnamedVariadicParamStub is nil")
		}
		panic("UnnamedVariadicParam unimplemented")
	}
	m.UnnamedVariadicParamStub(param1...)
}

// BlankParam is a stub for the Example.BlankParam
// method that records the number of times it has been called.
func (m *ExampleMock) BlankParam(param1 string) {
	atomic.AddInt32(&m.BlankParamCalled, 1)
	if m.BlankParamStub == nil {
		if m.T != nil {
			m.T.Error("BlankParamStub is nil")
		}
		panic("BlankParam unimplemented")
	}
	m.BlankParamStub(param1)
}

// BlankVariadicParam is a stub for the Example.BlankVariadicParam
// method that records the number of times it has been called.
func (m *ExampleMock) BlankVariadicParam(param1 ...string) {
	atomic.AddInt32(&m.BlankVariadicParamCalled, 1)
	if m.BlankVariadicParamStub == nil {
		if m.T != nil {
			m.T.Error("BlankVariadicParamStub is nil")
		}
		panic("BlankVariadicParam unimplemented")
	}
	m.BlankVariadicParamStub(param1...)
}

// NamedParam is a stub for the Example.NamedParam
// method that records the number of times it has been called.
func (m *ExampleMock) NamedParam(str string) {
	atomic.AddInt32(&m.NamedParamCalled, 1)
	if m.NamedParamStub == nil {
		if m.T != nil {
			m.T.Error("NamedParamStub is nil")
		}
		panic("NamedParam unimplemented")
	}
	m.NamedParamStub(str)
}

// NamedVariadicParam is a stub for the Example.NamedVariadicParam
// method that records the number of times it has been called.
func (m *ExampleMock) NamedVariadicParam(strs ...string) {
	atomic.AddInt32(&m.NamedVariadicParamCalled, 1)
	if m.NamedVariadicParamStub == nil {
		if m.T != nil {
			m.T.Error("NamedVariadicParamStub is nil")
		}
		panic("NamedVariadicParam unimplemented")
	}
	m.NamedVariadicParamStub(strs...)
}

// SameTypeNamedParams is a stub for the Example.SameTypeNamedParams
// method that records the number of times it has been called.
func (m *ExampleMock) SameTypeNamedParams(str1 string, str2 string) {
	atomic.AddInt32(&m.SameTypeNamedParamsCalled, 1)
	if m.SameTypeNamedParamsStub == nil {
		if m.T != nil {
			m.T.Error("SameTypeNamedParamsStub is nil")
		}
		panic("SameTypeNamedParams unimplemented")
	}
	m.SameTypeNamedParamsStub(str1, str2)
}

// InternalTypeParam is a stub for the Example.InternalTypeParam
// method that records the number of times it has been called.
func (m *ExampleMock) InternalTypeParam(internal internal.Internal) {
	atomic.AddInt32(&m.InternalTypeParamCalled, 1)
	if m.InternalTypeParamStub == nil {
		if m.T != nil {
			m.T.Error("InternalTypeParamStub is nil")
		}
		panic("InternalTypeParam unimplemented")
	}
	m.InternalTypeParamStub(internal)
}

// ImportedParam is a stub for the Example.ImportedParam
// method that records the number of times it has been called.
func (m *ExampleMock) ImportedParam(tmpl template.Template) {
	atomic.AddInt32(&m.ImportedParamCalled, 1)
	if m.ImportedParamStub == nil {
		if m.T != nil {
			m.T.Error("ImportedParamStub is nil")
		}
		panic("ImportedParam unimplemented")
	}
	m.ImportedParamStub(tmpl)
}

// ImportedVariadicParam is a stub for the Example.ImportedVariadicParam
// method that records the number of times it has been called.
func (m *ExampleMock) ImportedVariadicParam(tmpl ...template.Template) {
	atomic.AddInt32(&m.ImportedVariadicParamCalled, 1)
	if m.ImportedVariadicParamStub == nil {
		if m.T != nil {
			m.T.Error("ImportedVariadicParamStub is nil")
		}
		panic("ImportedVariadicParam unimplemented")
	}
	m.ImportedVariadicParamStub(tmpl...)
}

// RenamedImportParam is a stub for the Example.RenamedImportParam
// method that records the number of times it has been called.
func (m *ExampleMock) RenamedImportParam(tmpl renamed.Template) {
	atomic.AddInt32(&m.RenamedImportParamCalled, 1)
	if m.RenamedImportParamStub == nil {
		if m.T != nil {
			m.T.Error("RenamedImportParamStub is nil")
		}
		panic("RenamedImportParam unimplemented")
	}
	m.RenamedImportParamStub(tmpl)
}

// RenamedImportVariadicParam is a stub for the Example.RenamedImportVariadicParam
// method that records the number of times it has been called.
func (m *ExampleMock) RenamedImportVariadicParam(tmpls ...renamed.Template) {
	atomic.AddInt32(&m.RenamedImportVariadicParamCalled, 1)
	if m.RenamedImportVariadicParamStub == nil {
		if m.T != nil {
			m.T.Error("RenamedImportVariadicParamStub is nil")
		}
		panic("RenamedImportVariadicParam unimplemented")
	}
	m.RenamedImportVariadicParamStub(tmpls...)
}

// DotImportParam is a stub for the Example.DotImportParam
// method that records the number of times it has been called.
func (m *ExampleMock) DotImportParam(file File) {
	atomic.AddInt32(&m.DotImportParamCalled, 1)
	if m.DotImportParamStub == nil {
		if m.T != nil {
			m.T.Error("DotImportParamStub is nil")
		}
		panic("DotImportParam unimplemented")
	}
	m.DotImportParamStub(file)
}

// DotImportVariadicParam is a stub for the Example.DotImportVariadicParam
// method that records the number of times it has been called.
func (m *ExampleMock) DotImportVariadicParam(files ...File) {
	atomic.AddInt32(&m.DotImportVariadicParamCalled, 1)
	if m.DotImportVariadicParamStub == nil {
		if m.T != nil {
			m.T.Error("DotImportVariadicParamStub is nil")
		}
		panic("DotImportVariadicParam unimplemented")
	}
	m.DotImportVariadicParamStub(files...)
}

// SelfReferentialParam is a stub for the Example.SelfReferentialParam
// method that records the number of times it has been called.
func (m *ExampleMock) SelfReferentialParam(intf Example) {
	atomic.AddInt32(&m.SelfReferentialParamCalled, 1)
	if m.SelfReferentialParamStub == nil {
		if m.T != nil {
			m.T.Error("SelfReferentialParamStub is nil")
		}
		panic("SelfReferentialParam unimplemented")
	}
	m.SelfReferentialParamStub(intf)
}

// SelfReferentialVariadicParam is a stub for the Example.SelfReferentialVariadicParam
// method that records the number of times it has been called.
func (m *ExampleMock) SelfReferentialVariadicParam(intf ...Example) {
	atomic.AddInt32(&m.SelfReferentialVariadicParamCalled, 1)
	if m.SelfReferentialVariadicParamStub == nil {
		if m.T != nil {
			m.T.Error("SelfReferentialVariadicParamStub is nil")
		}
		panic("SelfReferentialVariadicParam unimplemented")
	}
	m.SelfReferentialVariadicParamStub(intf...)
}

// StructParam is a stub for the Example.StructParam
// method that records the number of times it has been called.
func (m *ExampleMock) StructParam(obj struct{ num int }) {
	atomic.AddInt32(&m.StructParamCalled, 1)
	if m.StructParamStub == nil {
		if m.T != nil {
			m.T.Error("StructParamStub is nil")
		}
		panic("StructParam unimplemented")
	}
	m.StructParamStub(obj)
}

// StructVariadicParam is a stub for the Example.StructVariadicParam
// method that records the number of times it has been called.
func (m *ExampleMock) StructVariadicParam(objs ...struct{ num int }) {
	atomic.AddInt32(&m.StructVariadicParamCalled, 1)
	if m.StructVariadicParamStub == nil {
		if m.T != nil {
			m.T.Error("StructVariadicParamStub is nil")
		}
		panic("StructVariadicParam unimplemented")
	}
	m.StructVariadicParamStub(objs...)
}

// EmbeddedStructParam is a stub for the Example.EmbeddedStructParam
// method that records the number of times it has been called.
func (m *ExampleMock) EmbeddedStructParam(obj struct{ int }) {
	atomic.AddInt32(&m.EmbeddedStructParamCalled, 1)
	if m.EmbeddedStructParamStub == nil {
		if m.T != nil {
			m.T.Error("EmbeddedStructParamStub is nil")
		}
		panic("EmbeddedStructParam unimplemented")
	}
	m.EmbeddedStructParamStub(obj)
}

// EmbeddedStructVariadicParam is a stub for the Example.EmbeddedStructVariadicParam
// method that records the number of times it has been called.
func (m *ExampleMock) EmbeddedStructVariadicParam(objs ...struct{ int }) {
	atomic.AddInt32(&m.EmbeddedStructVariadicParamCalled, 1)
	if m.EmbeddedStructVariadicParamStub == nil {
		if m.T != nil {
			m.T.Error("EmbeddedStructVariadicParamStub is nil")
		}
		panic("EmbeddedStructVariadicParam unimplemented")
	}
	m.EmbeddedStructVariadicParamStub(objs...)
}

// EmptyInterfaceParam is a stub for the Example.EmptyInterfaceParam
// method that records the number of times it has been called.
func (m *ExampleMock) EmptyInterfaceParam(intf interface{}) {
	atomic.AddInt32(&m.EmptyInterfaceParamCalled, 1)
	if m.EmptyInterfaceParamStub == nil {
		if m.T != nil {
			m.T.Error("EmptyInterfaceParamStub is nil")
		}
		panic("EmptyInterfaceParam unimplemented")
	}
	m.EmptyInterfaceParamStub(intf)
}

// EmptyInterfaceVariadicParam is a stub for the Example.EmptyInterfaceVariadicParam
// method that records the number of times it has been called.
func (m *ExampleMock) EmptyInterfaceVariadicParam(intf ...interface{}) {
	atomic.AddInt32(&m.EmptyInterfaceVariadicParamCalled, 1)
	if m.EmptyInterfaceVariadicParamStub == nil {
		if m.T != nil {
			m.T.Error("EmptyInterfaceVariadicParamStub is nil")
		}
		panic("EmptyInterfaceVariadicParam unimplemented")
	}
	m.EmptyInterfaceVariadicParamStub(intf...)
}

// InterfaceParam is a stub for the Example.InterfaceParam
// method that records the number of times it has been called.
func (m *ExampleMock) InterfaceParam(intf interface{ MyFunc(num int) error }) {
	atomic.AddInt32(&m.InterfaceParamCalled, 1)
	if m.InterfaceParamStub == nil {
		if m.T != nil {
			m.T.Error("InterfaceParamStub is nil")
		}
		panic("InterfaceParam unimplemented")
	}
	m.InterfaceParamStub(intf)
}

// InterfaceVariadicParam is a stub for the Example.InterfaceVariadicParam
// method that records the number of times it has been called.
func (m *ExampleMock) InterfaceVariadicParam(intf ...interface{ MyFunc(num int) error }) {
	atomic.AddInt32(&m.InterfaceVariadicParamCalled, 1)
	if m.InterfaceVariadicParamStub == nil {
		if m.T != nil {
			m.T.Error("InterfaceVariadicParamStub is nil")
		}
		panic("InterfaceVariadicParam unimplemented")
	}
	m.InterfaceVariadicParamStub(intf...)
}

// InterfaceVariadicFuncParam is a stub for the Example.InterfaceVariadicFuncParam
// method that records the number of times it has been called.
func (m *ExampleMock) InterfaceVariadicFuncParam(intf interface{ MyFunc(nums ...int) error }) {
	atomic.AddInt32(&m.InterfaceVariadicFuncParamCalled, 1)
	if m.InterfaceVariadicFuncParamStub == nil {
		if m.T != nil {
			m.T.Error("InterfaceVariadicFuncParamStub is nil")
		}
		panic("InterfaceVariadicFuncParam unimplemented")
	}
	m.InterfaceVariadicFuncParamStub(intf)
}

// InterfaceVariadicFuncVariadicParam is a stub for the Example.InterfaceVariadicFuncVariadicParam
// method that records the number of times it has been called.
func (m *ExampleMock) InterfaceVariadicFuncVariadicParam(intf ...interface{ MyFunc(nums ...int) error }) {
	atomic.AddInt32(&m.InterfaceVariadicFuncVariadicParamCalled, 1)
	if m.InterfaceVariadicFuncVariadicParamStub == nil {
		if m.T != nil {
			m.T.Error("InterfaceVariadicFuncVariadicParamStub is nil")
		}
		panic("InterfaceVariadicFuncVariadicParam unimplemented")
	}
	m.InterfaceVariadicFuncVariadicParamStub(intf...)
}

// EmbeddedInterfaceParam is a stub for the Example.EmbeddedInterfaceParam
// method that records the number of times it has been called.
func (m *ExampleMock) EmbeddedInterfaceParam(intf interface{ fmt.Stringer }) {
	atomic.AddInt32(&m.EmbeddedInterfaceParamCalled, 1)
	if m.EmbeddedInterfaceParamStub == nil {
		if m.T != nil {
			m.T.Error("EmbeddedInterfaceParamStub is nil")
		}
		panic("EmbeddedInterfaceParam unimplemented")
	}
	m.EmbeddedInterfaceParamStub(intf)
}

// UnnamedReturn is a stub for the Example.UnnamedReturn
// method that records the number of times it has been called.
func (m *ExampleMock) UnnamedReturn() error {
	atomic.AddInt32(&m.UnnamedReturnCalled, 1)
	if m.UnnamedReturnStub == nil {
		if m.T != nil {
			m.T.Error("UnnamedReturnStub is nil")
		}
		panic("UnnamedReturn unimplemented")
	}
	return m.UnnamedReturnStub()
}

// MultipleUnnamedReturn is a stub for the Example.MultipleUnnamedReturn
// method that records the number of times it has been called.
func (m *ExampleMock) MultipleUnnamedReturn() (int, error) {
	atomic.AddInt32(&m.MultipleUnnamedReturnCalled, 1)
	if m.MultipleUnnamedReturnStub == nil {
		if m.T != nil {
			m.T.Error("MultipleUnnamedReturnStub is nil")
		}
		panic("MultipleUnnamedReturn unimplemented")
	}
	return m.MultipleUnnamedReturnStub()
}

// BlankReturn is a stub for the Example.BlankReturn
// method that records the number of times it has been called.
func (m *ExampleMock) BlankReturn() (_ error) {
	atomic.AddInt32(&m.BlankReturnCalled, 1)
	if m.BlankReturnStub == nil {
		if m.T != nil {
			m.T.Error("BlankReturnStub is nil")
		}
		panic("BlankReturn unimplemented")
	}
	return m.BlankReturnStub()
}

// NamedReturn is a stub for the Example.NamedReturn
// method that records the number of times it has been called.
func (m *ExampleMock) NamedReturn() (err error) {
	atomic.AddInt32(&m.NamedReturnCalled, 1)
	if m.NamedReturnStub == nil {
		if m.T != nil {
			m.T.Error("NamedReturnStub is nil")
		}
		panic("NamedReturn unimplemented")
	}
	return m.NamedReturnStub()
}

// SameTypeNamedReturn is a stub for the Example.SameTypeNamedReturn
// method that records the number of times it has been called.
func (m *ExampleMock) SameTypeNamedReturn() (err1 error, err2 error) {
	atomic.AddInt32(&m.SameTypeNamedReturnCalled, 1)
	if m.SameTypeNamedReturnStub == nil {
		if m.T != nil {
			m.T.Error("SameTypeNamedReturnStub is nil")
		}
		panic("SameTypeNamedReturn unimplemented")
	}
	return m.SameTypeNamedReturnStub()
}

// RenamedImportReturn is a stub for the Example.RenamedImportReturn
// method that records the number of times it has been called.
func (m *ExampleMock) RenamedImportReturn() (tmpl renamed.Template) {
	atomic.AddInt32(&m.RenamedImportReturnCalled, 1)
	if m.RenamedImportReturnStub == nil {
		if m.T != nil {
			m.T.Error("RenamedImportReturnStub is nil")
		}
		panic("RenamedImportReturn unimplemented")
	}
	return m.RenamedImportReturnStub()
}

// DotImportReturn is a stub for the Example.DotImportReturn
// method that records the number of times it has been called.
func (m *ExampleMock) DotImportReturn() (file File) {
	atomic.AddInt32(&m.DotImportReturnCalled, 1)
	if m.DotImportReturnStub == nil {
		if m.T != nil {
			m.T.Error("DotImportReturnStub is nil")
		}
		panic("DotImportReturn unimplemented")
	}
	return m.DotImportReturnStub()
}

// SelfReferentialReturn is a stub for the Example.SelfReferentialReturn
// method that records the number of times it has been called.
func (m *ExampleMock) SelfReferentialReturn() (intf Example) {
	atomic.AddInt32(&m.SelfReferentialReturnCalled, 1)
	if m.SelfReferentialReturnStub == nil {
		if m.T != nil {
			m.T.Error("SelfReferentialReturnStub is nil")
		}
		panic("SelfReferentialReturn unimplemented")
	}
	return m.SelfReferentialReturnStub()
}

// StructReturn is a stub for the Example.StructReturn
// method that records the number of times it has been called.
func (m *ExampleMock) StructReturn() (obj struct{ num int }) {
	atomic.AddInt32(&m.StructReturnCalled, 1)
	if m.StructReturnStub == nil {
		if m.T != nil {
			m.T.Error("StructReturnStub is nil")
		}
		panic("StructReturn unimplemented")
	}
	return m.StructReturnStub()
}

// EmbeddedStructReturn is a stub for the Example.EmbeddedStructReturn
// method that records the number of times it has been called.
func (m *ExampleMock) EmbeddedStructReturn() (obj struct{ int }) {
	atomic.AddInt32(&m.EmbeddedStructReturnCalled, 1)
	if m.EmbeddedStructReturnStub == nil {
		if m.T != nil {
			m.T.Error("EmbeddedStructReturnStub is nil")
		}
		panic("EmbeddedStructReturn unimplemented")
	}
	return m.EmbeddedStructReturnStub()
}

// EmptyInterfaceReturn is a stub for the Example.EmptyInterfaceReturn
// method that records the number of times it has been called.
func (m *ExampleMock) EmptyInterfaceReturn() (intf interface{}) {
	atomic.AddInt32(&m.EmptyInterfaceReturnCalled, 1)
	if m.EmptyInterfaceReturnStub == nil {
		if m.T != nil {
			m.T.Error("EmptyInterfaceReturnStub is nil")
		}
		panic("EmptyInterfaceReturn unimplemented")
	}
	return m.EmptyInterfaceReturnStub()
}

// InterfaceReturn is a stub for the Example.InterfaceReturn
// method that records the number of times it has been called.
func (m *ExampleMock) InterfaceReturn() (intf interface{ MyFunc(num int) error }) {
	atomic.AddInt32(&m.InterfaceReturnCalled, 1)
	if m.InterfaceReturnStub == nil {
		if m.T != nil {
			m.T.Error("InterfaceReturnStub is nil")
		}
		panic("InterfaceReturn unimplemented")
	}
	return m.InterfaceReturnStub()
}

// InterfaceVariadicFuncReturn is a stub for the Example.InterfaceVariadicFuncReturn
// method that records the number of times it has been called.
func (m *ExampleMock) InterfaceVariadicFuncReturn() (intf interface{ MyFunc(nums ...int) error }) {
	atomic.AddInt32(&m.InterfaceVariadicFuncReturnCalled, 1)
	if m.InterfaceVariadicFuncReturnStub == nil {
		if m.T != nil {
			m.T.Error("InterfaceVariadicFuncReturnStub is nil")
		}
		panic("InterfaceVariadicFuncReturn unimplemented")
	}
	return m.InterfaceVariadicFuncReturnStub()
}

// EmbeddedInterfaceReturn is a stub for the Example.EmbeddedInterfaceReturn
// method that records the number of times it has been called.
func (m *ExampleMock) EmbeddedInterfaceReturn() (intf interface{ fmt.Stringer }) {
	atomic.AddInt32(&m.EmbeddedInterfaceReturnCalled, 1)
	if m.EmbeddedInterfaceReturnStub == nil {
		if m.T != nil {
			m.T.Error("EmbeddedInterfaceReturnStub is nil")
		}
		panic("EmbeddedInterfaceReturn unimplemented")
	}
	return m.EmbeddedInterfaceReturnStub()
}
