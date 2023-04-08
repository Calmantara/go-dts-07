package account

import (
	"context"
	"fmt"

	activitymodel "github.com/Calmantara/go-account/modules/models/accountactivity"
	"github.com/Calmantara/go-common/pkg/logger"
	"gorm.io/gorm"
)

type ActivityRepoGormImpl struct {
	master *gorm.DB
}

func NewActivityRepoGormImpl(master *gorm.DB) IAccountActivityRepo {
	return &ActivityRepoGormImpl{
		master: master,
	}
}

func (a *ActivityRepoGormImpl) CreateActivity(ctx context.Context, acc activitymodel.AccountActivity) (created activitymodel.AccountActivity, err error) {
	logCtx := fmt.Sprintf("%T - CreateActivity", a)
	logger.Info(ctx, "%v invoked", "logCtx", logCtx)

	err = a.master.
		Table("account_activities").
		Create(&acc).Error
	if err != nil {
		return
	}

	return acc, err
}
