package user

import "time"

type GetUserResponse struct {
	Id    uint64    `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Dob   time.Time `json:"dob"`
}

type CreateUserResponse struct {
	GetUserResponse
}
