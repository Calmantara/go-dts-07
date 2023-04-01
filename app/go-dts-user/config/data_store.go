package config

import model "github.com/Calmantara/go-dts-user/module/model/user"

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
