package domain

import (
	"context"
	"senao/pkg/models"
	"time"

	"github.com/uptrace/bun"
)

type Account struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Password  string
}

type CreateAccountReqDTO struct {
	Name     string `validate:"required,lte=32,gte=3"`
	Password string `validate:"required,lte=32,gte=8,contains=1Upper,contains=1Lower,contains=1Digit"`
}

func CreateAccountReq2DTO(req *models.CreateAccountReq) *CreateAccountReqDTO {
	createAccountReqDTO := &CreateAccountReqDTO{
		Name:     req.Name,
		Password: req.Password,
	}
	return createAccountReqDTO
}

func AccountDomain2Resp(account *Account) *models.AccountResp {
	accountResp := &models.AccountResp{
		ID:       account.ID,
		Name:     account.Name,
		Password: account.Password,
	}
	return accountResp
}

type AccountUsecase interface {
	CreateAccount(ctx context.Context, CreateAccountReqDTO *CreateAccountReqDTO) (*Account, error)
}

type AccountRepo interface {
	CreateOne(ctx context.Context, tx *bun.Tx, account *Account) (int, error)
	GetOneByID(ctx context.Context, tx *bun.Tx, id int) (*Account, error)
}
