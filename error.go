package gotry

import "fmt"

// Error get error from exception
func (e *exceptionInteractor) Error() error {
	if e.exception == nil {
		return nil
	}

	return fmt.Errorf("%v", e.exception)
}
