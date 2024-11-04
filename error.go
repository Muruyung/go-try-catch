package gotry

import "fmt"

// Error get error from exception
func (e *exceptionInteractor) Error() error {
	if e.exception == nil {
		return nil
	}

	switch err := e.exception.(type) {
	case error:
		return err
	default:
		return fmt.Errorf("%v", e.exception)
	}
}
