package perr

import "errors"

type immutableError struct {
	original       error
	initialMessage string
}

func (e immutableError) Error() string {
	return e.initialMessage
}

func (e immutableError) Same(err error) bool {
	err = e.convertError(err)
	return errors.Is(e.original, err) || errors.Is(err, e.original)
}

func (e immutableError) In(err ...error) bool {
	for _, ee := range err {
		if e.Same(ee) {
			return true
		}
	}

	return false
}

func (e immutableError) convertError(err error) error {
	switch err.(type) {
	case PreparedError:
		return err.(PreparedError).original
	case immutableError:
		return err.(immutableError).original
	default:
		unwrappedErr := e.unwrapFully(err)

		switch unwrappedErr.(type) {
		case PreparedError, immutableError:
			return e.convertError(unwrappedErr)
		default:
			return unwrappedErr
		}
	}
}

func (e immutableError) unwrapFully(err error) error {
	var beforeErr = err

	for {
		e := errors.Unwrap(beforeErr)
		if e == nil {
			return beforeErr
		}

		beforeErr = e
	}
}
