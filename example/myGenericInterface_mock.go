package example

import (
	"sync/atomic"

	"github.com/nicheinc/mock/example/internal"
)

// MyGenericInterfaceMock is a mock implementation of the MyGenericInterface
// interface.
type MyGenericInterfaceMock[T interface{ byte | internal.Internal }, U any] struct {
	TStub   func() T
	TCalled int32
	UStub   func() U
	UCalled int32
}

// T is a stub for the MyGenericInterface.T
// method that records the number of times it has been called.
func (m *MyGenericInterfaceMock[T, U]) T() T {
	atomic.AddInt32(&m.TCalled, 1)
	return m.TStub()
}

// U is a stub for the MyGenericInterface.U
// method that records the number of times it has been called.
func (m *MyGenericInterfaceMock[T, U]) U() U {
	atomic.AddInt32(&m.UCalled, 1)
	return m.UStub()
}
