package gotry

// GetStackTraceList get list stack trace from exception
func (e *exceptionInteractor) GetStackTraceList() []string {
	if e.stackTrace == nil {
		return nil
	}

	return e.stackTrace.GetList()
}
