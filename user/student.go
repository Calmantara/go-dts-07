package user

import (
	"fmt"

	"github.com/segmentio/fasthash/fnv1a"
)

type Student struct {
	NIM   int
	Email string
	Class string
}

func (s Student) GeneratePassword() (password string) {
	password = fmt.Sprintf(
		"%v",
		fnv1a.HashString64(s.Email+s.Class),
	) // random uint
	return
}

func (s Student) GetUsername() (username string) {
	username = s.Email
	return
}

func (s Student) GetLocation() (ret any) {
	return
}
