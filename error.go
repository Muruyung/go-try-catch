package gotry

import "fmt"

func (e *exceptionInteractor) Error() error {
	return fmt.Errorf("%v", e.exception)
}
