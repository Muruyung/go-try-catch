package gotry

type Exception interface {
	Catch(any) Exception
	Error() error
	GetException() any
	GetStackTrace() string
}

type StackTrace interface {
	Print()
	String() string
}
