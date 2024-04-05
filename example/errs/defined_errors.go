package errs

import "github.com/aivyss/perr"

var (
	MyFirstError  = perr.NewPreparedError("my-first-error")
	MySecondError = perr.NewPreparedError("my-second-error")
)

func GetFirstError() error {
	return MyFirstError
}

func GetSecondError() error {
	return MySecondError
}

func GetFirstErrorWithMessage(message string) error {
	return MyFirstError.AddMessage("added message: %s", message)
}

func GetSecondErrorWithMessage(message string) error {
	return MySecondError.AddMessage("added message: %s", message)
}
