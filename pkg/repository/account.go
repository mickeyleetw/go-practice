package repository

import (
	"context"
	"database/sql"
	"errors"
	"senao/pkg/core"
	"senao/pkg/domain"
	"time"

	"github.com/uptrace/bun"
)

type accountRepo struct {
}

func NewAccountRepo() *accountRepo {
	return &accountRepo{}
}

type Account struct {
	bun.BaseModel `bun:"table:account"`
	ID            int       `bun:"id,pk,nullzero"`
	CreatedAt     time.Time `bun:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at"`
	Password      string    `bun:"password"`
	Name          string    `bun:"name"`
}

func accountDomain2PO(account *domain.Account) *Account {
	return &Account{
		ID:        account.ID,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
		Password:  account.Password,
		Name:      account.Name,
	}
}

func accountPO2Domain(accountPO *Account) *domain.Account {
	account := &domain.Account{
		ID:       accountPO.ID,
		Name:     accountPO.Name,
		Password: accountPO.Password,
	}
	return account
}

func (o *accountRepo) CreateOne(ctx context.Context, tx *bun.Tx, account *domain.Account) (int, error) {
	// create account
	accountPO := accountDomain2PO(account)
	_, err := tx.NewInsert().Model(accountPO).Exec(ctx)
	if err != nil {
		return 0, err
	}

	return accountPO.ID, nil
}

func (o *accountRepo) GetOneByID(ctx context.Context, tx *bun.Tx, id int) (*domain.Account, error) {
	accountPO := &Account{}
	stmt := tx.NewSelect().Model(accountPO)
	err := stmt.Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NewResourceNotFoundError("Account")
		}
		return nil, err
	}

	account := accountPO2Domain(accountPO)
	return account, nil
}
