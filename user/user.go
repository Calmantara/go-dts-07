package user

// abstraksi
// user, harus memiliki setidaknya fungsi dibawah ini
// 1. generate password
// 2. get username

type Birthday struct {
	Date  string
	Place string
}

type User interface {
	// copy interface level
	Location // copy instead of extend
	GeneratePassword() (password string)
	GetUsername() (username string)
}

type Location interface {
	GetLocation() any
}

// implementasi user disini
// 1. teacher
// 2. student
