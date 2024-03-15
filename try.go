package gotry

type exceptionInteractor struct {
	exception  any
	stackTrace StackTrace
}

func Try(f func(ThrowNewException func(any))) Exception {
	e := new(exceptionInteractor)

	tryFunc(e, f)

	return e
}

func tryFunc(e *exceptionInteractor, try func(ThrowNewException func(any))) {
	defer recoverFunc(e)
	try(throwNewException)
}
