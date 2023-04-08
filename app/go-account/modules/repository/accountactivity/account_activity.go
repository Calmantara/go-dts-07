package account

import (
	"context"

	activitymodel "github.com/Calmantara/go-account/modules/models/accountactivity"
)

type IAccountActivityRepo interface {
	CreateActivity(ctx context.Context, acc activitymodel.AccountActivity) (created activitymodel.AccountActivity, err error)
}
