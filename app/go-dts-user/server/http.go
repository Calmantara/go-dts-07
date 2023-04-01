package server

import (
	"fmt"

	"github.com/Calmantara/go-common/config"
	"github.com/Calmantara/go-dts-user/module/router/v1/user"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GO DTS USER API DUCUMENTATION
// @version 2.0
// @description This is a go rgate api docs.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func NewHttpServer() {
	hdls := initDI()

	// init server
	ginServer := gin.Default()

	// swagger
	ginServer.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
