package server

import (
	"github.com/Calmantara/go-dts-07/config"
	"github.com/Calmantara/go-dts-07/module/router/v1/user"
	"github.com/gin-gonic/gin"
)

func NewHttpServer() {
	hdls := initDI()

	// init server
	ginServer := gin.Default()
	// init middleware
	ginServer.Use(
		gin.Logger(),   // untuk log request yang masuk
		gin.Recovery(), // untuk auto restart kalau panic
	)

	// register router
	v1 := ginServer.Group("/api/v1")
	user.NewUserRouter(v1, hdls.userHdl)

	ginServer.Run(config.PORT)
}
