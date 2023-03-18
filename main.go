package main

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"

	eg "golang.org/x/sync/errgroup"
)

func main() {
	// di golang
	// ada 3 yang perlu diketahuin
	// untuk menjalankan concurrency

}

func fanOutPattern() {
	// aku hanya mau
	// spawn 100 goroutine
	// untuk menjalankan longMultiply
	chanInt := generateDataChan()
	for i := 0; i < 100; i++ {
		poolFunc(chanInt)
	}
	fmt.Println("num of goroutines",
		runtime.NumGoroutine())
	time.Sleep(10 * time.Second)
}

func generateDataChan() <-chan int {
	chanInt := make(chan int)
	go func() {
		defer close(chanInt)
		for i := 0; i < 10000; i++ {
			chanInt <- i
		}
	}()
	return chanInt
}

func poolFunc(chanInt <-chan int) {
	go func() {
		for i := range chanInt {
			longMultiply(i)
		}
	}()
}

func goroutineLeaked() {
	// goroutine leaked
	for i := 0; i < 10000; i++ {
		go longMultiply(i)
	}
	fmt.Println("num of goroutines",
		runtime.NumGoroutine())
}

func dataRaceExp() {
	// DATA RACE
	x := 10
	y := 0
	go func() {
		y = x + 10
	}()
	fmt.Println("value of y", y)
}

func deadlockExp() {
	// deadlock
	chan1 := make(chan int)
	// chanData menunggu chan1 untuk memberikan data
	chanData := <-chan1
	// chan1 baru dapat data dibawahnya
	chan1 <- 100
	fmt.Println(chanData)
}

func deferExp() {
	defer func() {
		fmt.Println("hello world in the end")
	}()

	sum := 1
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(i)
	}
	if sum%2 == 0 {
		return
	}
	fmt.Println("after modulo")
}

func concurrent1() {
	chanInt := generateChannel()

	// do some process
	time.Sleep(2 * time.Second)

	// cara kita ambil data dari channel
	chanData := <-chanInt
	fmt.Println("data dari channel", chanData)
}

func generateChannel() chan int {
	// design pattern: generator
	chanInt := make(chan int)

	// proceed calculation
	go func() {
		// defer adalah
		// built in function
		// yang memastikan golang
		// untuk menjalankan process yang didefer
		// di paling akhir
		defer close(chanInt)
		time.Sleep(time.Second)
		// cara kita assign data ke channel
		chanInt <- 1000
	}()

	return chanInt
}

func channelWithBuffer() {
	// goroutine with channel
	ts := time.Now()
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := []int{}

	// define channel
	// unbuffered -> kalau tidak hati hati
	// bisa berakibat deadlock!!!
	// paling aman kita kasih buffer di channel kita
	chanInt := make(chan int, 9)

	var wg sync.WaitGroup
	for _, val := range arr {
		wg.Add(1)
		go longMultiplyTfDataChan(&wg, chanInt, val)
	}
	wg.Wait()
	fmt.Println("with concurrent", time.Since(ts))

	for i := 0; i < 9; i++ {
		res = append(res, <-chanInt)
	}
	fmt.Println("data", res)
}

func dataRaceProblem() {
	// goroutine with channel
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := &[]int{}
	ts := time.Now()
	var wg sync.WaitGroup
	for _, val := range arr {
		wg.Add(1)
		go longMultiplyTfData(&wg, res, val)
	}
	wg.Wait()
	fmt.Println("with concurrent", time.Since(ts))
	// akan adanya DATA RACE!!!!!
	// sehingga akan ada unconsistent data
	fmt.Println("data", *res)
}

func longMultiplyTfDataChan(wg *sync.WaitGroup, result chan int, x int) error {
	time.Sleep(time.Second)
	fmt.Println("inside long multiply:", x*100)
	// transfer data
	result <- x * 100

	wg.Done()
	return nil
}

func longMultiplyTfData(wg *sync.WaitGroup, result *[]int, x int) error {
	time.Sleep(time.Second)
	fmt.Println("inside long multiply:", x*100)
	// transfer data
	temp := *result
	temp = append(temp, x*100)
	*result = temp

	wg.Done()
	return nil
}

func errorGroupProcess() {
	// ERROR GROUP
	//https://levelup.gitconnected.com/coordinating-goroutines-errgroup-c78bb5d80232
	// hampir sama dengan wait group
	// bedanya dia akan cancel semua goroutine
	// ketika ada 1 process yang error

	// 1. kita bisa cancel semua goroutine process (context)
	// 2. kita bisa menangkap error yang terjadi di process concurrent

	errGroup := eg.Group{}
	errGroup.Go(func() error {
		time.Sleep(time.Second)
		fmt.Println("first process")
		return errors.New("some error")
	})
	errGroup.Go(func() error {
		time.Sleep(2 * time.Second)
		fmt.Println("second process")
		return nil
	})
	errGroup.Wait()
}

func longMultiply(x int) {
	time.Sleep(time.Second)
	fmt.Println("inside long multiply:", x*100)
}

func longMultiplyWG(wg *sync.WaitGroup, x int) error {
	time.Sleep(time.Second)
	fmt.Println("inside long multiply:", x*100)
	// process goroutine
	// memberitahu bahwa
	// processnya sudah selesai
	wg.Done()
	return nil
}

func waitGroupProcess() {
	// WAIT GROUP
	ts := time.Now()

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, val := range arr {
		longMultiply(val)
	}
	fmt.Println("without concurrent", time.Since(ts))

	ts = time.Now()
	var wg sync.WaitGroup
	for _, val := range arr {
		// mendaftarkan goroutine
		// ke wait group
		wg.Add(1)
		// process selanjutnya
		// go routine bertanggung jawab
		// untuk memberitahu wg
		// bahwa processnya sudah selesai
		go longMultiplyWG(&wg, val)
	}
	// menunggu sampai
	// semua goroutine yang didaftarkan
	// sudah done
	wg.Wait()
	fmt.Println("with concurrent", time.Since(ts))
}

func asyncProcess() {
	// procedural / sync / sequential
	ts := time.Now()
	veryLongFunc()
	time.Sleep(2 * time.Second)
	fmt.Println("main process executed:", time.Since(ts))

	// async
	go veryLongFunc()
	time.Sleep(2 * time.Second)
	fmt.Println("main process executed:", time.Since(ts))
}

func veryLongFunc() {
	fmt.Println("very long func")
	time.Sleep(1 * time.Second)
}
