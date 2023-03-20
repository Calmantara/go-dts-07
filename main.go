package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type (
	User struct {
		Id     uint64 `json:"id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Delete bool   `json:"-"`
		// - di json, menandakan golang akan mengabaikan properti
	}

	Response struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

const (
	// init config
	PORT = ":9090"
)

var (
	DataStore = make(map[uint64]User)
)

func main() {
	DataStore = map[uint64]User{
		1: {
			Id:     1,
			Name:   "Calman",
			Email:  "calman@gmail.com",
			Delete: false,
		},
		2: {
			Id:     2,
			Name:   "Tara",
			Email:  "tara@gmail.com",
			Delete: false,
		},
	}

	// init server
	ginServer := gin.Default()
	// init middleware
	ginServer.Use(
		gin.Logger(),   // untuk log request yang masuk
		gin.Recovery(), // untuk auto restart kalau panic
	)
	// init router

	// get all users
	// api ini akan menampilkan
	// semua users dari data store
	ginServer.GET("/users", func(ctx *gin.Context) {
		// FUNGSI HANDLER: layer yang berperan untuk
		// menangkap request

		// business logic
		// pisahkan menjadi business logic layer
		var users []User
		for _, usr := range DataStore {
			if !usr.Delete {
				users = append(users, usr)
			}
		}

		// response
		ctx.JSON(http.StatusOK, Response{
			Message: "success get users",
			Data:    users,
		})
	})
	// get user by id
	// api ini akan memunculkan detail
	// dari user oleh id yang di filter
	ginServer.GET("/user", func(ctx *gin.Context) {
		// mendapatkan id dari query
		id := ctx.Query("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, Response{
				Message: "invalid query param",
			})
			return
		}
		// transform id string to uint64
		idUint, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Message: "invalid id param",
			})
			return
		}

		// busines logic
		// get from data store
		user, ok := DataStore[idUint]
		if !ok || user.Delete { // berarti tidak ada data
			ctx.JSON(http.StatusNotFound, Response{
				Message: "user is not found",
			})
			return
		}

		// response
		ctx.JSON(http.StatusOK, Response{
			Message: "success get user",
			Data:    user,
		})
	})
	// create user
	// api ini akan menambahkan daftar user
	ginServer.POST("/user", func(ctx *gin.Context) {
		// API
		// body dalam json:
		// name <string>
		// email <string>

		// mendapatkan body
		var usrIn User

		// KENAPA HARUS &body??
		if err := ctx.Bind(&usrIn); err != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Message: "invalid payload",
			})
			return
		}

		// validate name and email
		if usrIn.Email == "" || usrIn.Name == "" {
			ctx.JSON(http.StatusBadRequest, Response{
				Message: "invalid payload",
			})
			return
		}

		// business logic
		// tidak ada yang boleh memiliki email yang sama
		for _, usr := range DataStore {
			if strings.EqualFold(usr.Email, usrIn.Email) {
				// equal fold dia akan mengabaikan
				// besar kecilnya huruf
				// CALMAN@GMAIL.COM == calman@gmail.com
				ctx.JSON(http.StatusBadRequest, Response{
					Message: "duplicate user",
				})
				return
			}
		}
		// if success
		idInsert := len(DataStore) + 1
		usrIn.Id = uint64(idInsert)
		DataStore[uint64(idInsert)] = usrIn

		ctx.JSON(http.StatusAccepted, Response{
			Message: "success create user",
			Data:    usrIn,
		})
	})
	// update user
	ginServer.PUT("/user/:id", func(ctx *gin.Context) {
		// :id
		// bisa mendapaykan path param
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, Response{
				Message: "invalid query param",
			})
			return
		}
		// transform id string to uint64
		idUint, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Message: "invalid id param",
			})
			return
		}

		// binding payload
		var usrIn User
		if err := ctx.Bind(&usrIn); err != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Message: "invalid payload",
			})
			return
		}

		// validate name
		if usrIn.Name == "" {
			ctx.JSON(http.StatusBadRequest, Response{
				Message: "invalid payload",
			})
			return
		}

		// business logic
		// get existing user data
		user, ok := DataStore[idUint]
		if !ok || user.Delete { // berarti tidak ada data
			ctx.JSON(http.StatusNotFound, Response{
				Message: "user is not found",
			})
			return
		}

		user.Name = usrIn.Name
		DataStore[idUint] = user

		ctx.JSON(http.StatusAccepted, Response{
			Message: "success update user",
		})
	})
	// delete user
	ginServer.DELETE("/user/:id", func(ctx *gin.Context) {
		// :id
		// bisa mendapaykan path param
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, Response{
				Message: "invalid query param",
			})
			return
		}
		// transform id string to uint64
		idUint, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Message: "invalid id param",
			})
			return
		}

		// business logic
		// get existing user data
		user, ok := DataStore[idUint]
		if !ok || user.Delete { // berarti tidak ada data
			ctx.JSON(http.StatusNotFound, Response{
				Message: "user is not found",
			})
			return
		}

		// soft delete
		user.Delete = true
		DataStore[idUint] = user
		ctx.JSON(http.StatusAccepted, Response{
			Message: "success delete user",
		})
	})
	// default port is 8080
	ginServer.Run(PORT)
}
