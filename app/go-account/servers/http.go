package servers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Calmantara/go-account/modules/router/v1/account"
	"github.com/Calmantara/go-common/config"
	commonmidware "github.com/Calmantara/go-common/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func NewHttpServer() (srv *http.Server) {
	hdls := initDI()

	// init server
	ginServer := gin.Default()
	if config.Load.Server.Env == config.ENV_PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}

	// init middleware
	ginServer.Use(
		gin.Logger(),                             // untuk log request yang masuk
		gin.Recovery(),                           // untuk auto restart kalau panic
		commonmidware.CorrelationIDInterceptor(), // tracing purpose
	)

	// register router
	v1 := ginServer.Group("/api/v1")
	account.NewAccountRouter(v1, hdls.accountHdl)

	srv = &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Load.Server.Http.Port),
		Handler: ginServer,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	return
}
