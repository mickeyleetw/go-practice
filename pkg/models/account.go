package models

import (
	"senao/pkg/core"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("ValidateAccountPassword", core.ValidateAccountPassword)
	}

}

type CreateAccountReq struct {
	Name     string `json:"name" binding:"required,lte=32,gte=3"`
	Password string `json:"password" binding:"required,ValidateAccountPassword"`
}

type VerifyAccountReq struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AccountResp struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
