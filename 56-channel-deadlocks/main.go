package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	//ch1 := make(chan int)
	done := make(chan struct{})
	go func() {
		for {
			time.Sleep(time.Second * 1)
			<-done
			break
		}
		ch <- 100
	}()
	//go func() {
	go func() {
		count := 1
		for {
			time.Sleep(time.Second * 1)
			count++
			fmt.Println("Waiting here...--->", count)
			if count == 11 {
				done <- struct{}{}
			}
		}
		//ch1 <- 100
	}()
	fmt.Println(<-ch)
	//fmt.Println(<-ch1)
	//}()
}

// if there is a receiver but not sender and there is a block it leads to deadlock
//
