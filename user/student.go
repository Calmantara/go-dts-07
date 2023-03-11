package user

import "fmt"

type Student struct {
	User
	Batch int
}

func (s Student) GreetingWithBatch() {
	fmt.Printf("[%v] Hallo, %v %v tahun dari batch %v\n",
		generateID(),
		s.Name,
		s.calculateAge(),
		s.Batch)
}

func init() {
	fmt.Println("HALO INI DARI INIT FUNCTION PACKAGE USER")
}
