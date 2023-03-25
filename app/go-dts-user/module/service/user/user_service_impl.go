package user

import (
	"context"
	"log"

	"github.com/Calmantara/go-dts-user/module/model"
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

func (u *UserSvcImpl) FindUserByIdSvc(ctx context.Context, userId uint64) (user model.User, err error) {
	log.Printf("[INFO] %T FindUserById invoked\n", u)
	if user, err = u.userRepo.FindUserById(ctx, userId); err != nil {
		log.Printf("[ERROR] error FindUserById :%v\n", err)
	}
	return
}

func (u *UserSvcImpl) FindAllUsersSvc(ctx context.Context) (users []model.User, err error) {
	log.Printf("[INFO] %T FindAllUsers invoked\n", u)
	if users, err = u.userRepo.FindAllUsers(ctx); err != nil {
		log.Printf("[ERROR] error FindAllUsers :%v\n", err)
	}
	return
}

func (u *UserSvcImpl) InsertUserSvc(ctx context.Context, userIn model.User) (user model.User, err error) {
	log.Printf("[INFO] %T InsertUser invoked\n", u)
	if user, err = u.userRepo.InsertUser(ctx, userIn); err != nil {
		log.Printf("[ERROR] error InsertUser :%v\n", err)
	}
	return
}

func (u *UserSvcImpl) UpdateUserSvc(ctx context.Context, userIn model.User) (err error) {
	log.Printf("[INFO] %T UpdateUser invoked\n", u)
	if err = u.userRepo.UpdateUser(ctx, userIn); err != nil {
		log.Printf("[ERROR] error InsertUser :%v\n", err)
	}
	return
}

func (u *UserSvcImpl) DeleteUserByIdSvc(ctx context.Context, userId uint64) (deletedUser model.User, err error) {
	log.Printf("[INFO] %T DeleteUserById invoked\n", u)
	if deletedUser, err = u.userRepo.DeleteUserById(ctx, userId); err != nil {
		log.Printf("[ERROR] error DeleteUserById :%v\n", err)
	}
	return
}
