package gotry

import (
	"errors"
	"reflect"
	"runtime"
)

// Catch running catch function if it catch exception (error)
func (e *exceptionInteractor) Catch(f any) (except Exception) {
	if e == nil {
		return nil
	}

	if e.exception != nil {
		catchFunc(e, f)
	}

	return e
}

func catchFunc(e *exceptionInteractor, f any) {
	defer recoverFunc(e)

	var (
		fVal  = reflect.ValueOf(f)
		fType = fVal.Type()
		fCall = make([]reflect.Value, 0)
	)

	if fType.NumIn() > 2 {
		panic("invalid parameter for catch")
	}

	switch exception := e.exception.(type) {
	case runtime.Error:
		e.exception = errors.New(exception.Error())
	}

	var (
		exceptionType  = reflect.TypeOf(e.exception)
		stType         StackTrace
		stackTraceType = reflect.TypeOf(&stType).Elem()
		err            error
		errorType      = reflect.TypeOf(&err).Elem()
	)

	for i := 0; i < fType.NumIn(); i++ {
		paramType := fType.In(i)
		if paramType.AssignableTo(errorType) {
			fCall = append(fCall, reflect.ValueOf(e.exception))
		} else if paramType.AssignableTo(exceptionType) {
			fCall = append(fCall, reflect.ValueOf(e.exception))
		} else if paramType.AssignableTo(stackTraceType) {
			fCall = append(fCall, reflect.ValueOf(e.stackTrace))
		} else {
			panic("invalid parameter for catch")
		}
	}

	fVal.Call(fCall)
}
