package user

import (
	"context"

	model "github.com/Calmantara/go-dts-user/module/model/user"
)

type UserService interface {
	FindUserByIdSvc(ctx context.Context, userId uint64) (user model.GetUserResponse, err error)
	FindAllUsersSvc(ctx context.Context) (users []model.GetUserResponse, err error)
	InsertUserSvc(ctx context.Context, userIn model.CreateUser) (user model.CreateUserResponse, err error)
	UpdateUserSvc(ctx context.Context, userIn model.UpdateUser) (err error)
	DeleteUserByIdSvc(ctx context.Context, userId uint64) (user model.GetUserResponse, err error)
}
