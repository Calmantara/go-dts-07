package user

import (
	"context"
	"errors"

	"github.com/Calmantara/go-dts-07/config"
	"github.com/Calmantara/go-dts-07/module/model"
)

type UserMapImpl struct {
	dataStore  config.DataStore
	cacheEmail map[string]bool
}

func NewUserMap(dataStore config.DataStore) UserRepo {
	cache := make(map[string]bool)
	return &UserMapImpl{
		dataStore:  dataStore,
		cacheEmail: cache,
	}
}

func (um *UserMapImpl) FindUserById(ctx context.Context, userId uint64) (user model.User, err error) {
	user, ok := um.dataStore.UserData[userId]
	if !ok || user.Delete {
		err = errors.New("user is not found")
	}
	return
}

func (um *UserMapImpl) FindAllUsers(ctx context.Context) (users []model.User, err error) {
	for _, usr := range um.dataStore.UserData {
		if !usr.Delete {
			users = append(users, usr)
		}
	}
	return
}

func (um *UserMapImpl) InsertUser(ctx context.Context, userIn model.User) (user model.User, err error) {
	if um.cacheEmail[userIn.Email] {
		err = errors.New("error duplication email")
		return
	}
	// insert
	userIn.Id = uint64(len(um.dataStore.UserData) + 1)
	um.dataStore.UserData[userIn.Id] = userIn // store user
	um.cacheEmail[userIn.Email] = true        // store email
	return userIn, err
}

func (um *UserMapImpl) UpdateUser(ctx context.Context, userIn model.User) (err error) {
	// update
	user, err := um.FindUserById(ctx, userIn.Id)
	if err != nil {
		return err
	}
	delete(um.cacheEmail, user.Email)         // delete previous email cache
	um.dataStore.UserData[userIn.Id] = userIn // update user map
	um.cacheEmail[userIn.Email] = true        // store email cache
	return err
}

func (um *UserMapImpl) DeleteUserById(ctx context.Context, userId uint64) (user model.User, err error) {

	return
}
