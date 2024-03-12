package gotry

import "fmt"

func (e *exceptionInteractor) Error() error {
	if e.exception == nil {
		return nil
	}

	return fmt.Errorf("%v", e.exception)
}
