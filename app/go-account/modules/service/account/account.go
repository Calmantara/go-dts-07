package account

import (
	"context"

	accountmodel "github.com/Calmantara/go-account/modules/models/account"
	token "github.com/Calmantara/go-account/modules/models/token"
)

type IAccountService interface {
	CreateAccount(ctx context.Context, acc accountmodel.CreateAccount) (created accountmodel.AccountResponse, err error)
	LoginAccountByUserName(ctx context.Context, loginAcc accountmodel.LoginAccount) (tokens token.Tokens, err error)
	GetAccount(ctx context.Context, userId string) (account accountmodel.AccountResponse, err error)
}
