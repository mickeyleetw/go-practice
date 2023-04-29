package apps

import (
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
		errors.WriteErrorResp(c, err2)
		return
	}

	// to req dto and create
	createAccountReqDTO := domain.CreateAccountReq2DTO(req)
	account, err := o.accountUsecase.CreateAccount(c, createAccountReqDTO)
	if err != nil {
		errors.WriteErrorResp(c, err)
		return
	}

	// to resp dto
	resp := domain.AccountDomain2Resp(account)
	c.JSON(http.StatusCreated, resp)
}
