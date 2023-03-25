package server

import (
	"github.com/Calmantara/go-dts-07/config"
	userhdl "github.com/Calmantara/go-dts-07/module/handler/user"
	userrepo "github.com/Calmantara/go-dts-07/module/repository/user"
	usersvc "github.com/Calmantara/go-dts-07/module/service/user"
)

type handlers struct {
	userHdl userhdl.UserHandler
}

func initDI() handlers {
	dataStore := config.ConnectDataStore()
	userRepo := userrepo.NewUserMap(dataStore)
	if config.DATABASE == "PG" {
		pgConn := config.NewPostgresConn()
		userRepo = userrepo.NewUserPgRepo(pgConn)
	}

	userSvc := usersvc.NewUserSvc(userRepo)
	userHdl := userhdl.NewUserHandler(userSvc)

	return handlers{
		userHdl: userHdl,
	}
}
