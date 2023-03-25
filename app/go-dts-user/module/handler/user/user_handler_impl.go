package user

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Calmantara/go-dts-user/module/model"
	"github.com/Calmantara/go-dts-user/module/service/user"
	"github.com/Calmantara/go-dts-user/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserHdlImpl struct {
	userSvc user.UserService
}

func NewUserHandler(userSvc user.UserService) UserHandler {
	return &UserHdlImpl{
		userSvc: userSvc,
	}
}

func (u *UserHdlImpl) FindUserByIdHdl(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to find user",
			ErrorMsg: response.InvalidQuery,
		})
		return
	}
	// transform id string to uint64
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to find user",
			ErrorMsg: response.InvalidParam,
		})
		return
	}

	// call service
	user, err := u.userSvc.FindUserByIdSvc(ctx, idUint)
	if err != nil {
		if err.Error() == "user is not found" {
			ctx.JSON(http.StatusNotFound, response.ErrorResponse{
				Message:  "failed to find user",
				ErrorMsg: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message:  "failed to find user",
			ErrorMsg: response.SomethingWentWrong,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success find user",
		Data:    user,
	})
}

func (u *UserHdlImpl) FindAllUsersHdl(ctx *gin.Context) {
	users, err := u.userSvc.FindAllUsersSvc(ctx)
	if err != nil {
		// bad code, should be wrapped in other package
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message:  "failed to get users",
			ErrorMsg: response.SomethingWentWrong,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success get users",
		Data:    users,
	})
}

func (u *UserHdlImpl) InsertUserHdl(ctx *gin.Context) {
	// mendapatkan body
	var usrIn model.User

	if err := ctx.Bind(&usrIn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to insert user",
			ErrorMsg: response.InvalidBody,
		})
		return
	}

	// validate name and email
	if usrIn.Email == "" || usrIn.Name == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to insert user",
			ErrorMsg: response.InvalidParam,
		})
		return
	}

	insertedUser, err := u.userSvc.InsertUserSvc(ctx, usrIn)
	if err != nil {
		// bad code, should be wrapped in other package
		if err.Error() == "error duplication email" {
			ctx.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
				Message:  "failed to insert user",
				ErrorMsg: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message:  "failed to insert user",
			ErrorMsg: response.SomethingWentWrong,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Message: "success create user",
		Data:    insertedUser,
	})
}

func (u *UserHdlImpl) UpdateUserHdl(ctx *gin.Context) {
	idUint, err := u.getIdFromParam(ctx)
	if err != nil {
		return
	}
	// binding payload
	var usrIn model.User
	if err := ctx.Bind(&usrIn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to update user",
			ErrorMsg: response.InvalidBody,
		})
		return
	}
	usrIn.Id = idUint

	// validate name
	if usrIn.Name == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to update user",
			ErrorMsg: response.InvalidParam,
		})
		return
	}

	if err := u.userSvc.UpdateUserSvc(ctx, usrIn); err != nil {
		if err.Error() == "user is not found" {
			ctx.JSON(http.StatusNotFound, response.ErrorResponse{
				Message:  "failed to update user",
				ErrorMsg: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message:  "failed to update user",
			ErrorMsg: response.SomethingWentWrong,
		})
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Message: "success update user",
	})
}

func (u *UserHdlImpl) DeleteUserByIdHdl(ctx *gin.Context) {
	idUint, err := u.getIdFromParam(ctx)
	if err != nil {
		return
	}
	deletedUser, err := u.userSvc.DeleteUserByIdSvc(ctx, idUint)
	if err != nil {
		if err.Error() == "user is not found" {
			ctx.JSON(http.StatusNotFound, response.ErrorResponse{
				Message:  "failed to delete user",
				ErrorMsg: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message:  "failed to delete user",
			ErrorMsg: response.SomethingWentWrong,
		})
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Message: "success delete user",
		Data:    deletedUser,
	})
}

func (u *UserHdlImpl) getIdFromParam(ctx *gin.Context) (idUint uint64, err error) {
	id := ctx.Param("id")
	if id == "" {
		err = errors.New("failed id")
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to update user",
			ErrorMsg: response.InvalidParam,
		})
		return
	}
	// transform id string to uint64
	idUint, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		err = errors.New("failed parse id")
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to update user",
			ErrorMsg: response.InvalidParam,
		})
		return
	}
	return
}
