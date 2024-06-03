package gotry

// GetException return exception with custom type
func (e *exceptionInteractor) GetException() any {
	return e.exception
}
