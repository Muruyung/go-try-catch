package gotry

type exceptionInteractor struct {
	exception  any
	stackTrace StackTrace
}

// Try running try function
func Try(f func(ThrowNewException func(any))) Exception {
	e := new(exceptionInteractor)

	tryFunc(e, f)

	return e
}

// tryFunc running try function and check exception
func tryFunc(e *exceptionInteractor, try func(ThrowNewException func(any))) {
	defer recoverFunc(e)
	try(throwNewException)
}
