package user

import (
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	FindUserByIdHdl(ctx *gin.Context)
	FindAllUsersHdl(ctx *gin.Context)
	InsertUserHdl(ctx *gin.Context)
	UpdateUserHdl(ctx *gin.Context)
	DeleteUserByIdHdl(ctx *gin.Context)
}
