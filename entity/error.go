package entity

import (
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	shortenerv1 "github.com/indrasaputra/url-shortener/proto/indrasaputra/shortener/v1"
)

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

// ErrInternal returns codes.Internal.
// It explains that unexpected behavior occurred in system.
func ErrInternal(message string) error {
	st := status.New(codes.Internal, message)
	se := &shortenerv1.URLShortenerError{
		ErrorCode: shortenerv1.URLShortenerErrorCode_INTERNAL,
	}
	res, err := st.WithDetails(se)
	if err != nil {
		return st.Err()
	}
	return res.Err()
}

// ErrNilURL returns codes.InvalidArgument.
// It explains that the instance is empty or nil.
func ErrNilURL() error {
	st := status.New(codes.InvalidArgument, "")
	br := createBadRequest(&errdetails.BadRequest_FieldViolation{
		Field:       "URL instance",
		Description: "empty or nil",
	})

	se := &shortenerv1.URLShortenerError{
		ErrorCode: shortenerv1.URLShortenerErrorCode_EMPTY_URL,
	}
	res, err := st.WithDetails(br, se)
	if err != nil {
		return st.Err()
	}
	return res.Err()
}

// ErrAlreadyExists returns codes.AlreadyExists.
// It explains that the code / short URL already exists.
func ErrAlreadyExists() error {
	st := status.New(codes.AlreadyExists, "")
	se := &shortenerv1.URLShortenerError{
		ErrorCode: shortenerv1.URLShortenerErrorCode_ALREADY_EXISTS,
	}
	res, err := st.WithDetails(se)
	if err != nil {
		return st.Err()
	}
	return res.Err()
}

// ErrNotFound returns codes.NotFound.
// It explains that short URL is not found.
func ErrNotFound(message string) error {
	st := status.New(codes.NotFound, message)
	se := &shortenerv1.URLShortenerError{
		ErrorCode: shortenerv1.URLShortenerErrorCode_NOT_FOUND,
	}
	res, err := st.WithDetails(se)
	if err != nil {
		return st.Err()
	}
	return res.Err()
}

func createBadRequest(details ...*errdetails.BadRequest_FieldViolation) *errdetails.BadRequest {
	return &errdetails.BadRequest{
		FieldViolations: details,
	}
}
