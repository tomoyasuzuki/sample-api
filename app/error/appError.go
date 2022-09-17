package appError

type Code string

const (
	ConnectionFailed Code = "ConnectionFailed"
	MigrationFailed  Code = "MigrationFailed"
	RecordNotFound   Code = "RecordNotFound"
	Unknown          Code = "Unknown"
)

type AppError struct {
	error
	Code Code
}

func (a AppError) Error() string {
	var message string

	switch a.Code {
	case RecordNotFound:
		message = "record not found"
	}

	return message
}

func (a AppError) GetHttpCode() int {
	return 200
}

func New(code Code) AppError {
	return AppError{Code: code}
}
