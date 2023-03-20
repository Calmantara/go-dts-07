package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/pkg/errors"
)

type (
	Item struct {
		Amount int
	}

	Customer struct {
		Balance uint64
		Items   []Item
	}
)

var (
	customers = []Customer{
		{
			Balance: 100000,
			Items:   []Item{},
		},
		{
			Balance: 100000,
			Items: []Item{
				{Amount: -10},
				{Amount: 1000},
			},
		},
		{
			Balance: 100000,
			Items: []Item{
				{Amount: 1000},
				{Amount: 2000},
			},
		},
	}
)

func main() {
	// error sangat berguna
	// untuk mendeteksi ketidak sesuaian
	// input atau process

	defer func() {
		// recover hanya akan
		// menangkap value panic
		// biasanya func recover
		// digunakan oleh web server
		// untuk auto restart!!!!
		// udah tersedia di web framework!
		// gin/echo/http
		r := recover()
		if r != nil {
			fmt.Println("WUT?? PANIC BOSS")
			return
		}
		fmt.Println("oh yaudah error biasa / di kill")
	}()

	// akan mengconvert string menjadi int
	var input string
	fmt.Scanln(&input)
	// panic handling

	// panic -> kondisi yang memaksa
	// program untuk keluar

	// err := connectToDB(input)
	// fmt.Println(err)

	// apakah kita bisa 100% menghindari panic?
	// tidak
	var intPtr *int
	fmt.Println(*intPtr * 1000)

	// recover -> untuk menangkap panic
}

func connectToDB(password string) error {
	// handle error / false condition
	// with panic
	if password == "" {
		panic("password is invalid empty")
	} else if password != "mysupersecretpassword" {
		return errors.New("invalid password")
	}
	return nil
}

func convert(in string) (val int, err error) {
	// handling error in JS
	// - catchErr

	val, err = strconv.Atoi(in)
	// 1. error handling
	if err != nil {
		// 1. error terjadi di mana sih?
		// 2. siapa yang manggil error?

		// menggunakan error tracer
		// jaeger
		// datadog
		// newrelic
		// open telemetry

		// error aslinya dari Atoi,
		// cukup kita log
		// sehingga hanya developer yang mengetahui
		err = errors.WithStack(err)
		fmt.Printf("ERROR BOS!!!:%+v\n", err)

		// custom error
		// err = errors.New("something went wrong")
		err = fmt.Errorf("invalid input:%v", in)
		return
	}
	fmt.Println("OK:", val)
	return
}

func exitExample() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for id, cust := range customers {
			fmt.Println("customer num:", id)
			CalculateBillRate(cust)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)

		// 0 -> keluar dengan baik baik (success)
		// != 0 -> keluar dengan error
		os.Exit(1)
	}()
	wg.Wait()
}

func CalculateBillRate(c Customer) {
	// definisi paling akhir dari function ini
	// adalah ketika sebelum return
	var (
		sumAmount int
		sumTax    float64
	)

	// golang ada built in function
	// yang ensure suatu process
	// dijalankan paling akhir dari function
	// DEFER

	// Cari sendiri jawabannya :)
	//1
	defer getCustomerBalance(c, sumAmount, sumTax)
	//2
	defer func() {
		getCustomerBalance(c, sumAmount, sumTax)
	}()

	if len(c.Items) <= 0 {
		sumAmount, sumTax = 0, 0.0
		fmt.Println("nothing to do here, item is empty")
		// return 1
		// getCustomerBalance(c, sumAmount, sumTax)
		return
	}

	for _, item := range c.Items {
		if item.Amount <= 0 {
			fmt.Println("invalid item amount")
			// return 2
			// getCustomerBalance(c, sumAmount, sumTax)
			return
		}
		sumAmount += item.Amount
		if item.Amount > 1000 {
			sumTax += float64(item.Amount) * 0.1
		}
	}
	// getCustomerBalance(c, sumAmount, sumTax)
	// return 3
}

func getCustomerBalance(c Customer, sumAmount int, sumTax float64) {
	if sumAmount <= 0 {
		fmt.Println("amount is zero")
		return
	}

	if c.Balance < (uint64(sumAmount) + uint64(math.Ceil(sumTax))) {
		fmt.Println("issuficient balance")
		return
	}

	c.Balance -= (uint64(sumAmount) + uint64(math.Ceil(sumTax)))
	fmt.Println("balance deducted, final balance is:", c.Balance)
}

// menangkap sinyal dari os
// akan sangat dibutuhkan untuk
// membuat web server (backend)
func signalUnix() {
	// membuat 1 channel
	// untuk menangkap sinyal
	sigs := make(chan os.Signal, 1)
	// memasukkan channel ke
	// procedure penangkapan sinyal
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	done := make(chan bool, 1)

	go func() {
		// assign sigs channel
		// to variable
		// jika sigs menangkap signal
		// yang diberikan oleh OS kita

		// karena sigs adalah channel
		// dia akan stuck di process ini
		// sampai dia mendapatkan sinyal

		// sigs akan mendapatkan sinyal /
		// diassign suatu value, di signal.Notify
		sig := <-sigs
		fmt.Println("signal received")
		switch sig {
		case syscall.SIGINT:
			fmt.Println("killed by ctrl+c:", sig)
		case syscall.SIGTERM:
			fmt.Println("killed by system:", sig)
		case syscall.SIGHUP:
			fmt.Println("hot reloading by system:", sig)
		default:
			fmt.Println("received other signal:", sig)
		}

		// process end of signal
		done <- true
	}()

	// apakah os.Exit
	// adalah termasuk dari signal?
	// os.exit tidak termasuk kedalam signal.
	// signal hanya perintah exit yang dikirimkan
	// oleh OS ke program GO
	// sedangkan os.exit, keinginan program go sendiri
	// untuk keluar

	go func() {
		time.Sleep(2 * time.Second)

		// 0 -> keluar dengan baik baik (success)
		// != 0 -> keluar dengan error
		os.Exit(1)
	}()

	fmt.Println("awaiting signal")
	// ketika done tidak mendapatkan value
	// dia akan stuck di process ini

	// done akan mendapatkan value
	// process end of signal
	<-done
	fmt.Println("exiting")
}
