package user

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	ID    uint
	Name  string
	Email string
	DOB   time.Time
}

func (user User) Greeting() {
	fmt.Printf("Hallo, %v %v tahun\n", user.Name, user.calculateAge())
}

func (user User) AddAgeYear() {
	fmt.Println(user.DOB)
	user.DOB = user.DOB.AddDate(-1, 0, 0)
	fmt.Println(user.DOB)
}

func (user *User) AddAgeYearPtr() {
	user.DOB = user.DOB.AddDate(-1, 0, 0)
}

// aku mau ada 1 function
// yang hanya bisa diakses oleh package user
// untuk menghitung umur

func (user User) calculateAge() int {
	return int(time.Since(user.DOB).Hours()) / (365 * 24)
}

func generateID() int {
	return rand.Int()
}
