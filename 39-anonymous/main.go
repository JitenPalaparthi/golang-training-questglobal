package main

import "fmt"

// anonymous function does not contain a name
// functions can be stored in variables
// functions can be passed as arguments same way as variables
// bcz functions can be stored in variables then they can be stored in any types..maps, slices etc.
// func(params){}

// func f4(num uint)bool{
// 	return false
// }

var f3 func(uint) bool // sinature of a function

func main() {

	//func() // parameters
	func() {
		fmt.Println("Hello World")
	}() // arguments

	var f1 func()
	f1 = func() {
		fmt.Println("Hello World")
	}

	f1()
	// anonymous function, takes arguments and returns

	c1 := func(a int, b int) int {
		return a + b
	}(10, 20)

	fmt.Println(c1)

	fmt.Println(func(a int, b int) int {
		return a + b
	}(10, 20)) // So function can be passed as an argument to another function

	f2 := func(a int, b int) int {
		return a + b
	} // not executed..
	c2 := f2(10, 20)
	fmt.Println(c2)

	c3 := f2(100, 200)
	fmt.Println(c3)

	// prime number..
	f3 = func(num uint) bool {
		if num == 0 {
			return false
		}
		if num == 1 || num == 2 || num == 3 {
			return true
		}
		for i := 2; i <= int(num)/2; i++ {
			if int(num)%i == 0 {
				return false
			}
		}
		return true
	}

	fmt.Println("Is Prime number?", f3(8))
}
