package config

import "github.com/Calmantara/go-dts-user/module/model"

type DataStore struct {
	UserData map[uint64]model.User
}

func ConnectDataStore() (ds DataStore) {
	// init map
	userData := make(map[uint64]model.User)

	return DataStore{
		UserData: userData,
	}
}
