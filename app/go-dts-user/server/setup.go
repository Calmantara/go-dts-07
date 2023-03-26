package server

import (
	"context"

	"github.com/Calmantara/go-dts-user/config"
	userhdl "github.com/Calmantara/go-dts-user/module/handler/user"
	userrepo "github.com/Calmantara/go-dts-user/module/repository/user"
	usersvc "github.com/Calmantara/go-dts-user/module/service/user"
	c "github.com/Calmantara/go-dts-user/pkg/context"
	"github.com/Calmantara/go-dts-user/pkg/logger"
)

type handlers struct {
	userHdl userhdl.UserHandler
}

func initDI() handlers {
	ctx, _ := c.GetCorrelationID(context.Background())

	logger.Info(ctx, "setup repository")
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

	logger.Info(ctx, "setup service")
	userSvc := usersvc.NewUserSvc(userRepo)

	logger.Info(ctx, "setup handler")
	userHdl := userhdl.NewUserHandler(userSvc)

	return handlers{
		userHdl: userHdl,
	}
}
