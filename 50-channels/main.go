// Do not communicate by sharing memory; instead, share memory by communicating.
// share by communicating is done through channels in go (chan)
// chan is a reference type.. so use make function to instantiate a channel
// when a goroutine tries to communicate to another goroutine
//
//	1 - sender goroutine: is blocked until the receiver goroutine receives the value (provided by the size of the channel)
//	2 - receiver goroutine: is blocked until the sender sends the value
//	3 - if runtime understands that any one of them is unable to proceed, it reaches to a deadlock
//
// there are two types of channels. Buffered and unbuffered. Rules 1 and 2 are little differnet for buffered and unbuffered
// can check whether a channel is nil or not
// channel can be created with all valid types
// to send data to channel ch <- 100
// to receive a value from channel k:= <- ch
// can also receive a value and dont assign to any varialbe <-ch (similar to a func returns a value but you dont take that value still the func executes)
// channels can be bidirectional or unidirectional
package main

import (
	"fmt"
	"time"
)

func main() {
	// how to declare a channel
	var ch chan int // ch is a channel of type int
	// how to instantiate a channel
	ch = make(chan int) // This is a unbuffered channel
	go func() {
		ch <- 100 // sending
		//v = 100
	}()
	go receive(ch)
	//ch1 := make(chan int) // short hand declaration of channel
	//add(10, 20)
	time.Sleep(time.Millisecond * 1)
}

func receive(ch chan int) {
	val := <-ch //receiving
	fmt.Println(val)
}
