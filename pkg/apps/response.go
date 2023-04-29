package apps

import (
	"senao/pkg/core"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type errorResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func WriteErrorResp(c *gin.Context, err error) {
	domainErr, ok := err.(core.DomainError)
	if ok {
		errResp := errorResp{Success: false, Message: domainErr.Msg()}
		c.JSON(domainErr.HttpStatus(), errResp)
	} else {
		unexpectedErr := core.NewUnexpectedError(err.Error())
		errResp := errorResp{Success: false, Message: unexpectedErr.Msg()}
		log.Warn().Msgf("Internal Error: %+v", errResp)
		c.JSON(unexpectedErr.HttpStatus(), errResp)
	}
}

type successResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func WriteSuccessResp(sc *gin.Context, status int, msg string) {
	success := core.NewSuccess(status, msg)
	successResp := successResp{Success: true, Message: success.Msg()}
	sc.JSON(success.HttpStatus(), successResp)
}
