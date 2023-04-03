package user

import (
	"time"

	jsonplaceholder "github.com/Calmantara/go-dts-user/client/json-placeholder"
)

type GetUserResponse struct {
	Id    uint64    `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Dob   time.Time `json:"dob"`
}

type GetUserResponseWithTodo struct {
	GetUserResponse
	Todo jsonplaceholder.JsonPlaceholderResp `json:"todo"`
}

type CreateUserResponse struct {
	GetUserResponse
}
