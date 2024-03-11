package gotry

func recoverFunc(e *exceptionInteractor) {
	if err := recover(); err != nil {
		e.exception = err
		e.stackTrace = getStackTrace()
	}
}
