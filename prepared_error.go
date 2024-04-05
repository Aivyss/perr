package perr

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
)

type persistError struct {
	original       error
	initialMessage string
}

func (e persistError) Error() string {
	return e.initialMessage
}

func (e persistError) Same(err error) bool {
	err = e.convertError(err)
	return errors.Is(e.original, err) || errors.Is(err, e.original)
}

func (e persistError) In(err ...error) bool {
	for _, ee := range err {
		if e.Same(ee) {
			return true
		}
	}

	return false
}

func (e persistError) convertError(err error) error {
	switch err.(type) {
	case PreparedError:
		return err.(PreparedError).original
	case persistError:
		return err.(persistError).original
	default:
		unwrappedErr := e.unwrapFully(err)

		switch unwrappedErr.(type) {
		case PreparedError, persistError:
			return e.convertError(unwrappedErr)
		default:
			return unwrappedErr
		}
	}
}

func (e persistError) unwrapFully(err error) error {
	var beforeErr = err

	for {
		e := errors.Unwrap(beforeErr)
		if e == nil {
			return beforeErr
		}

		beforeErr = e
	}
}

type PreparedError struct {
	persistError
	message     []string
	stackNumber int
	stackTrace  string
}

func (p PreparedError) Error() string {
	if len(p.message) == 0 {
		return p.initialMessage
	}

	return fmt.Sprintf("%s\n%v\n[stacks]\n%s", p.original.Error(), p.message, p.stackTrace)
}

func (p PreparedError) AddMessage(format string, args ...any) PreparedError {
	message := fmt.Sprintf(format, args...)

	var lines string
	for _, line := range strings.Split(string(debug.Stack()), "\n")[5:] {
		lines += fmt.Sprintf("\t%s\n", line)
	}

	modifiedPErr := PreparedError{
		persistError: p.persistError,
		message:      append(p.message, message),
		stackNumber:  p.stackNumber + 1,
		stackTrace:   fmt.Sprintf("%s\n", lines),
	}

	return modifiedPErr
}

type preparedErrorGroup []PreparedError

func (g preparedErrorGroup) Contains(err error) bool {
	for _, pErr := range g {
		if pErr.Same(err) {
			return true
		}
	}

	return false
}

func NewPreparedError(errorMessage string) PreparedError {
	original := errors.New(errorMessage)

	return PreparedError{
		persistError: persistError{
			original:       original,
			initialMessage: original.Error(),
		},
	}
}

func ErrorGroup(err ...PreparedError) preparedErrorGroup {
	return err
}
