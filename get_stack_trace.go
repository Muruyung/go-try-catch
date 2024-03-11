package gotry

func (e *exceptionInteractor) GetStackTrace() string {
	return e.stackTrace.String()
}
