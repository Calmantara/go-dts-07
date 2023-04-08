package account

import (
	"context"
	"fmt"

	accountmodel "github.com/Calmantara/go-account/modules/models/account"
	"github.com/Calmantara/go-common/pkg/logger"
	"gorm.io/gorm"
)

type AccountRepoGormImpl struct {
	master *gorm.DB
}

func NewAccountRepoGormImpl(master *gorm.DB) IAccountRepo {
	return &AccountRepoGormImpl{
		master: master,
	}
}

func (a *AccountRepoGormImpl) CreateAccount(ctx context.Context, acc accountmodel.Account) (created accountmodel.Account, err error) {
	logCtx := fmt.Sprintf("%T - CreateAccount", a)
	logger.Info(ctx, "%v invoked", "logCtx", logCtx)

	err = a.master.
		Table("accounts").
		Create(&acc).Error
	if err != nil {
		return
	}

	return acc, err
}

func (a *AccountRepoGormImpl) GetAccountByUserName(ctx context.Context, username string) (account accountmodel.Account, err error) {
	logCtx := fmt.Sprintf("%T - GetAccountByUserName", a)
	logger.Info(ctx, "%v invoked", "logCtx", logCtx)

	err = a.master.
		Table("accounts").
		Where("username = ?", username).
		Find(&account).Error
	if err != nil {
		return
	}

	return account, err
}

func (a *AccountRepoGormImpl) GetAccountByUserID(ctx context.Context, userId string) (account accountmodel.Account, err error) {
	logCtx := fmt.Sprintf("%T - GetAccountByUserID", a)
	logger.Info(ctx, "%v invoked", "logCtx", logCtx)

	err = a.master.
		Table("accounts").
		Where("id = ?", userId).
		Find(&account).Error

	if err != nil {
		return
	}
	return account, err
}
