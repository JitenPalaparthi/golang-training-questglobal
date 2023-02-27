package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("hello world from go routine")
	// never use global variables inside a goroutine
	go func() {
		for i := 1; i <= 100; i++ {
			go fmt.Println(i)
		}
	}()

	time.Sleep(time.Second * 2) // this is not a solution to wait goroutine. This is a work around
}
