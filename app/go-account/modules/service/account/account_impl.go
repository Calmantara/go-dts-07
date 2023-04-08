package account

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Calmantara/go-common/pkg/logger"
	"github.com/google/uuid"

	accountmodel "github.com/Calmantara/go-account/modules/models/account"
	"github.com/Calmantara/go-account/modules/models/accountactivity"
	token "github.com/Calmantara/go-account/modules/models/token"
	accountrepo "github.com/Calmantara/go-account/modules/repository/account"
	activityrepo "github.com/Calmantara/go-account/modules/repository/accountactivity"
	crypto "github.com/Calmantara/go-account/pkg/crypto"
)

type AccountServiceImpl struct {
	accountRepo  accountrepo.IAccountRepo
	activityRepo activityrepo.IAccountActivityRepo
}

func NewAccountServiceImpl(
	accountRepo accountrepo.IAccountRepo,
	activityRepo activityrepo.IAccountActivityRepo,
) IAccountService {
	return &AccountServiceImpl{
		accountRepo:  accountRepo,
		activityRepo: activityRepo,
	}
}

func (a *AccountServiceImpl) CreateAccount(ctx context.Context, acc accountmodel.CreateAccount) (created accountmodel.AccountResponse, err error) {
	logCtx := fmt.Sprintf("%T - CreatedAccount", a)
	logger.Info(ctx, "invoked", "logCtx", logCtx)

	// need to hash password
	hashedPassowrd, err := crypto.GenerateHash(acc.Password)
	if err != nil {
		logger.Error(ctx, "error when hashing password",
			"logCtx", logCtx,
			"error", err)
		return
	}
	// update passowrd with hashed password
	acc.Password = hashedPassowrd
	// store to db
	createdAcc, err := a.accountRepo.CreateAccount(ctx, accountmodel.Account{
		ID:       uuid.New(),
		Username: acc.Username,
		Password: acc.Password,
		Role:     acc.Role,
	})
	if err != nil {
		logger.Error(ctx, "error when storing account",
			"logCtx", logCtx,
			"error", err)
		return
	}

	return accountmodel.AccountResponse{
		ID:        createdAcc.ID,
		Username:  createdAcc.Username,
		Role:      createdAcc.Role,
		CreatedAt: createdAcc.CreatedAt,
	}, err
}

func (a *AccountServiceImpl) LoginAccountByUserName(ctx context.Context, loginAcc accountmodel.LoginAccount) (tokens token.Tokens, err error) {
	logCtx := fmt.Sprintf("%T - LoginAccountByUserName", a)
	logger.Info(ctx, "invoked", "logCtx", logCtx)

	// get account by username
	acc, err := a.getAccountWithPassword(ctx, loginAcc.Username)
	if err != nil {
		logger.Error(ctx, "error when fetching account by username",
			"logCtx", logCtx,
			"error", err)
		return
	}

	// compare password
	// password acc -> hashed password
	// password login acc -> plain password
	if err = crypto.CompareHash(acc.Password, loginAcc.Password); err != nil {
		logger.Error(ctx, "error when comparing password",
			"logCtx", logCtx,
			"error", err)
		return
	}

	// record activity
	createdActivity, err := a.activityRepo.CreateActivity(ctx, accountactivity.AccountActivity{
		ID:     uuid.New(),
		UserID: acc.ID,
		Type:   accountactivity.ACTIVITY_LOGIN,
	})
	if err != nil {
		logger.Error(ctx, "error when creating activity",
			"logCtx", logCtx,
			"error", err)
		return
	}

	idToken, accessToken, refreshToken, err := a.generateAllTokensConcurrent(ctx,
		acc.ID.String(),
		acc.Username,
		string(acc.Role),
		createdActivity.ID.String())
	if err != nil {
		return
	}

	return token.Tokens{
		IDToken:      (idToken),
		AccessToken:  (accessToken),
		RefreshToken: (refreshToken),
	}, err
}

func (a *AccountServiceImpl) GetAccount(ctx context.Context, userId string) (account accountmodel.AccountResponse, err error) {
	logCtx := fmt.Sprintf("%T - GetAccount", a)
	logger.Info(ctx, "invoked", "logCtx", logCtx)
	// get account from database
	acc, err := a.accountRepo.GetAccountByUserID(ctx, userId)
	if err != nil {
		return
	}
	return accountmodel.AccountResponse{
		ID:        acc.ID,
		Username:  acc.Username,
		Role:      acc.Role,
		CreatedAt: acc.CreatedAt,
	}, err
}

func (a *AccountServiceImpl) getAccountWithPassword(ctx context.Context, username string) (account accountmodel.AccountResponseWithPassword, err error) {
	logCtx := fmt.Sprintf("%T - getAccountWithPassword", a)
	logger.Info(ctx, "invoked", "logCtx", logCtx)

	// get account from database
	acc, err := a.accountRepo.GetAccountByUserName(ctx, username)
	if err != nil {
		return
	}
	return accountmodel.AccountResponseWithPassword{
		AccountResponse: accountmodel.AccountResponse{
			ID:        acc.ID,
			Username:  acc.Username,
			Role:      acc.Role,
			CreatedAt: acc.CreatedAt,
		},
		Password: acc.Password,
	}, err
}

