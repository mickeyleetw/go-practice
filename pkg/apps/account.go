package apps

import (
	"fmt"
	"net/http"
	errors "senao/pkg/core"

	// "senao/pkg/models"

	// "senao/pkg/models"

	"senao/pkg/domain"

	"github.com/gin-gonic/gin"
)

type accountHandler struct {
	accountUsecase domain.AccountUsecase
}

func NewAccountHandler(AccountUsecase domain.AccountUsecase) *accountHandler {
	return &accountHandler{accountUsecase: AccountUsecase}
}

type CreateAccountReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// type CreateAccountReqDTO struct {
// 	Name     string `validate:"required,lte=32,gte=3"`
// 	Password string `validate:"required,lte=32,gte=8,contains=1Upper,contains=1Lower,contains=1Digit"`
// }

type AccountResp struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func CreateAccountReq2DTO(req *CreateAccountReq) *domain.CreateAccountReqDTO {
	createAccountReqDTO := &domain.CreateAccountReqDTO{
		Name:     req.Name,
		Password: req.Password,
	}
	return createAccountReqDTO
}

func AccountDomain2Resp(account *domain.Account) *AccountResp {
	accountResp := &AccountResp{
		ID:       account.ID,
		Name:     account.Name,
		Password: account.Password,
	}
	return accountResp
}

func (o *accountHandler) CreateAccount(c *gin.Context) {

	// bind req
	fmt.Print(c.Request.Body)
	req := &CreateAccountReq{}

	if err := c.ShouldBindJSON(req); err != nil {
		err2 := errors.NewValidationError(err.Error())
		errors.WriteErrorResp(c, err2)
		return
	}

	// to req dto and create
	createAccountReqDTO := CreateAccountReq2DTO(req)
	account, err := o.accountUsecase.CreateAccount(c, createAccountReqDTO)
	if err != nil {
		errors.WriteErrorResp(c, err)
		return
	}

	// to resp dto
	resp := AccountDomain2Resp(account)
	c.JSON(http.StatusCreated, resp)
}
