package user

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/Calmantara/go-common/pkg/json"
	"github.com/Calmantara/go-common/pkg/logger"
	jsonplaceholder "github.com/Calmantara/go-dts-user/client/json-placeholder"
	model "github.com/Calmantara/go-dts-user/module/model/user"
	"github.com/Calmantara/go-dts-user/module/repository/user"
)

type UserSvcImpl struct {
	userRepo           user.UserRepo
	jsonPlaceholderCln jsonplaceholder.JsonPlaceholderClient
}

func NewUserSvc(userRepo user.UserRepo, jsonPlaceholderCln jsonplaceholder.JsonPlaceholderClient) UserService {
	return &UserSvcImpl{
		userRepo:           userRepo,
		jsonPlaceholderCln: jsonPlaceholderCln,
	}
}

func (u *UserSvcImpl) FindUserByIdSvc(ctx context.Context, userId uint64) (user model.GetUserResponseWithTodo, err error) {
	logger.Info(ctx, "FindUserById invoked", "logCtx", u)
	tchan := u.getTodoAsync(ctx)

	// asumsi user akan selalu ada
	var usermdl model.User
	if usermdl, err = u.userRepo.FindUserById(ctx, userId); err != nil {
		logger.Error(ctx, "error FindUserById", "logCtx", u, "err", err)
		return
	}

	if err = json.ObjectMapper(usermdl, &user); err != nil {
		logger.Error(ctx, "error mapping user object", "logCtx", *u, "err", err)
		return
	}
	// get channel data
	todo := <-tchan
	if err = todo.err; err != nil {
		return
	}
	user.Todo = todo.todo
	return
}

type todoChan struct {
	todo jsonplaceholder.JsonPlaceholderResp
	err  error
}

func (u *UserSvcImpl) getTodoAsync(ctx context.Context) (resChan <-chan todoChan) {
	logCtx := fmt.Sprintf("%T", u)
	logger.Info(ctx, "getTodoAsync invoked", "logCtx", logCtx)

	rc := make(chan todoChan)

	go func() {
		// generate random number
		min := 10
		max := 30
		phMdl, err := u.jsonPlaceholderCln.GetTodo(ctx, rand.Intn(max-min+1)+min)
		if err != nil {
			logger.Error(ctx, "error fetching json placeholder", "err", err)
		}
		rc <- todoChan{
			todo: phMdl,
			err:  err,
		}
	}()
	return rc
}

func (u *UserSvcImpl) FindAllUsersSvc(ctx context.Context) (users []model.GetUserResponse, err error) {
	logger.Info(ctx, "FindAllUsers invoked", "logCtx", *u)

	var usersmdl []model.User
	if usersmdl, err = u.userRepo.FindAllUsers(ctx); err != nil {
		logger.Error(ctx, "error FindAllUsers", "logCtx", *u, "err", err)
		return
	}

	if err = json.ObjectMapper(usersmdl, &users); err != nil {
		logger.Error(ctx, "error mapping user object", "logCtx", *u, "err", err)
		return
	}
	return
}

func (u *UserSvcImpl) InsertUserSvc(ctx context.Context, userIn model.CreateUser) (user model.CreateUserResponse, err error) {
	logger.Info(ctx, "InsertUser invoked", "logCtx", *u)

	var usermdl model.User
	if err = json.ObjectMapper(userIn, &usermdl); err != nil {
		logger.Error(ctx, "error mapping user object", "logCtx", *u, "err", err)
		return
	}

	if usermdl, err = u.userRepo.InsertUser(ctx, usermdl); err != nil {
		logger.Error(ctx, "error InsertUser", "logCtx", *u, "err", err)
		return
	}

	if err = json.ObjectMapper(usermdl, &user); err != nil {
		logger.Error(ctx, "error mapping user object", "logCtx", *u, "err", err)
		return
	}
	return
}

func (u *UserSvcImpl) UpdateUserSvc(ctx context.Context, userIn model.UpdateUser) (err error) {
	logger.Info(ctx, "UpdateUser invoked", "logCtx", *u)

	var usermdl model.User
	if err = json.ObjectMapper(userIn, &usermdl); err != nil {
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

	if err = json.ObjectMapper(deletedUser, &user); err != nil {
		logger.Error(ctx, "error mapping user object", "logCtx", *u, "err", err)
		return
	}

	return
}
