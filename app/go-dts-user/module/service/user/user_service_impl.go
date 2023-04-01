package user

import (
	"context"

	"github.com/Calmantara/go-common/pkg/json"
	"github.com/Calmantara/go-common/pkg/logger"
	model "github.com/Calmantara/go-dts-user/module/model/user"
	"github.com/Calmantara/go-dts-user/module/repository/user"
)

type UserSvcImpl struct {
	userRepo user.UserRepo
}

func NewUserSvc(userRepo user.UserRepo) UserService {
	return &UserSvcImpl{
		userRepo: userRepo,
	}
}

func (u *UserSvcImpl) FindUserByIdSvc(ctx context.Context, userId uint64) (user model.GetUserResponse, err error) {
	logger.Info(ctx, "FindUserById invoked", "logCtx", u)

	var usermdl model.User
	if usermdl, err = u.userRepo.FindUserById(ctx, userId); err != nil {
		logger.Error(ctx, "error FindUserById", "logCtx", u, "err", err)
		return
	}

	if err = json.ObjectMapper(usermdl, user); err != nil {
		logger.Error(ctx, "error mapping user object", "logCtx", *u, "err", err)
		return
	}
	return
}

func (u *UserSvcImpl) FindAllUsersSvc(ctx context.Context) (users []model.GetUserResponse, err error) {
	logger.Info(ctx, "FindAllUsers invoked", "logCtx", *u)

	var usersmdl []model.User
	if usersmdl, err = u.userRepo.FindAllUsers(ctx); err != nil {
		logger.Error(ctx, "error FindAllUsers", "logCtx", *u, "err", err)
		return
	}

	if err = json.ObjectMapper(usersmdl, users); err != nil {
		logger.Error(ctx, "error mapping user object", "logCtx", *u, "err", err)
		return
	}
	return
}

func (u *UserSvcImpl) InsertUserSvc(ctx context.Context, userIn model.CreateUser) (user model.CreateUserResponse, err error) {
	logger.Info(ctx, "InsertUser invoked", "logCtx", *u)

	var usermdl model.User
	if err = json.ObjectMapper(userIn, usermdl); err != nil {
		logger.Error(ctx, "error mapping user object", "logCtx", *u, "err", err)
		return
	}

	if usermdl, err = u.userRepo.InsertUser(ctx, usermdl); err != nil {
		logger.Error(ctx, "error InsertUser", "logCtx", *u, "err", err)
		return
	}

	if err = json.ObjectMapper(usermdl, user); err != nil {
		logger.Error(ctx, "error mapping user object", "logCtx", *u, "err", err)
		return
	}
	return
}

func (u *UserSvcImpl) UpdateUserSvc(ctx context.Context, userIn model.UpdateUser) (err error) {
	logger.Info(ctx, "UpdateUser invoked", "logCtx", *u)

	var usermdl model.User
	if err = json.ObjectMapper(userIn, usermdl); err != nil {
		logger.Error(ctx, "error mapping user object", "logCtx", *u, "err", err)
		return
	}

	if err = u.userRepo.UpdateUser(ctx, usermdl); err != nil {
		logger.Error(ctx, "error UpdateUser", "logCtx", *u, "err", err)
		return
	}
	return
}

func (u *UserSvcImpl) DeleteUserByIdSvc(ctx context.Context, userId uint64) (user model.GetUserResponse, err error) {
	logger.Info(ctx, "DeleteUserById invoked", "logCtx", *u)

	var deletedUser model.User
	if deletedUser, err = u.userRepo.DeleteUserById(ctx, userId); err != nil {
		logger.Error(ctx, "error DeleteUserById", "logCtx", *u, "err", err)
		return
	}

	if err = json.ObjectMapper(deletedUser, user); err != nil {
		logger.Error(ctx, "error mapping user object", "logCtx", *u, "err", err)
		return
	}

	return
}
