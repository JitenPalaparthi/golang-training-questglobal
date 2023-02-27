package main

import "fmt"

// defer defers the execution towards the end of the caller
// defer maintains a stack
func main() {
	defer func() {
		defer fmt.Println("end of main-1")
		fmt.Println("end of main-2")
	}()
	fmt.Println("start of main")
	a := new(int)
	*a = 10
	defer printA(a) // defer maintains its stack if not call by reference/pointers.. even defer is called after a=20
	*a = 20

	str := "Hello World"
	fmt.Println()

	for _, v := range str {
		defer fmt.Print(string(v))
	}

}
func printA(a *int) {
	fmt.Println(*a)
}
