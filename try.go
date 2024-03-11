package try

type exceptionInteractor struct {
	exception  any
	stackTrace StackTrace
}

func Try(f func()) Exception {
	e := new(exceptionInteractor)

	tryFunc(e, f)

	return e
}

func tryFunc(e *exceptionInteractor, try func()) {
	defer recoverFunc(e)
	try()
}
