package mock

import (
	"context"

	errors "senao/pkg/core"
	"senao/pkg/domain"

	"github.com/jinzhu/copier"
	"github.com/uptrace/bun"
)

type accountRepo struct {
	storage map[int]*domain.Account
}

func NewAccountRepo() *accountRepo {
	storage := make(map[int]*domain.Account)
	storage[999] = &domain.Account{ID: 999}
	return &accountRepo{storage: storage}
}

func (o *accountRepo) CreateOne(ctx context.Context, tx *bun.Tx, account *domain.Account) (int, error) {
	account2 := &domain.Account{}
	account.ID = 1
	copier.Copy(account2, account)
	o.storage[1] = account2
	return 1, nil
}

func (o *accountRepo) GetOneByID(ctx context.Context, tx *bun.Tx, id int, withEventTypes bool) (*domain.Account, error) {
	account, ok := o.storage[id]
	if !ok {
		return nil, errors.NewResourceNotFoundError("")
	}
	return account, nil
}

type AccountUsecase struct{}

func NewAccountUsecase() *AccountUsecase {
	return &AccountUsecase{}
}

func (o *AccountUsecase) CreateAccount(ctx context.Context, createAccountReqDTO *domain.CreateAccountReqDTO) (*domain.Account, error) {
	if createAccountReqDTO.Name == "" {
		return nil, errors.NewValidationError("")
	}
	return &domain.Account{
		ID:       1,
		Name:     createAccountReqDTO.Name,
		Password: createAccountReqDTO.Password,
	}, nil
}
