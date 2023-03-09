package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func(*sync.WaitGroup) {
		for i := 1; i <= 10; i++ {
			ping(ch1, ch2)
		}
		wg.Done()
	}(wg)

	go func(*sync.WaitGroup) {
		for i := 1; i <= 10; i++ {
			pong(ch1, ch2)
		}
		wg.Done()
	}(wg)

	wg.Wait()

	// go func() {
	// 	for i := 1; i <= 10; i++ {
	// 		ping(ch1, ch2)
	// 	}

	// }()

	// go func() {
	// 	for i := 1; i <= 10; i++ {
	// 		pong(ch1, ch2)
	// 	}

	// }()

	//runtime.Goexit()
}

func ping(ch1 chan<- string, ch2 <-chan string) {
	ch1 <- "ping"
	fmt.Println(<-ch2)
}

func pong(ch1 <-chan string, ch2 chan<- string) {
	fmt.Println(<-ch1)
	ch2 <- "pong"
}

// How does two or more threads communicate?
// What is the meaning of communication between threads?

// debit
// credit
// what is the shared data? account_balance --> data race here.
// go routines
// go routines comminicate to each other by channels

// One go routine sends a message called ping
// as soon the ping message is received the other goroutine has to send another message called pong
