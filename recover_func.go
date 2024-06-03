package gotry

// recoverFunc check exception and recover panic error
func recoverFunc(e *exceptionInteractor) {
	if err := recover(); err != nil {
		e.exception = err
		e.stackTrace = getStackTrace()
	}
}
