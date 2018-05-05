package example

import (
	"html/template"
	. "os"
	renamed "text/template"
)

//go:generate mock -o mock.go MyInterface
type MyInterface interface {
	NoParamsOrReturn()
	UnnamedParam(string)
	UnnamedVariadicParam(...string)
	BlankParam(_ string)
	BlankVariadicParam(_ ...string)
	NamedParam(str string)
	NamedVariadicParam(strs ...string)
	SameTypeNamedParams(str1, str2 string)
	ImportedParam(tmpl template.Template)
	ImportedVariadicParam(tmpl ...template.Template)
	RenamedImportParam(tmpl renamed.Template)
	RenamedImportVariadicParam(tmpls ...renamed.Template)
	DotImportParam(file File)
	DotImportVariadicParam(files ...File)
	SelfReferentialParam(intf MyInterface)
	SelfReferentialVariadicParam(intf ...MyInterface)
	StructParam(obj struct{ num int })
	StructVariadicParam(objs ...struct{ num int })
	EmbeddedStructParam(obj struct{ int })
	EmbeddedStructVariadicParam(obj ...struct{ int })
	EmptyInterfaceParam(intf interface{})
	InterfaceParam(intf interface {
		MyFunc(num int) error
	})
	EmbeddedInterfaceParam(intf interface {
		MyInterface
	})

	UnnamedReturn() error
	MultipleUnnamedReturn() (int, error)
	BlankReturn() (_ error)
	NamedReturn() (err error)
	SameTypeNamedReturn() (err1, err2 error)
	RenamedImportReturn() (tmpl renamed.Template)
	DotImportReturn() (file File)
	SelfReferentialReturn() (intf MyInterface)
	StructReturn() (obj struct{ num int })
	EmbeddedStructReturn() (obj struct{ int })
	EmptyInterfaceReturn() (intf interface{})
	InterfaceReturn() (intf interface {
		MyFunc(num int) error
	})
	//	EmbeddedInterfaceReturn() (intf interface {
	//		MyInterface
	//	})
}
