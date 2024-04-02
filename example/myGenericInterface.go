package example

import "github.com/nicheinc/mock/example/internal"

// MyGenericInterface is a sample generic interface with a complex type
// parameter list.
//
//go:generate mock -o myGenericInterface_mock.go MyGenericInterface
type MyGenericInterface[T interface{ byte | internal.Internal }, U any] interface {
	T() T
	U() U
}
