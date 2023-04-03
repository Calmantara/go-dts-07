package user

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	response "github.com/Calmantara/go-common/pkg/response"
	model "github.com/Calmantara/go-dts-user/module/model/user"
	"github.com/Calmantara/go-dts-user/module/service/user"
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

// @BasePath /api/v1/user
// @Accept json -> api kita akan menerima payload dalam bentuk apa
// @Produce json -> api kita akan produce response dalam bentuk apa
// @Param id query int true "user id" -> namaparam tipeparam datatipe mandatory desc

// Find User Detail
// @Tags User
// @Summary finding user record
// @Schemes http
// @Description fetch user information by id
// @Accept json
// @Produce json
// @Param Authorization header string true "basic authentication"
// @Param id query int true "user id"
// @Success 200 {object} response.SuccessResponse{data=model.GetUserResponseWithTodo}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 422 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /user [get]
func (u *UserHdlImpl) FindUserByIdHdl(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "failed to find user",
			Error:   response.InvalidQuery,
		})
		return
	}

	userBasic := ctx.GetStringMapString("userBasic")
	fmt.Println(userBasic)

	// transform id string to uint64
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "failed to find user",
			Error:   response.InvalidParam,
		})
		return
	}

	// call service
	user, err := u.userSvc.FindUserByIdSvc(ctx, idUint)
	if err != nil {
		if err.Error() == "user is not found" {
			ctx.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
				Message: "failed to find user",
				Error:   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: "failed to find user",
			Error:   response.SomethingWentWrong,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success find user",
		Data:    user,
	})
}

// Find All User
// @Tags User
// @Summary finding all user record
// @Schemes http
// @Description fetch all user information
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]model.GetUserResponse}
// @Param Authorization header string true "basic authentication"
// @Router /user/all [get]
func (u *UserHdlImpl) FindAllUsersHdl(ctx *gin.Context) {
	users, err := u.userSvc.FindAllUsersSvc(ctx)
	if err != nil {
		// bad code, should be wrapped in other package
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: "failed to get users",
			Error:   response.SomethingWentWrong,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success get users",
		Data:    users,
	})
}

// Insert New User
// @Tags User
// @Summary insert new user info
// @Schemes http
// @Description insert new user info
// @Accept json
// @Param user body model.CreateUser true "user payload"
// @Produce json
// @Param Authorization header string true "basic authentication"
// @Success 202 {object} response.SuccessResponse{data=model.CreateUserResponse}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 422 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /user [post]
func (u *UserHdlImpl) InsertUserHdl(ctx *gin.Context) {
	// mendapatkan body
	var usrIn model.CreateUser

	if err := ctx.Bind(&usrIn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "failed to insert user",
			Error:   response.InvalidBody,
		})
		return
	}

	// validate name and email
	if usrIn.Email == "" || usrIn.Name == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "failed to insert user",
			Error:   response.InvalidParam,
		})
		return
	}

	insertedUser, err := u.userSvc.InsertUserSvc(ctx, usrIn)
	if err != nil {
		// bad code, should be wrapped in other package
		if err.Error() == "error duplication email" {
			ctx.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
				Message: "failed to insert user",
				Error:   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: "failed to insert user",
			Error:   response.SomethingWentWrong,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Message: "success create user",
		Data:    insertedUser,
	})
}

// Update User Detail
// @Tags User
// @Summary update user detail
// @Schemes http
// @Description update user info by id
// @Accept json
// @Param user body model.UpdateUser false "user payload"
// @Param id path int true "user id"
// @Param Authorization header string true "basic authentication"
// @Produce json
// @Success 202 {object} response.SuccessResponse{}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 422 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /user/{id} [put]
func (u *UserHdlImpl) UpdateUserHdl(ctx *gin.Context) {
	idUint, err := u.getIdFromParam(ctx)
	if err != nil {
		return
	}
	// binding payload
	var usrIn model.UpdateUser
	if err := ctx.Bind(&usrIn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "failed to update user",
			Error:   response.InvalidBody,
		})
		return
	}
	usrIn.Id = idUint

	// validate name
	if usrIn.Name == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "failed to update user",
			Error:   response.InvalidParam,
		})
		return
	}

	if err := u.userSvc.UpdateUserSvc(ctx, usrIn); err != nil {
		if err.Error() == "user is not found" {
			ctx.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
				Message: "failed to update user",
				Error:   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: "failed to update user",
			Error:   response.SomethingWentWrong,
		})
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Message: "success update user",
	})
}

// Delete User Detail
// @Tags User
// @Summary soft delete user record
// @Schemes http
// @Description add deleted_at param flag
// @Accept json
// @Param id path int true "user id"
// @Param Authorization header string true "basic authentication"
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=model.GetUserResponse}
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 422 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
// @Router /user/{id} [delete]
func (u *UserHdlImpl) DeleteUserByIdHdl(ctx *gin.Context) {
	idUint, err := u.getIdFromParam(ctx)
	if err != nil {
		return
	}
	deletedUser, err := u.userSvc.DeleteUserByIdSvc(ctx, idUint)
	if err != nil {
		if err.Error() == "user is not found" {
			ctx.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
				Message: "failed to delete user",
				Error:   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: "failed to delete user",
			Error:   response.SomethingWentWrong,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success delete user",
		Data:    deletedUser,
	})
}

func (u *UserHdlImpl) getIdFromParam(ctx *gin.Context) (idUint uint64, err error) {
	id := ctx.Param("id")
	if id == "" {
		err = errors.New("failed id")
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "failed to update user",
			Error:   response.InvalidParam,
		})
		return
	}
	// transform id string to uint64
	idUint, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		err = errors.New("failed parse id")
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "failed to update user",
			Error:   response.InvalidParam,
		})
		return
	}
	return
}
