// channels can be bidirectional or unidirectional
// chan <-int is sendonly channel
// <- chan is receive only channel
// for better readability
// for encapsulation
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := make(chan int) // bi-directional channel
	done := make(chan *struct{})
	//done := make(chan bool)
	go func() {
		for i := 1; i <= 100; i++ {
			// create a random number
			rnum := rand.Intn(100)
			go sender(ch, rnum)
		}
	}()
	go func() {
		for i := 1; i <= 100; i++ {
			receiver(ch)
		}
		//done <- true
		done <- new(struct{}) // if struct is not a pointer struct{}{}
	}()
	//time.Sleep(time.Second * 2)
	<-done
	// type employee struct{}
	// e1 := employee{}
}

//type employee struct{}

func sender(ch chan<- int, value int) { // send only channel
	ch <- value
}

func receiver(ch <-chan int) { // receive only channel
	fmt.Print(<-ch, " ")
}
