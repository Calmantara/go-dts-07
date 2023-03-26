package server

import (
	"fmt"

	"github.com/Calmantara/go-dts-user/config"
	"github.com/Calmantara/go-dts-user/module/router/v1/user"
	"github.com/gin-gonic/gin"
)

func NewHttpServer() {
	hdls := initDI()

	// init server
	ginServer := gin.Default()

	if config.Load.Server.Env == config.ENV_PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}

	// init middleware
	ginServer.Use(
		gin.Logger(),   // untuk log request yang masuk
		gin.Recovery(), // untuk auto restart kalau panic
	)

	// register router
	v1 := ginServer.Group("/api/v1")
	user.NewUserRouter(v1, hdls.userHdl)

	ginServer.Run(fmt.Sprintf(":%v", config.Load.Server.Http.Port))
}
