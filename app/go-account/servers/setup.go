package servers

import (
	"context"

	"github.com/Calmantara/go-common/config"

	accounthdl "github.com/Calmantara/go-account/modules/handler/account"
	accountrepo "github.com/Calmantara/go-account/modules/repository/account"
	activityrepo "github.com/Calmantara/go-account/modules/repository/accountactivity"
	accountsvc "github.com/Calmantara/go-account/modules/service/account"
	c "github.com/Calmantara/go-common/pkg/context"
	"github.com/Calmantara/go-common/pkg/logger"
)

type handlers struct {
	accountHdl accounthdl.IAccountHandler
}

func initDI() handlers {
	ctx, _ := c.GetCorrelationID(context.Background())

	logger.Info(ctx, "setup repository")
	pgConn := config.NewPostgresGormConn()
	accountRepo := accountrepo.NewAccountRepoGormImpl(pgConn)
	activityRepo := activityrepo.NewActivityRepoGormImpl(pgConn)

	logger.Info(ctx, "setup service")
	accountSvc := accountsvc.NewAccountServiceImpl(accountRepo, activityRepo)

	logger.Info(ctx, "setup handler")
	accountHdl := accounthdl.NewAccountHandlerImpl(accountSvc)

	return handlers{
		accountHdl: accountHdl,
	}
}
