package main

import (
	"fmt"
)

// alias
type MyString string // MyString adalah alias dari string

const (
	key1 MyString = "key1"
	key2 MyString = "key2"
)

func main() {
	// array
	arrOfInt := make([]int, 5)
	fmt.Println("array of integer:", arrOfInt)

	// slice
	sliceOfInt := []int{}
	fmt.Println("slice of integer:", sliceOfInt)

	// array dan slice
	// index -> untuk melakukan action di value suatu array
	// arr    [0 0 0 0 0]
	// index   0 1 2 3 4
	arrOfInt[2] = 100
	fmt.Println("array of integer after edited index 2:", arrOfInt)

	// panic 1: mengakses memory di empty slice
	// sliceOfInt[2] = 100 // mau akses index ke 2
	// fmt.Println("slice of integer after edited index 2:", sliceOfInt)

	// panic 2:
	// arrOfInt[5] = 100
	// fmt.Println("array of integer after edited index 5:", arrOfInt)

	// memasukkan data ke empty slice
	sliceOfInt = append(sliceOfInt, 10)
	fmt.Println("slice of integer after appended:", sliceOfInt)
	// access index of slice
	fmt.Println("value in slice with index num 0:", sliceOfInt[0])

	// len in array
	// untuk mendapatkan banyaknya slot
	// [1,2,3,4] -> len(arr) = 4

	// [a:b] -> aku mau mengambil value array
	// mulai dari a sampai kurang dari b

	sliceOfInt = append(sliceOfInt, 20, 30, 40, 50) // [10, 20, 30, 40, 50]
	// cara pop atau mengeluarkan value paling terakhir dari array
	popedValue := sliceOfInt[len(sliceOfInt)-1]
	sliceOfInt = sliceOfInt[0 : len(sliceOfInt)-1] // aku cuma mau ngambil index 0 sampai sebelum terakhir
	fmt.Printf("value yang dikeluarkan:%v dan arr final:%v\n", popedValue, sliceOfInt)

	// mengambil 2 value awal dari array
	fmt.Println("dua value awal dari array adalah", sliceOfInt[0:2])

	// mini challenge
	slice1 := []int{1, 2, 3, 4}
	slice2 := slice1

	// mengubah slice 2
	slice2[2] = 5
	fmt.Printf("slice1:%v slice2:%v\n", slice1, slice2)

	// mini challenge
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := arr1

	// mengubah slice 2
	arr2[2] = 5
	fmt.Printf("arr1:%v arr2:%v\n", arr1, arr2)

	// copy
	// -> hanya bisa dilakukan
	// ketika len(dest) != 0
	slice3 := []int{1, 2, 3}
	slice4 := []int{4, 5}
	elmCopied := copy(slice4, slice3)
	fmt.Printf("elmCopied:%v, slice3:%v, slice4:%v\n", elmCopied, slice3, slice4)

	// mengubah element setelah di copy
	slice4[0] = 100
	fmt.Printf("slice3:%v, slice4:%v\n", slice3, slice4)

	// map -> map[key]value
	// key value pair
	var map1 map[string]string // map1 akan bernilai nil
	fmt.Println(map1)
	// assign value
	// map1 = make(map[string]string)
	map1 = map[string]string{}
	map1["key1"] = "value1"
	map1["key1"] = "value2"
	fmt.Printf("banyak key value %v, isinya:%v\n", len(map1), map1)

	// interface{} -> dynamic type
	map2 := map[string]interface{}{}
	map2["program"] = "dts kominfo"
	map2["kelas"] = 7
	fmt.Printf("banyak key value %v, isinya:%v\n", len(map2), map2)

	// untuk mengakses map
	// 1. langsung dengan key
	// 2. lebih secure, key with validation

	// cara langsung
	program := map2["program"]
	fmt.Println(program)
	baju := map2["koko"]
	baju = "ini baju koko"
	fmt.Println(baju)

	// cara aman
	program, ok := map2["program"] // value, bool -> bool akan true ketika nilainya ada
	if ok {
		fmt.Println("OKE nemu value dari kunci program", program)
	} else {
		fmt.Println("woi lokernya kosong")
	}

	baju, ok = map2["koko"] // value, bool -> bool akan true ketika nilainya ada
	if ok {
		fmt.Println("OKE nemu value dari kunci koko", baju)
	} else {
		fmt.Println("woi lokernya kosong")
	}

	// catatan:
	// 1. map tidak bisa dijadikan constant (cari tau sendiri)
	// 2. akan lebih aman, kalau menggunakan type sendiri

	var myStringVariable MyString
	var myStringNormal string

	myStringVariable = "hello world"
	myStringNormal = "hello world"

	// secara syntax MyString != string
	// ketika mau compare, harus di type cast
	fmt.Println(myStringVariable == MyString(myStringNormal))

	map3 := map[MyString]interface{}{}
	map3[key1] = "hello world"
	map3[key2] = "hello world juga"
	fmt.Println("value sebelum di delete", map3)
	delete(map3, key2)
	fmt.Println("value sesudah di delete", map3)

	// condition
	// if(){} -> didalam kurung nanti harus komparasi
	// atau suatu variable bool
	// else if -> nested condition
	// else -> kondisi ketika tidak ada yang memenuhi

	socialMedia := "facebook"
	if socialMedia == "twitter" {
		fmt.Println("elon")
	} else if socialMedia == "facebook" {
		fmt.Println("mark")
	} else {
		fmt.Println("ansos")
	}

	// case 1
	switch socialMedia {
	case "twitter":
		fmt.Println("elon")
	case "facebook":
		fmt.Println("mark")
		fallthrough // ketika masuk ke facebook -> dia akan mengeksekusi bawahnya juga
	case "whatsapp":
		fmt.Println("mark lagi dong")
	default:
		fmt.Println("ansos")
	}

	// case 2
	switch {
	case socialMedia == "twitter":
		fmt.Println("elon")
	case socialMedia == "facebook":
		fmt.Println("mark")
	default:
		fmt.Println("ansos")
	}

	var nilable interface{}
	switch nilable {
	case nil:
		fmt.Println("ini nil lho")
	default:
		fmt.Println("gak nil")
	}

	number := 10
	if number > 10 {
		fmt.Println("gede")
	} else if number > 0 && number <= 5 {
		fmt.Println("normal lah")
	} else if number == 10 || number == 100 {
		fmt.Println("ok")
	}

	exist := false
	// exist di atas ini, adalah global variable

	if _, exist := map3[key1]; exist {
		// exist di sini, akan menjadi
		// local variable di dalam block ini doang

		// aku ngambil val dan exist dari map
		// kalau dia exist / exist == true
		// dia akan masuk ke block ini
		fmt.Println("exist nih, masuk ke block", exist)
	}
	fmt.Println(exist)

	// looping
	// aku mau ngelooping 10 kali
	for i := 0; i < 10; i++ {
		// 1. kita declare i sebagai integer
		// 2. kita check apakah i < 10,
		// 		jika i >= 10, dia akan keluar dari loop
		// 3. i++ -> i = i + 1 -> i += 1
		fmt.Println("print yang ke-", i)
	}

	// infinit loop
	i := 0
	for {
		i = i + 1 // incremental
		if i >= 10 {
			// break -> program akan keluar dari block looping
			break
		} else {
			if i == 5 {
				// continue -> dia akan melewatkan
				// semua process yang ada di bawahnya
				// yang berada di dalam loop block
				continue
			}
		}
		// 1 + 1
		fmt.Println("print ke 2 untuk i-", i)
	}

	// loop range
	// untuk mengakses map / array,
	// kita bisa menggunakan key / index
	// jika kita ingin mengakses semua isinya,
	// kita bisa menggunakan loop range

	arrRange := [5]int{100, 200, 300, 400, 500}
	for idx, val := range arrRange {
		fmt.Printf("index:%v val:%v\n", idx, val)
	}

	mapRange := map[string]int{
		"key1": 100,
		"key2": 200,
		"key3": 300,
		"key4": 400,
		"key5": 500,
	}
	for key, val := range mapRange {
		fmt.Printf("key:%v val:%v\n", key, val)
	}

	// MINI CHALLENGE
	// 1. diberikan array [1,1,2,5,4,7,3,4,4,6,5,7,9,9]
	// 2. printlah value dalam array yang berjumlah > 1
	// 3. expected output: 1, 4, 5, 7, 9

	inputArray := []int{1, 1, 2, 5, 4, 7, 3, 4, 4, 6, 5, 7, 9, 9}
	duplicate := map[int]int{}

	for _, val := range inputArray {
		duplicate[val] += 1
	}

	for key, val := range duplicate {
		if val > 1 {
			fmt.Println(key)
		}
	}
}
