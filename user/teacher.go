package user

import (
	"fmt"

	"github.com/segmentio/fasthash/fnv1a"
)

type Teacher struct {
	Birthday // copy struct level
	NIP     int
	Name    string
	Email   string
	Subject string
}

// gimana caranya kita mengimplementasikan
// abstract interface ke teacher
// sehingga teacher termasuk ke
// golongan user

// DENGAN CARA
// MENGIMPLEMENTASIKAN FUNGSI ABSTRACK
// SEBAGAI METHOD

func (t Teacher) GeneratePassword() (password string) {
	password = fmt.Sprintf(
		"%v%v%v",
		fnv1a.HashString64(t.Email+t.Name+t.Subject),
		t.NIP,
		t.Date,
	) // random uint
	return
}

func (t Teacher) GetUsername() (username string) {
	username = t.Name
	return
}

func (t Teacher) GetLocation() (ret any) {
	return
}
