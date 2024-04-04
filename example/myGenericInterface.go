package example

import "github.com/nathanjcochran/mock/example/internal"

// MyGenericInterface is a sample generic interface with a complex type
// parameter list.
//
//go:generate mock -o myGenericInterface_mock.go MyGenericInterface
type MyGenericInterface[T interface{ byte | internal.Internal }, U any] interface {
	GetT() T
	GetU() U
}
