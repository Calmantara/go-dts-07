package account

import (
	accounthandler "github.com/Calmantara/go-account/modules/handler/account"
	"github.com/Calmantara/go-account/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func NewAccountRouter(v1 *gin.RouterGroup, accountHdl accounthandler.IAccountHandler) {
	g := v1.Group("/account")

	// register all router
	g.POST("",
		accountHdl.CreateAccount)
	g.POST("/login",
		accountHdl.LoginAccount)
	g.GET("",
		middleware.BearerOAuth(),
		accountHdl.GetAccount)
}
