package main

import "fmt"

func main() {

	c1 := arithmetic(10, 20, func(a1, b1 int) int {
		return a1 + b1
	})
	fmt.Println("c1:", c1)

	c2 := arithmetic(10, 20, sub)
	fmt.Println("c2:", c2)

	f3 := func(a, b int) int {
		return a * b
	}

	c3 := arithmetic(10, 20, f3)
	fmt.Println("c3:", c3)

	var funcMap map[string]func(int, int) int

	funcMap = make(map[string]func(int, int) int)

	funcMap["add"] = func(a, b int) int {
		return a + b
	}

	funcMap["sub"] = sub

	funcMap["multiply"] = f3

	a1 := funcMap["add"]

	fmt.Println("addition=", a1(10, 20))

	a2 := funcMap["sub"]

	fmt.Println("substractions=", a2(10, 20))

	a3 := funcMap["multiply"]

	fmt.Println("multiplication=", a3(10, 20))

}

func arithmetic(a, b int, fn func(int, int) int) int {
	c := fn(a, b)
	return c
}

func sub(a, b int) int {
	return a - b
}
