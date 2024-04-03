package example

import (
	"sync/atomic"
	"testing"

	"github.com/nicheinc/mock/example/internal"
)

// MyGenericInterfaceMock is a mock implementation of the MyGenericInterface
// interface.
type MyGenericInterfaceMock[T interface{ byte | internal.Internal }, U any] struct {
	T          *testing.T
	GetTStub   func() T
	GetTCalled int32
	GetUStub   func() U
	GetUCalled int32
}

// Verify that *MyGenericInterfaceMock implements MyGenericInterface.
func _[T interface{ byte | internal.Internal }, U any]() {
	var _ MyGenericInterface[T, U] = &MyGenericInterfaceMock[T, U]{}
}

// GetT is a stub for the MyGenericInterface.GetT
// method that records the number of times it has been called.
func (m *MyGenericInterfaceMock[T, U]) GetT() T {
	atomic.AddInt32(&m.GetTCalled, 1)
	if m.GetTStub == nil {
		if m.T != nil {
			m.T.Error("GetTStub is nil")
		}
		panic("GetT unimplemented")
	}
	return m.GetTStub()
}

// GetU is a stub for the MyGenericInterface.GetU
// method that records the number of times it has been called.
func (m *MyGenericInterfaceMock[T, U]) GetU() U {
	atomic.AddInt32(&m.GetUCalled, 1)
	if m.GetUStub == nil {
		if m.T != nil {
			m.T.Error("GetUStub is nil")
		}
		panic("GetU unimplemented")
	}
	return m.GetUStub()
}
