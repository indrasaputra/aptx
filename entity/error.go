package entity

import "fmt"

var (
	// ErrInternalServer indicates there is unexpected problem occurs in the system itself.
	// The detail of the error/problem should be known in internal message.
	ErrInternalServer = NewError("01-001", "Internal server error")

	// ErrEmptyURL is returned when the original URL is empty.
	ErrEmptyURL = NewError("02-001", "URL is empty")
	// ErrDuplicatedShortURL is returned when the generated short URL already exists in the system.
	ErrDuplicatedShortURL = NewError("02-002", "Short URL already exists")
	// ErrURLNotFound is returned when the URL can't be found in the system.
	ErrURLNotFound = NewError("02-003", "URL not found")
)

// Error represents a data structure for error.
type Error struct {
	// Code represents error code.
	Code string `json:"code"`
	// Message represents error message.
	// This is the message that exposed to the user.
	Message string `json:"message"`
	// internalMessage represents deep error message.
	// This is should not be exposed to the user directly.
	// This attributes should be used as log.
	internalMessage string
}

// NewError creates an instance of Error.
func NewError(code, message string) *Error {
	return &Error{
		Code:            code,
		Message:         message,
		internalMessage: message,
	}
}

// Error returns internal message in one string.
func (err *Error) Error() string {
	return err.internalMessage
}

// WrapError wraps Error with given message.
// The message will be put in internalMessage attribute
// and can be accessed via Error() method.
func WrapError(err *Error, message string) *Error {
	return &Error{
		Code:            err.Code,
		Message:         err.Message,
		internalMessage: fmt.Sprintf("%s. %s", err.internalMessage, message),
	}
}
