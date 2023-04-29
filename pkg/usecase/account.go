package usecase

import (
	"context"
	"database/sql"
	"senao/pkg/core"
	"senao/pkg/domain"
	"time"

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

func (o *accountUsecase) CreateAccount(ctx context.Context, createAccountReqDTO *domain.CreateAccountReqDTO) (int, error) {
	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// validate request
	if err := core.Validate.Struct(createAccountReqDTO); err != nil {
		return 0, core.NewValidationError(err.Error())
	}

	// create account
	account := createAccountReqDTO2Domain(createAccountReqDTO)
	accountID, err := o.accountRepo.CreateOne(ctx, &tx, account)
	if err != nil {
		return accountID, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return accountID, nil
}

func (o *accountUsecase) VerifyAccount(ctx context.Context, verifyAccountReqDTO *domain.VerifyAccountReqDTO) error {
	release_tag := false
	update_retries := 0
	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// validate request
	if err := core.Validate.Struct(verifyAccountReqDTO); err != nil {
		return core.NewValidationError(err.Error())
	}

	// get account by name
	existed_account, err := o.accountRepo.GetOneByName(ctx, &tx, verifyAccountReqDTO.Name)
	if err != nil {
		return err
	}

	// check account is locked
	if existed_account.Retries >= 5 {
		release_time := existed_account.UpdatedAt.Add(1 * time.Minute)
		if time.Now().Before(release_time) {
			return core.NewAccountLockedError("account is locked")
		} else {
			release_tag = true
		}
	}

	// verify account
	if verifyAccountReqDTO.Password == existed_account.Password {
		if release_tag {
			o.accountRepo.UpdateRetries(ctx, &tx, existed_account.ID, update_retries)
			if err := tx.Commit(); err != nil {
				return err
			}
		}
		return nil
	} else {
		if !release_tag {
			update_retries = existed_account.Retries + 1
		}
		o.accountRepo.UpdateRetries(ctx, &tx, existed_account.ID, update_retries+1)
		if err := tx.Commit(); err != nil {
			return err
		}
		return core.NewPasswordNotMatchError("password is wrong")
	}
}
