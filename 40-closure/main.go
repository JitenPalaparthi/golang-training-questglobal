package main

import "fmt"

func main() {

	c1 := arithmetic(10, 20, func(a, b int) int {
		return a + b
	})
	fmt.Println("c1:", c1)

	c2 := arithmetic(10, 20, sub)
	fmt.Println("c2:", c2)

	f3 := func(a, b int) int {
		return a * b
	}

	c3 := arithmetic(10, 20, f3)
	fmt.Println("c3:", c3)

}

func arithmetic(a, b int, fn func(int, int) int) int {
	c := fn(a, b)
	return c
}

func sub(a, b int) int {
	return a - b
}
