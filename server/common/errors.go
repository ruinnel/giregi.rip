package common

import (
	"fmt"
	"net/http"
)

type ErrorCode string

const (
	ErrorUnauthorized ErrorCode = "UNAUTHORIZED"
	ErrorForbidden    ErrorCode = "FORBIDDEN"
	ErrorInvalidParam ErrorCode = "INVALID_PARAM"
	ErrorDatabase     ErrorCode = "DATABASE_ERROR"
	ErrorFirebase     ErrorCode = "FIREBASE_ERROR"
	ErrorArchiveFail  ErrorCode = "ARCHIVE_FAIL"
	ErrorNotFound     ErrorCode = "NOT_FOUND"
	ErrorUnknown      ErrorCode = "UNKNOWN"
)

type StatusError struct {
	Status  int
	Code    ErrorCode
	Message string
	Cause   []error
}

func (e *StatusError) Error() string {
	if len(e.Cause) > 0 {
		return fmt.Sprintf("(%d, %s) %s - cause: [%v]", e.Status, e.Code, e.Message, e.Cause)
	} else {
		return fmt.Sprintf("(%d, %s) %s", e.Status, e.Code, e.Message)
	}
}

func NewUnauthorizedError(message string, cause ...error) error {
	return &StatusError{
		Status:  http.StatusUnauthorized,
		Code:    ErrorUnauthorized,
		Message: message,
		Cause:   cause,
	}
}

func NewForbiddenError(message string, cause ...error) error {
	return &StatusError{
		Status:  http.StatusForbidden,
		Code:    ErrorForbidden,
		Message: message,
		Cause:   cause,
	}
}

func NewInvalidParamError(message string, cause ...error) error {
	return &StatusError{
		Status:  http.StatusBadRequest,
		Code:    ErrorInvalidParam,
		Message: message,
		Cause:   cause,
	}
}

func NewDatabaseError(message string, cause ...error) error {
	return &StatusError{
		Status:  http.StatusInternalServerError,
		Code:    ErrorDatabase,
		Message: message,
		Cause:   cause,
	}
}

func NewFirebaseError(message string, cause ...error) error {
	return &StatusError{
		Status:  http.StatusInternalServerError,
		Code:    ErrorFirebase,
		Message: message,
		Cause:   cause,
	}
}

func NewArchiveFailError(message string, cause ...error) error {
	return &StatusError{
		Status:  http.StatusInternalServerError,
		Code:    ErrorArchiveFail,
		Message: message,
		Cause:   cause,
	}
}

func NewNotFoundError(message string, cause ...error) error {
	return &StatusError{
		Status:  http.StatusBadRequest,
		Code:    ErrorNotFound,
		Message: message,
		Cause:   cause,
	}
}

func NewUnknownError(message string, cause ...error) error {
	return &StatusError{
		Status:  http.StatusInternalServerError,
		Code:    ErrorUnknown,
		Message: message,
		Cause:   cause,
	}
}
