package account

import "github.com/gin-gonic/gin"

type IAccountHandler interface {
	LoginAccount(ctx *gin.Context)
	CreateAccount(ctx *gin.Context)
	GetAccount(ctx *gin.Context)
}
