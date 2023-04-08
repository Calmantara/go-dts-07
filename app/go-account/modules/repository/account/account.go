package account

import (
	"context"

	accountmodel "github.com/Calmantara/go-account/modules/models/account"
)

type IAccountRepo interface {
	CreateAccount(ctx context.Context, acc accountmodel.Account) (created accountmodel.Account, err error)
	GetAccountByUserName(ctx context.Context, username string) (account accountmodel.Account, err error)
	GetAccountByUserID(ctx context.Context, userId string) (account accountmodel.Account, err error)
}
