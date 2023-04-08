package account

import (
	"context"
	"errors"
	"regexp"
	"testing"

	accountmodel "github.com/Calmantara/go-account/modules/models/account"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetAccountByUserID(t *testing.T) {
	id := uuid.New()
	query := `SELECT * 
	FROM "accounts" 
	WHERE id = $1 
		AND "accounts"."deleted_at" IS NULL`

	type (
		input struct {
			userId string
		}
		want struct {
			err     error
			account accountmodel.Account
		}
	)

	testCases := []struct {
		desc   string
		input  input
		want   want
		doMock func(mock sqlmock.Sqlmock)
	}{
		{
			desc: "happy case",
			input: input{
				userId: id.String(),
			},
			want: want{
				err: nil,
				account: accountmodel.Account{
					ID:       id,
					Username: "this-is-username",
					Password: "password",
					Role:     accountmodel.ROLE_NORMAL,
				},
			},
			doMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows(
					[]string{"id", "username", "password", "role"}).
					AddRow(id, "this-is-username", "password", "normal")

				mock.
					ExpectQuery(
						regexp.QuoteMeta(query),
					).
					WithArgs(id).
					WillReturnError(nil).
					WillReturnRows(rows)
			},
		},
		{
			desc: "error case",
			input: input{
				userId: id.String(),
			},
			want: want{
				err:     errors.New("some error"),
				account: accountmodel.Account{},
			},
			doMock: func(mock sqlmock.Sqlmock) {
				mock.
					ExpectQuery(
						regexp.QuoteMeta(query),
					).
					WithArgs(id).
					WillReturnError(errors.New("some error"))
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			DB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: db,
			}), &gorm.Config{})
			tC.doMock(mock)

			repo := AccountRepoGormImpl{
				master: DB,
			}

			acc, err := repo.GetAccountByUserID(context.Background(), tC.input.userId)
			if tC.want.err != nil {
				assert.EqualError(t, err, tC.want.err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tC.want.account, acc)
		})
	}
}
