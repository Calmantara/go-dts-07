package user

import "time"

type CreateUser struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Dob   time.Time `json:"dob"`
}

type UpdateUser struct {
	Id    uint64    `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Dob   time.Time `json:"dob"`
}

type BasicUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
