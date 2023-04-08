package account

import (
	"errors"
	"net/http"

	accountmodel "github.com/Calmantara/go-account/modules/models/account"
	"github.com/Calmantara/go-account/modules/models/token"
	accountservice "github.com/Calmantara/go-account/modules/service/account"
	"github.com/Calmantara/go-account/pkg/middleware"
	"github.com/Calmantara/go-common/pkg/json"
	"github.com/Calmantara/go-common/pkg/logger"
	"github.com/Calmantara/go-common/pkg/response"
	"github.com/gin-gonic/gin"
)

type AccountHandlerImpl struct {
	accService accountservice.IAccountService
}

func NewAccountHandlerImpl(accService accountservice.IAccountService) IAccountHandler {
	return &AccountHandlerImpl{
		accService: accService,
	}
}

func (a *AccountHandlerImpl) LoginAccount(ctx *gin.Context) {
	// binding payload
	var loginAccount accountmodel.LoginAccount
	if err := ctx.BindJSON(&loginAccount); err != nil {
		logger.Error(ctx, "error binding payload",
			"error", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			response.ErrorResponse{
				Message: response.InvalidBody,
				Error:   "error binding payload",
			},
		)
		return
	}

	tokens, err := a.accService.LoginAccountByUserName(ctx, loginAccount)
	if err != nil {
		logger.Error(ctx, "error create account",
			"error", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			response.ErrorResponse{
				Message: response.InternalServer,
				Error:   response.SomethingWentWrong,
			},
		)
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Message: "success created",
		Data:    tokens,
	})
}

func (a *AccountHandlerImpl) CreateAccount(ctx *gin.Context) {
	// binding payload
	var createAccount accountmodel.CreateAccount
	if err := ctx.BindJSON(&createAccount); err != nil {
		logger.Error(ctx, "error binding payload",
			"error", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			response.ErrorResponse{
				Message: response.InvalidBody,
				Error:   "error binding payload",
			},
		)
		return
	}

	created, err := a.accService.CreateAccount(ctx, createAccount)
	if err != nil {
		logger.Error(ctx, "error create account",
			"error", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			response.ErrorResponse{
				Message: response.InternalServer,
				Error:   response.SomethingWentWrong,
			},
		)
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Message: "success created",
		Data:    created,
	})
}

func (a *AccountHandlerImpl) GetAccount(ctx *gin.Context) {
	// get user_id from context first
	accessClaimI, ok := ctx.Get(middleware.AccessClaim.String())
	if !ok {
		err := errors.New("error get claim from context")
		logger.Error(ctx, "error get payload",
			"error", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: response.InvalidPayload,
			Error:   "invalid user id",
		})
		return
	}

	var accessClaim token.AccessClaim
	if err := json.ObjectMapper(accessClaimI, &accessClaim); err != nil {
		logger.Error(ctx, "error mapping object payload",
			"error", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Message: response.InvalidPayload,
			Error:   "invalid payload",
		})
		return
	}

	account, err := a.accService.GetAccount(ctx, accessClaim.UserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: response.InternalServer,
			Error:   "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success",
		Data:    account,
	})
}
