package user

import (
	"context"

	"github.com/Calmantara/go-dts-user/module/model"
)

type UserRepo interface {
	FindUserById(ctx context.Context, userId uint64) (user model.User, err error)
	FindAllUsers(ctx context.Context) (users []model.User, err error)
	InsertUser(ctx context.Context, userIn model.User) (user model.User, err error)
	UpdateUser(ctx context.Context, userIn model.User) (err error)
	DeleteUserById(ctx context.Context, userId uint64) (user model.User, err error)
}
