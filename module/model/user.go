package model

import "time"

type User struct {
	Id     uint64    `json:"id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Dob    time.Time `json:"dob"`
	Delete bool      `json:"-"`
	// - di json, menandakan golang akan mengabaikan properti
}
