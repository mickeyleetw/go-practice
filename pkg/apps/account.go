package apps

import (
	"fmt"
	"net/http"
	errors "senao/pkg/core"
	"senao/pkg/domain"
	"senao/pkg/models"

	"github.com/gin-gonic/gin"
)

type accountHandler struct {
	accountUsecase domain.AccountUsecase
}

func NewAccountHandler(AccountUsecase domain.AccountUsecase) *accountHandler {
	return &accountHandler{accountUsecase: AccountUsecase}
}

func (o *accountHandler) CreateAccount(c *gin.Context) {

	// bind req
	req := &models.CreateAccountReq{}

	if err := c.ShouldBindJSON(req); err != nil {
		err2 := errors.NewValidationError(err.Error())
		WriteErrorResp(c, err2)
		return
	}

	// to req dto and create
	createAccountReqDTO := domain.CreateAccountReq2DTO(req)
	accountID, err := o.accountUsecase.CreateAccount(c, createAccountReqDTO)
	if err != nil {
		WriteErrorResp(c, err)
		return
	}

	if accountID != 0 {
		WriteSuccessResp(c, http.StatusCreated, fmt.Sprintf("account id: %d Created", accountID))
		return
	}
}

func (o *accountHandler) VerifyAccount(c *gin.Context) {

	// bind req
	req := &models.VerifyAccountReq{}

	if err := c.ShouldBindJSON(req); err != nil {
		err2 := errors.NewValidationError(err.Error())
		WriteErrorResp(c, err2)
		return
	}

	// to req dto and verify
	verifyAccountReqDTO := domain.VerifyAccountReq2DTO(req)
	err := o.accountUsecase.VerifyAccount(c, verifyAccountReqDTO)
	if err != nil {
		WriteErrorResp(c, err)
		return
	}

	WriteSuccessResp(c, http.StatusAccepted, "account successfully verified")
}
