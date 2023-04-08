package account

import (
	"context"
	"errors"
	"testing"
	"time"

	accountmodel "github.com/Calmantara/go-account/modules/models/account"
	repomock "github.com/Calmantara/go-account/modules/repository/account/mock"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	createdAt := time.Now()
	id := uuid.New()

	type (
		input struct {
			acc accountmodel.CreateAccount
		}
		want struct {
			created accountmodel.AccountResponse
			err     error
		}
	)

	testCases := []struct {
		desc string
		// dia akan mocking seakan akan function yang kita test menjalankan procedure function kita
		doMock func(repoMock *repomock.MockIAccountRepo)
		input  input
		want   want
	}{
		{
			desc: "happy case",
			input: input{
				acc: accountmodel.CreateAccount{
					Username: "test",
					Password: "tong",
					Role:     accountmodel.ROLE_NORMAL,
				},
			},
			want: want{
				err: nil,
				created: accountmodel.AccountResponse{
					ID:        id,
					Username:  "test",
					Role:      accountmodel.ROLE_NORMAL,
					CreatedAt: createdAt,
				},
			},
			doMock: func(repoMock *repomock.MockIAccountRepo) {
				repoMock.
					EXPECT().
					CreateAccount(gomock.Any(), gomock.Any()).
					Return(accountmodel.Account{
						ID:        id,
						Username:  "test",
						Role:      accountmodel.ROLE_NORMAL,
						CreatedAt: createdAt,
					}, nil).
					MaxTimes(1).
					AnyTimes()
			},
		},
		{
			desc: "error repo create account",
			input: input{
				acc: accountmodel.CreateAccount{
					Username: "test",
					Password: "tong",
					Role:     accountmodel.ROLE_NORMAL,
				},
			},
			want: want{
				err:     errors.New("some error"),
				created: accountmodel.AccountResponse{},
			},
			doMock: func(repoMock *repomock.MockIAccountRepo) {
				repoMock.
					EXPECT().
					CreateAccount(gomock.Any(), gomock.Any()).
					Return(accountmodel.Account{}, errors.New("some error")).
					MaxTimes(1).
					AnyTimes()
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoMock := repomock.NewMockIAccountRepo(ctrl)
			tC.doMock(repoMock)

			svc := AccountServiceImpl{
				accountRepo: repoMock,
			}
			created, err := svc.CreateAccount(context.Background(), tC.input.acc)
			if tC.want.err != nil {
				assert.EqualError(t, err, tC.want.err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tC.want.created, created)
		})
	}
}
