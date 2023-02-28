package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})
	go func() {
		fmt.Println("Hello World;Receiving value")
		fmt.Println(<-ch)
		done <- struct{}{}
	}()

	go func() {
		fmt.Println("sending value--100")
		//time.Sleep(time.Second * 5)
		ch <- 100
	}()
	<-done
}
