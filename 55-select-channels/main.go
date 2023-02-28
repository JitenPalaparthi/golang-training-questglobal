package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := getServer("Server-1")
	ch2 := getServer("Server-2")
	ch3 := getServer("Server-3")
	ch4 := time.After(time.Millisecond * 1)
	select {
	case s1 := <-ch1:
		fmt.Println("Connecting to", s1)
	case s2 := <-ch2:
		fmt.Println("Connecting to", s2)
	case s3 := <-ch3:
		fmt.Println("Connecting to", s3)
	case <-ch4:
		fmt.Println("timed out")
		// default:
		// 	fmt.Println("no servers found")
	}
}

func getServer(str string) <-chan string {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Millisecond * 1)
		ch <- str
	}()
	return ch
}
