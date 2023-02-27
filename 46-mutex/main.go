package main

import (
	"fmt"
	"sync"
)

var (
	num, numM int = 1000, 1000

	wg *sync.WaitGroup
	mu *sync.Mutex
)

func main() {
	wg = &sync.WaitGroup{}
	mu = &sync.Mutex{}

	wg.Add(404)
	go func(wg *sync.WaitGroup) {
		for i := 1; i <= 100; i++ {
			go increment()
		}
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		for i := 1; i <= 100; i++ {
			go decrement()
		}
		wg.Done()
	}(wg)
	// with mutex
	go func(wg *sync.WaitGroup) {
		for i := 1; i <= 100; i++ {
			go incrementM()
		}
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		for i := 1; i <= 100; i++ {
			go decrementM()
		}
		wg.Done()
	}(wg)

	wg.Wait()

	//time.Sleep(time.Second * 2)
	fmt.Println("num without mutex", num)
	fmt.Println("numM with mutex", numM)

}

func increment() {
	num++
	wg.Done()
}

func decrement() {
	num--
	wg.Done()
}

func incrementM() {
	mu.Lock()
	numM++
	mu.Unlock()
	wg.Done()
}

func decrementM() {
	mu.Lock()
	numM--
	mu.Unlock()
	wg.Done()
}
