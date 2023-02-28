package main

import "fmt"

// buffered channels dont block until the buffer is full
func main() {
	ch := make(chan int, 3) // buffered channel
	ch <- 100
	ch <- 200
	ch <- 300
	//ch <- 400
	fmt.Println(<-ch, <-ch, <-ch)
	ch <- 101
	ch <- 202
	ch <- 303
	fmt.Println(<-ch, <-ch, <-ch)
	ch <- 1002
	fmt.Println(<-ch)
}
