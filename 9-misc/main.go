package main

import "fmt"

func main() {

	a1, b1 := 10, 20
	fmt.Println("Before swap a1:", a1, "b1:", b1)
	t1 := a1
	a1 = b1
	b1 = t1
	fmt.Println("After swap a1:", a1, "b1:", b1)

	a2, b2 := 10, 20
	fmt.Println("Before swap a2:", a2, "b2:", b2)

	a2, b2 = b2, a2 // simple swap in go
	fmt.Println("After swap a2:", a2, "b2:", b2)

	a3, b3, c3 := 10, 20, 30
	a3, b3, c3 = b3, c3, a3
}