func (a *AccountServiceImpl) generateAllTokensConcurrent(ctx context.Context, userid, username, role, jti string) (idToken, accessToken, refreshToken string, err error) {
	logCtx := fmt.Sprintf("%T - generateAllTokens", a)
	logger.Info(ctx, "invoked", "logCtx", logCtx)

	// https://github.com/kataras/jwt
	timeNow := time.Now()
	defaultClaim := token.DefaultClaim{
		Expired:   int(timeNow.Add(24 * time.Hour).Unix()),
		NotBefore: int(timeNow.Unix()),
		IssuedAt:  int(timeNow.Unix()),
		Issuer:    "http://go-account",
		Audience:  "http://dts-07",
		JTI:       jti,
		Type:      token.ID_TOKEN,
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func(defaultClaim_ token.DefaultClaim) {
		defer wg.Done()
		// generate id token
		idTokenClaim := struct {
			token.DefaultClaim
			token.IDClaim
		}{
			DefaultClaim: defaultClaim_,
			IDClaim: token.IDClaim{
				Username: username,
				Role:     role,
			},
		}
		idToken, err = crypto.SignJWT(idTokenClaim)
		if err != nil {
			logger.Error(ctx, "error when creating id token",
				"logCtx", logCtx,
				"error", err)
			return
		}
	}(defaultClaim)

	go func(defaultClaim_ token.DefaultClaim) {
		defer wg.Done()
		// generate access token
		defaultClaim_.Expired = int(timeNow.Add(20 * time.Minute).UnixMilli())
		defaultClaim_.Type = token.ACCESS_TOKEN
		accessTokenClaim := struct {
			token.DefaultClaim
			token.AccessClaim
		}{
			DefaultClaim: defaultClaim_,
			AccessClaim: token.AccessClaim{
				Role:   role,
				UserID: userid,
			},
		}
		accessToken, err = crypto.SignJWT(accessTokenClaim)
		if err != nil {
			logger.Error(ctx, "error when creating access token",
				"logCtx", logCtx,
				"error", err)
			return
		}
	}(defaultClaim)

	go func(defaultClaim_ token.DefaultClaim) {
		defer wg.Done()
		// generate refresh token
		defaultClaim_.Expired = int(timeNow.Add(time.Hour).UnixMilli())
		defaultClaim_.Type = token.REFRESH_TOKEN
		refreshTokenClaim := struct {
			token.DefaultClaim
		}{
			DefaultClaim: defaultClaim_,
		}
		refreshToken, err = crypto.SignJWT(refreshTokenClaim)
		if err != nil {
			logger.Error(ctx, "error when creating refresh token",
				"logCtx", logCtx,
				"error", err)
			return
		}
	}(defaultClaim)

	wg.Wait()
	return
}

func (a *AccountServiceImpl) generateAllTokens(ctx context.Context, userid, username, role, jti string) (idToken, accessToken, refreshToken string, err error) {
	logCtx := fmt.Sprintf("%T - generateAllTokens", a)
	logger.Info(ctx, "invoked", "logCtx", logCtx)

	// https://github.com/kataras/jwt
	timeNow := time.Now()
	defaultClaim_ := token.DefaultClaim{
		Expired:   int(timeNow.Add(24 * time.Hour).Unix()),
		NotBefore: int(timeNow.Unix()),
		IssuedAt:  int(timeNow.Unix()),
		Issuer:    "http://go-account",
		Audience:  "http://dts-07",
		JTI:       jti,
		Type:      token.ID_TOKEN,
	}

	var wg sync.WaitGroup
	wg.Add(3)

	// generate id token
	idTokenClaim := struct {
		token.DefaultClaim
		token.IDClaim
	}{
		DefaultClaim: defaultClaim_,
		IDClaim: token.IDClaim{
			Username: username,
			Role:     role,
		},
	}
	idToken, err = crypto.SignJWT(idTokenClaim)
	if err != nil {
		logger.Error(ctx, "error when creating id token",
			"logCtx", logCtx,
			"error", err)
		return
	}

	// generate access token
	defaultClaim_.Expired = int(timeNow.Add(20 * time.Minute).UnixMilli())
	defaultClaim_.Type = token.ACCESS_TOKEN
	accessTokenClaim := struct {
		token.DefaultClaim
		token.AccessClaim
	}{
		DefaultClaim: defaultClaim_,
		AccessClaim: token.AccessClaim{
			Role:   role,
			UserID: userid,
		},
	}
	accessToken, err = crypto.SignJWT(accessTokenClaim)
	if err != nil {
		logger.Error(ctx, "error when creating access token",
			"logCtx", logCtx,
			"error", err)
		return
	}

	// generate refresh token
	defaultClaim_.Expired = int(timeNow.Add(time.Hour).UnixMilli())
	defaultClaim_.Type = token.REFRESH_TOKEN
	refreshTokenClaim := struct {
		token.DefaultClaim
	}{
		DefaultClaim: defaultClaim_,
	}
	refreshToken, err = crypto.SignJWT(refreshTokenClaim)
	if err != nil {
		logger.Error(ctx, "error when creating refresh token",
			"logCtx", logCtx,
			"error", err)
		return
	}

	return
}

// mockgen -source=modules/repository/account/account.go \ interface kita
// -destination=modules/repository/account/mock/account_mock.go \ mock kita mau diletakin mana
// -package=mock
