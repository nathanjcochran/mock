package example

//go:generate mocker MyInterface
type MyInterface interface {
	MyMethod1(param1 string, param2 int) (float64, error)
	MyMethod2() error
}
