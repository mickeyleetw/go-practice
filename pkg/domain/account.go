package domain

import (
	"context"
	"senao/pkg/models"
	"time"

	"github.com/uptrace/bun"
)

type Account struct {
	ID        int
	UpdatedAt time.Time
	Retries   int
	Name      string
	Password  string
}

type CreateAccountReqDTO struct {
	Name     string `validate:"required,lte=32,gte=3"`
	Password string `validate:"required,lte=32,gte=8"`
}

func CreateAccountReq2DTO(req *models.CreateAccountReq) *CreateAccountReqDTO {
	createAccountReqDTO := &CreateAccountReqDTO{
		Name:     req.Name,
		Password: req.Password,
	}
	return createAccountReqDTO
}

type VerifyAccountReqDTO struct {
	Name     string `validate:"required"`
	Password string `validate:"required"`
}

func VerifyAccountReq2DTO(req *models.VerifyAccountReq) *VerifyAccountReqDTO {
	verifyAccountReqDTO := &VerifyAccountReqDTO{
		Name:     req.Name,
		Password: req.Password,
	}
	return verifyAccountReqDTO
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
	CreateAccount(ctx context.Context, CreateAccountReqDTO *CreateAccountReqDTO) (int, error)
	VerifyAccount(ctx context.Context, VerifyAccountReqDTO *VerifyAccountReqDTO) error
}

type AccountRepo interface {
	CreateOne(ctx context.Context, tx *bun.Tx, account *Account) (int, error)
	GetOneByName(ctx context.Context, tx *bun.Tx, name string) (*Account, error)
	UpdateRetries(ctx context.Context, tx *bun.Tx, id int, value int) error
}
