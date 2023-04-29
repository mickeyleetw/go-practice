package usecase

import (
	"context"
	"database/sql"
	"senao/pkg/core"
	"senao/pkg/domain"

	"github.com/uptrace/bun"
)

type accountUsecase struct {
	db          *bun.DB
	accountRepo domain.AccountRepo
}

func NewAccountUsecase(db *bun.DB, accountRepo domain.AccountRepo) *accountUsecase {
	return &accountUsecase{
		db:          db,
		accountRepo: accountRepo,
	}
}

func createAccountReqDTO2Domain(createAccountReqDTO *domain.CreateAccountReqDTO) *domain.Account {
	account := &domain.Account{
		Name:     createAccountReqDTO.Name,
		Password: createAccountReqDTO.Password,
	}
	return account
}

func (o *accountUsecase) CreateAccount(ctx context.Context, createAccountReqDTO *domain.CreateAccountReqDTO) (*domain.Account, error) {
	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// validate
	if err := core.Validate.Struct(createAccountReqDTO); err != nil {
		return nil, core.NewValidationError(err.Error())
	}

	// create account
	account := createAccountReqDTO2Domain(createAccountReqDTO)
	accountID, err := o.accountRepo.CreateOne(ctx, &tx, account)
	if err != nil {
		return nil, err
	}

	// get account
	account, err = o.accountRepo.GetOneByID(ctx, &tx, accountID)
	if err != nil {
		return nil, err
	}

	return account, nil
}
