package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := new(sync.WaitGroup)

	//wg.Add(5) //add 6 to the counter there are 5 goroutines .need to wait for 5 goroutines to complete

	go evenNumbers(wg, 100)

	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		isPrime(109)
		wg.Done() // decrement the counter
	}(wg)

	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		isPrime(47)
		wg.Done() // decrement the counter
	}(wg)

	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		isPrime(67)
		wg.Done() // decrement the counter
	}(wg)

	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		isPrime(58)
		wg.Done() // decrement the counter
	}(wg)

	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		isPrime(45)
		wg.Done() // decrement the counter
	}(wg)

	wg.Wait() // wait untile the counter is zero

	fmt.Println("End of main")
}

func isPrime(num int) {
	if num == 0 {
		fmt.Println(num, "is not a prime number")
		return
	} else if num == 1 || num == 2 || num == 3 {
		fmt.Println(num, "is a prime number")
		return
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			fmt.Println(num, " not a prime number")
			return
		}
	}
	fmt.Println(num, "is a prime number")
}

func evenNumbers(wg *sync.WaitGroup, num int) {
	wg.Add(1)
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			fmt.Println(i, " is even number")
		}
	}
	wg.Done()
}
