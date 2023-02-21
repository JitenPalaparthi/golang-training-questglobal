package main

import "fmt"

func main() {
	age := 18
	// > < >= <= != == && || & |
	// if age >= 18 {
	// 	fmt.Println("eligible for vote")
	// } else {
	// 	fmt.Println("not eligible for vote")
	// }

	// if age := 21; age >= 21 {
	// 	fmt.Println("eligible for vote;because age is", age)
	// } else {
	// 	fmt.Println("not eligible for vote;because age is", age)
	// }

	if age >= 18 {
		fmt.Println("eligible for vote")
		return
	}
	fmt.Println("not eligible for vote")
}

// what is an expressions and what is a statement?
// what is return statemrent
