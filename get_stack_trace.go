package gotry

func (e *exceptionInteractor) GetStackTrace() string {
	if e.stackTrace == nil {
		return ""
	}

	return e.stackTrace.String()
}
