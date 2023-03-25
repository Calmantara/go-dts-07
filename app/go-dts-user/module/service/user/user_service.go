package user

import (
	"context"

	"github.com/Calmantara/go-dts-user/module/model"
)

type UserService interface {
	FindUserByIdSvc(ctx context.Context, userId uint64) (user model.User, err error)
	FindAllUsersSvc(ctx context.Context) (users []model.User, err error)
	InsertUserSvc(ctx context.Context, userIn model.User) (user model.User, err error)
	UpdateUserSvc(ctx context.Context, userIn model.User) (err error)
	DeleteUserByIdSvc(ctx context.Context, userId uint64) (user model.User, err error)
}
