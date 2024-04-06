package perr

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
)

type PreparedError struct {
	immutableError
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
		immutableError: p.immutableError,
		message:        append(p.message, message),
		stackNumber:    p.stackNumber + 1,
		stackTrace:     fmt.Sprintf("%s\n", lines),
	}

	return modifiedPErr
}

func NewPreparedError(errorMessage string) PreparedError {
	original := errors.New(errorMessage)

	return PreparedError{
		immutableError: immutableError{
			original:       original,
			initialMessage: original.Error(),
		},
	}
}
