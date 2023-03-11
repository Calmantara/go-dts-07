package main

import (
	"fmt"
	"time"

	_ "github.com/Calmantara/go-dts-07/photo"
	"github.com/Calmantara/go-dts-07/user"
)

type (
	OperationFunc func(num1, num2 int) int // alias operation func
	Operation     int                      // alias operation type
)

const (
	ADD Operation = iota + 1
	SUB
	MUL
	DIV
)

// func ()  {}
// 1. sebelum () -> nama fungsi
// 2. di dalam (x) -> input dari fungsi
// 3. setelah () -> output
// 4. di dalam {. . .} -> code block

func add(num1, num2 int) int {
	// nama: add
	// input: num1, num2 int
	// output: int
	// block -> return num1 + num2
	return num1 + num2
}

func multiply(num1, num2, multiplier int) (int, int) {
	return num1 * multiplier, num2 * multiplier
}

// aku belum tau berapa banyak
// angka yang harus aku kali
// yang aku tau multiplier dan outputnya
// berupa array of integer
func multiplyMany(multiplier int, nums ...int) []int {
	// ... -> input variable menjadi optional
	var result []int

	for _, num := range nums {
		result = append(result, num*multiplier)
	}
	return result
}

func generateOperation(ops Operation) OperationFunc {
	var fn OperationFunc
	switch ops {
	case ADD:
		fn = func(num1, num2 int) int {
			return num1 + num2
		}
	case SUB:
		fn = func(num1, num2 int) int {
			return num1 - num2
		}
	case MUL:
		fn = func(num1, num2 int) int {
			return num1 * num2
		}
	case DIV:
		fn = func(num1, num2 int) int {
			return num1 / num2
		}
	default:
		fn = func(num1, num2 int) int {
			return 0
		}
	}
	return fn
}

func doOperation(num1, num2 int, fn OperationFunc) int {
	// kita akan menjalankan fn didalam doOperation
	res := fn(num1, num2)
	return res
}

func main() {
	// call function add
	res := add(1, 2)
	fmt.Println("hasil dari 1 + 2:", res)

	// multi output
	res1, res2 := multiply(3, 4, 2) // res1: 6, res2: 8
	fmt.Println("hasil dari multiply 3, 4 dengan 2:", res1, res2)

	// optional input param
	resArr := multiplyMany(5, 1, 2, 3, 4, 5, 6, 7, 8, 9, 123)
	fmt.Println("multiplyMany from numbers", resArr)

	inputArr := []int{3, 4, 5, 23, 13, 5, 2, 3, 2, 5, 4}
	resArr = multiplyMany(5, inputArr...)
	fmt.Println("multiplyMany from array", resArr)

	// closure -> function sebagai input / ouput dari function
	fn := generateOperation(ADD)
	fmt.Println("hasil dari generate ops div", fn(2, 1))
	// dijalankan di dalam doOperation
	fmt.Println("hasil dari doOperation fn", doOperation(3, 4, fn))

	// ga ada function yang aku inginin di dalam generate operation
	// tapi type sama sama operation function
	res = doOperation(2, 3, func(num1, num2 int) int {
		// function ini akan menjumlah num1 dan 2
		// lalu kali 5
		return (num1 + num2) * 5
	}) // 25
	fmt.Println("hasil dari custom doOperation:", res)

	// POINTER
	var pt *int
	num1 := 1

	// *pt = 100 // dengan ditambahin *, artinya kita mengakses value
	pt = &num1 // & menandakan dia memberikan address dari num1 ke pt1
	fmt.Println("address", pt, &num1)
	fmt.Println("value", *pt, num1)
	// ubah nilai dari num1
	num1 = 100
	fmt.Println("value", *pt, num1)

	// var arrpt *[]int //-> pointer of array of int
	// var arrOfPt []*int

	// STRUCT
	dob, _ := time.Parse("2006-01-02", "2000-03-03")
	calman := user.User{
		ID:    1,
		Name:  "Calmantara",
		Email: "calmantarasp@gmail.com",
		DOB:   dob,
	}
	fmt.Printf("isi dari calman:%+v\n", calman)
	fmt.Printf("nama lengkap:%v email:%v\n", calman.Name, calman.Email)

	// user.Greeting(calman)
	calman.Greeting()
	// exported atau unexported dari suatu package
	// exported -> diawali dengan huruf CAPITAL
	// unexported -> diawali dengan huruf KECIL
	// calman.calculateAge() -> tidak bisa dipanggil dari luat user package

	tara := user.Student{
		User: user.User{
			ID:    2,
			Name:  "Tara",
			Email: "tara@gmail.com",
			DOB:   dob,
		},
		Batch: 7,
	}
	tara.GreetingWithBatch()

	var interfaceOfInt interface{}
	interfaceOfInt = "this is string"

	numInt, ok := interfaceOfInt.(int)
	if !ok {
		fmt.Println("interface bukan int")
	} else {
		fmt.Println("interface adalah int:", numInt)
	}

	// mini challenge
	calman.AddAgeYear()
	calman.Greeting()
	//ptr
	calman.AddAgeYearPtr()
	calman.Greeting()
}
