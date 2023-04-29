package core

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type DomainError interface {
	Error() string
	HttpStatus() int
	Code() int
	Msg() string
}

type BaseError struct {
	httpStatus int
	code       int
	msg        string
}

func (o BaseError) Error() string { return fmt.Sprintf("msg: %s", o.msg) }

func (o BaseError) HttpStatus() int { return o.httpStatus }

func (o BaseError) Code() int { return o.code }

func (o BaseError) Msg() string { return o.msg }

type ValidationError struct {
	BaseError
}

func NewValidationError(msg string) ValidationError {
	return ValidationError{BaseError{
		httpStatus: http.StatusUnprocessableEntity, code: 1002, msg: fmt.Sprintf("[ValidationError]%s", msg),
	}}
}

type UnexpectedError struct {
	BaseError
}

func NewUnexpectedError(msg string) UnexpectedError {
	return UnexpectedError{BaseError{
		httpStatus: http.StatusInternalServerError, code: 1001, msg: fmt.Sprintf("[UnexpectedError]%s", msg),
	}}
}

type ResourceNotFoundError struct {
	BaseError
}

func NewResourceNotFoundError(msg string) ResourceNotFoundError {
	return ResourceNotFoundError{BaseError{
		httpStatus: http.StatusNotFound, code: 1003, msg: fmt.Sprintf("[ResourceNotFoundError]%s", msg),
	}}
}

type errorResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func WriteErrorResp(c *gin.Context, err error) {
	domainErr, ok := err.(DomainError)
	if ok {
		errResp := errorResp{Code: domainErr.Code(), Message: domainErr.Msg()}
		c.JSON(domainErr.HttpStatus(), errResp)
	} else {
		unexpectedErr := NewUnexpectedError(err.Error())
		errResp := errorResp{Code: unexpectedErr.Code(), Message: unexpectedErr.Msg()}
		log.Warn().Msgf("Internal Error: %+v", errResp)
		c.JSON(unexpectedErr.HttpStatus(), errResp)
	}
}
