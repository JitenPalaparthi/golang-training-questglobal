package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		fmt.Println("I am running inside a go routine")
		i := 1
		for {
			if i == 10 {
				runtime.Goexit() // return or break statement..
			}
			go fmt.Println("Hello World-->", i)
			i++
		}
	}()
	time.Sleep(time.Second * 1)
}
