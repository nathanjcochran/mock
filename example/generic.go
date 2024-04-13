package example

import "github.com/nathanjcochran/mock/example/internal"

// Generic is a sample generic interface with a complex type
// parameter list.
//
//go:generate mock -o generic_mock.go Generic
type Generic[T interface{ byte | internal.Internal }, U any] interface {
	GetT() T
	GetU() U
}
