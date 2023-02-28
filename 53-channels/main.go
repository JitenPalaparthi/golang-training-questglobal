package main

import "fmt"

func main() {
	ch1 := GetSquare(20)
	ch2 := GetSquare(40)

	// for {
	// 	val, ok := <-ch
	// 	if !ok { // false means channel is closed
	// 		break
	// 	} else {
	// 		fmt.Println(val)
	// 	}
	// }
	go func() {
		for v := range ch1 { // tries to retrive until the channel is closed
			fmt.Println("chan-1", v)
		}
	}()
	for v := range ch2 { // tries to retrive until the channel is closed
		fmt.Println("chan-2", v)
	}
}

func GetSquare(num int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= num; i++ {
			ch <- i * i
		}
		close(ch) //1. closing a channel .2 only a sender can close a channel
	}()
	return ch
}
