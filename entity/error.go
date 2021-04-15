package entity

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	aptxv1 "github.com/indrasaputra/aptx/proto/indrasaputra/aptx/v1"
)

// ErrInternal returns codes.Internal.
// It explains that unexpected behavior occurred in system.
func ErrInternal(message string) error {
	st := status.New(codes.Internal, message)
	se := &aptxv1.AptxError{
		ErrorCode: aptxv1.AptxErrorCode_INTERNAL,
	}
	res, err := st.WithDetails(se)
	if err != nil {
		return st.Err()
	}
	return res.Err()
}

// ErrEmptyURL returns codes.InvalidArgument.
// It explains that the instance is empty or nil.
func ErrEmptyURL() error {
	st := status.New(codes.InvalidArgument, "")
	br := createBadRequest(&errdetails.BadRequest_FieldViolation{
		Field:       "URL instance",
		Description: "empty or nil",
	})

	se := &aptxv1.AptxError{
		ErrorCode: aptxv1.AptxErrorCode_EMPTY_URL,
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
	se := &aptxv1.AptxError{
		ErrorCode: aptxv1.AptxErrorCode_ALREADY_EXISTS,
	}
	res, err := st.WithDetails(se)
	if err != nil {
		return st.Err()
	}
	return res.Err()
}

// ErrNotFound returns codes.NotFound.
// It explains that short URL is not found.
func ErrNotFound() error {
	st := status.New(codes.NotFound, "")
	se := &aptxv1.AptxError{
		ErrorCode: aptxv1.AptxErrorCode_NOT_FOUND,
	}
	res, err := st.WithDetails(se)
	if err != nil {
		return st.Err()
	}
	return res.Err()
}

// ErrURLTooLong returns codes.InvalidArgument.
// It explains that original URL's length exceeds 65535.
func ErrURLTooLong() error {
	st := status.New(codes.InvalidArgument, "")
	se := &aptxv1.AptxError{
		ErrorCode: aptxv1.AptxErrorCode_URL_TOO_LONG,
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
