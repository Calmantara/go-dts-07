package user

import (
	"github.com/Calmantara/go-dts-07/module/handler/user"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(v1 *gin.RouterGroup, userHdl user.UserHandler) {
	g := v1.Group("/user")

	// register all router
	g.GET("/all", userHdl.FindAllUsersHdl)
	g.GET("", userHdl.FindUserByIdHdl)
	g.POST("", userHdl.InsertUserHdl)
	g.PUT("/:id", userHdl.UpdateUserHdl)
	g.DELETE("/:id", userHdl.DeleteUserByIdHdl)
}
