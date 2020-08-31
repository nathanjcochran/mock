# Mock

`mock` is a code generation tool meant to be used with `go generate`. It
generates simple mock implementations of interfaces for use in testing.

Mocks are thread-safe.

## Installation

`go get -u github.com/nathanjcochran/mock`

## Usage

The only required argument to `mock` is the name of the interface to mock,
which must be provided after all other flags:

```
Usage: mock [options] interface
Options:
  -d string
    	Directory to search for interface in (default ".")
  -o string
    	Output file (default stdout)
```

## Example

Given this interface:

```go
package main

type Getter interface {
	GetByID(id int) ([]string, error)
	GetByName(name string) ([]string, error)
}
```

`mock Getter` will generate an implementation like this, and print it to
stdout:

```go
package main

// GetterMock is a mock implementation of the {{ .Name }}
// interface.
type GetterMock struct {
	GetByIDStub     func(id int) ([]string, error)
	GetByIDCalled   int32
	GetByNameStub   func(name string) ([]string, error)
	GetByNameCalled int32
}

var _ Getter = &GetterMock{}

// GetByID is a stub for the Getter.GetByID
// method that records the number of times it has been called.
func (m *GetterMock) GetByID(id int) ([]string, error) {
	atomic.AddInt32(m.GetByIDCalled, 1)
	return m.GetByIDStub(id)
}

// GetByName is a stub for the Getter.GetByName
// method that records the number of times it has been called.
func (m *GetterMock) GetByName(name string) ([]string, error) {
	atomic.AddInt32(m.GetByNameCalled, 1)
	return m.GetByNameStub(name)
}
```

## Go Generate

To use with `go generate`, simple place a `go:generate` comment somewhere in
your package (e.g. above the interface definition), like so:

`//go:generate mock -o getterMock.go Getter`

Note the use of the `-o` flag, which specifies the output file. If this flag
is not provided, the mocked implementation will be printed to stdout.

Then run the `go generate` command from the package directory.

Voila! There should now be a `getterMock.go` file containing your new mock, in
the same package as the interface definition. Subsequent runs of `go generate`
will overwrite the file, so be careful not to edit it!
