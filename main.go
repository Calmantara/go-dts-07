package main

import (
	"fmt"

	"github.com/Calmantara/go-dts-07/user"
)

const (
	pi = 3.14
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return pi * c.radius * c.radius
}

type rect struct {
	length float64
}

func (r rect) area() float64 {
	return r.length * r.length
}

type printer struct{}

func (p printer) printArea(s shape) {
	fmt.Println("area of shape is:", s.area())
}

func (p printer) printSumArea(s ...shape) {
	sum := 0.0
	for _, sh := range s {
		sum += sh.area()
	}
	fmt.Println("sum area of shapes is:", sum)
}

type DB interface {
	Connect()
}

type MySql struct{}

func (m MySql) Connect() {
	fmt.Println("connected to mysql server")
}

type Postgres struct{}

func (p Postgres) Connect() {
	fmt.Println("connected to postgres server")
}

type Mongo struct{}

func (m Mongo) Connect() {
	fmt.Println("connected to mongo server")
}

type MyApp struct {
	Db DB
}

func (m MyApp) ConnectToDB() {
	m.Db.Connect()
}

func main() {
	// interface as variable
	var interface1 interface{}

	// ada variable integer
	int1 := 100
	// assign interface as integer
	interface1 = int1
	interface1 = 1000
	fmt.Printf("interface:%v integer:%v\n", interface1, int1)
	// re assign to string
	interface1 = "HELLO WORLD" // interface of string
	fmt.Printf("interface:%v after reassign\n", interface1)

	// type cast interface of string
	var1, ok := interface1.(string) // var1 akan mendapatkan value dari interface
	fmt.Printf("status:%v variable type casted:%v\n", ok, var1)

	// type cast interface of string to int
	var2, ok := interface1.(int) // avoid panic
	fmt.Printf("status: %v variable type casted to int:%v\n", ok, var2)

	// cara lainnya
	// 1. reflect
	// 2. json marshal & unmarshal -> struct

	// map of interface
	mapInterface := make(map[string]any) // any = alias dari interface
	mapInterface["int"] = 100
	mapInterface["string"] = "hello world"
	mapInterface["float"] = 100.0
	fmt.Printf("map of interface: %+v\n", mapInterface)
	// map[string]any atau map[string]interface{}
	// dia bisa di transform menjadi struct

	var user1 user.User // <- interface of user
	user1 = user.Teacher{
		NIP:     123456789,
		Email:   "tara@gmail.com",
		Name:    "tara",
		Subject: "golang",
	}
	fmt.Printf("declaration of user1:%+v\n", user1)
	fmt.Println(
		user1.GeneratePassword(),
		user1.GetUsername(),
	)

	// reassign to student
	user1 = user.Student{
		NIM:   987654321,
		Email: "calman@gmail.com",
		Class: "Batch7",
	}
	fmt.Printf("declaration of user1 to student:%+v\n", user1)
	fmt.Println(
		user1.GeneratePassword(),
		user1.GetUsername(),
	)

	// case tidak bisa reassign
	// dari struct ke struct lain
	// meskipun satu golongan
	// interface

	// student := user.Student{
	// 	NIM:   987654321,
	// 	Email: "calman@gmail.com",
	// 	Class: "Batch7",
	// }
	// student = user.Teacher{
	// 	NIP:     123456789,
	// 	Email:   "tara@gmail.com",
	// 	Name:    "tara",
	// 	Subject: "golang",
	// }
	// fmt.Println(student)

	// student := &user.Student{
	// 	NIM:   987654321,
	// 	Email: "calman@gmail.com",
	// 	Class: "Batch7",
	// }
	// student = &user.Teacher{
	// 	NIP:     123456789,
	// 	Email:   "tara@gmail.com",
	// 	Name:    "tara",
	// 	Subject: "golang",
	// }

	circleShape := circle{radius: 4}
	r := rect{length: 10}

	printerVar := printer{}
	printerVar.printArea(circleShape)
	printerVar.printArea(r)
	printerVar.printSumArea(circleShape, r)

	//
	myApp := MyApp{
		Db: Postgres{},
	}
	myApp.ConnectToDB()
	/// postgres lemot harus ganti ke my sql

	myApp.Db = MySql{}
	myApp.ConnectToDB()
}
