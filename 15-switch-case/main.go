package main

import "fmt"

func main() {
	val := 88
	switch {
	case val%8 == 0:
		fmt.Println(val, "is divisible by 8")
		fallthrough
	case val%4 == 0:
		fmt.Println(val, "is divisible by 4")
		fallthrough
	case val%2 == 0:
		fmt.Println(val, "is divisible by 2")
	case val%5 == 0:
		fmt.Println(val, "is divisible by 5")
	default:
		fmt.Println(val, "is divisible by some other numbers")
	}
}

// break is must in other programming languages to break the case flow.
// in Go break is not required. It executes only specific case
