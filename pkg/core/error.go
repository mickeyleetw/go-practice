package core

import (
	"fmt"
	"net/http"
)

type DomainError interface {
	Error() string
	HttpStatus() int
	Msg() string
}

type BaseError struct {
	httpStatus int
	msg        string
}

func (o BaseError) Error() string { return fmt.Sprintf("msg: %s", o.msg) }

func (o BaseError) HttpStatus() int { return o.httpStatus }

func (o BaseError) Msg() string { return o.msg }

type ValidationError struct {
	BaseError
}

func NewValidationError(msg string) ValidationError {
	return ValidationError{BaseError{
		httpStatus: http.StatusUnprocessableEntity, msg: fmt.Sprintf("[ValidationError]%s", msg),
	}}
}

type UnexpectedError struct {
	BaseError
}

func NewUnexpectedError(msg string) UnexpectedError {
	return UnexpectedError{BaseError{
		httpStatus: http.StatusInternalServerError, msg: fmt.Sprintf("[UnexpectedError]%s", msg),
	}}
}

type ResourceNotFoundError struct {
	BaseError
}

func NewResourceNotFoundError(msg string) ResourceNotFoundError {
	return ResourceNotFoundError{BaseError{
		httpStatus: http.StatusNotFound, msg: fmt.Sprintf("[ResourceNotFoundError]%s", msg),
	}}
}

type PasswordNotMatchError struct {
	BaseError
}

func NewPasswordNotMatchError(msg string) PasswordNotMatchError {
	return PasswordNotMatchError{BaseError{
		httpStatus: http.StatusForbidden, msg: fmt.Sprintf("[PasswordNotMatchError]%s", msg),
	}}
}

type AccountLockedError struct {
	BaseError
}

func NewAccountLockedError(msg string) AccountLockedError {
	return AccountLockedError{BaseError{
		httpStatus: http.StatusForbidden, msg: fmt.Sprintf("[AccountLockedError]%s", msg),
	}}
}
