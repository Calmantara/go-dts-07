package server

import (
	"github.com/Calmantara/go-dts-user/config"
	userhdl "github.com/Calmantara/go-dts-user/module/handler/user"
	userrepo "github.com/Calmantara/go-dts-user/module/repository/user"
	usersvc "github.com/Calmantara/go-dts-user/module/service/user"
)

type handlers struct {
	userHdl userhdl.UserHandler
}

func initDI() handlers {
	dataStore := config.ConnectDataStore()
	userRepo := userrepo.NewUserMap(dataStore)

	switch config.Load.DataSource.Mode {
	case config.MODE_GORM:
		pgConn := config.NewPostgresConn()
		userRepo = userrepo.NewUserPgRepo(pgConn)
	case config.MODE_PG:
		pgConn := config.NewPostgresConn()
		userRepo = userrepo.NewUserPgRepo(pgConn)
	}

	userSvc := usersvc.NewUserSvc(userRepo)
	userHdl := userhdl.NewUserHandler(userSvc)

	return handlers{
		userHdl: userHdl,
	}
}
