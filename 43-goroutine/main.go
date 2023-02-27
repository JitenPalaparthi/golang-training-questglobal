package main

import (
	"fmt"
	"time"
)

// 1- main is also a go routine
// 2- to run any func/method as a goroutine just need to use a keyword called go infront of that func/method
// 3- no goroutine waits for other goroutine to complete its execution
// 4- we cannot guarantee the order of execution when goroutines are running

func main() {
	go Sayhi()
	fmt.Println("Hello World from main")

	go fmt.Println("Can I be executed as a goroutine?")

	go func() {
		for i := 1; i <= 100; i++ {
			fmt.Println("Go Routine-1", i)
		}
	}()

	go func() {
		for i := 1; i <= 100; i++ {
			fmt.Println("Go Routine-2", i)
		}
	}()
	fmt.Println("end of main")

	time.Sleep(time.Second * 1)
}

func Sayhi() {
	fmt.Println("Hello World from a function")
}
