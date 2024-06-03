package gotry

// Exception template for exception try catch
type Exception interface {
	Catch(any) Exception
	Error() error
	GetException() any
	GetStackTrace() string
	GetStackTraceList() []string
}

// StackTrace template for stacktrace try catch
type StackTrace interface {
	Print()
	String() string
	GetList() []string
}
