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
	Retries       int       `bun:"retries"`
	Password      string    `bun:"password"`
	Name          string    `bun:"name"`
}

func accountDomain2PO(account *domain.Account) *Account {
	return &Account{
		ID:       account.ID,
		Password: account.Password,
		Name:     account.Name,
	}
}

func accountPO2Domain(accountPO *Account) *domain.Account {
	account := &domain.Account{
		ID:        accountPO.ID,
		Retries:   accountPO.Retries,
		UpdatedAt: accountPO.UpdatedAt,
		Name:      accountPO.Name,
		Password:  accountPO.Password,
	}
	return account
}

func (o *Account) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		o.CreatedAt = time.Now()
		o.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		o.UpdatedAt = time.Now()
	}
	return nil
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

func (o *accountRepo) GetOneByName(ctx context.Context, tx *bun.Tx, name string) (*domain.Account, error) {
	accountPO := &Account{}
	stmt := tx.NewSelect().Model(accountPO).Where("name = ?", name)
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

func (o *accountRepo) UpdateRetries(ctx context.Context, tx *bun.Tx, id int, value int) error {
	_, err := tx.NewUpdate().Model((*Account)(nil)).
		Set("retries = ?", value).Set("updated_at = ?", time.Now()).
		Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
