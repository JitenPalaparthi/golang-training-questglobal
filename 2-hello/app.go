package main

import (
	"fmt"
	_ "os"
	"time"
)

// import "fmt"
// import "time"
// import "os"

func main() {
	fmt.Println()
	fmt.Println("Hello World!", time.Now())
	fmt.Println("Hello", "World", "!", "How", "are", "you", "doing?")

	fmt.Print("Hello", " World\n") // report it as a bug in golang

	fmt.Printf("Time is %v", time.Now())
	fmt.Printf("\nDay:%d Month:%s Year:%d\n", time.Now().Day(), time.Now().Month(), time.Now().Year())
}
