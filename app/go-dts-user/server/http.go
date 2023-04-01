package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Calmantara/go-common/config"
	"github.com/Calmantara/go-common/pkg/middleware"
	"github.com/Calmantara/go-dts-user/docs"
	"github.com/Calmantara/go-dts-user/module/router/v1/user"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GO DTS USER API DUCUMENTATION
// @version 2.0
// @description go-dts-user api documentation
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9090
// @BasePath /
// @schemes http
func NewHttpServer() (srv *http.Server) {
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
		middleware.CorrelationIDInterceptor(),
	)

	// register router
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := ginServer.Group("/api/v1")
	user.NewUserRouter(v1, hdls.userHdl)

	// swagger
	ginServer.GET("/go-dts-user/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
