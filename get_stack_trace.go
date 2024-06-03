package gotry

// GetStackTrace get stack trace from exception
func (e *exceptionInteractor) GetStackTrace() string {
	if e.stackTrace == nil {
		return ""
	}

	return e.stackTrace.String()
}
