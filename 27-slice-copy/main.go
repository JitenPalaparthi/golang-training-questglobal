package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var slice1 []int

	slice1 = make([]int, 10)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = rand.Intn(100)
	}
	slice2 := slice1 // both reffered to same memory // shallow copy

	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)

	slice2[0] = 2000
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)

	slice3 := make([]int, len(slice1))

	// for deep copy of slice there is a builtin function called copy

	copy(slice3, slice1) // deep copy

	// for i, v := range slice1 {
	// 	slice3[i] = v
	// }

	slice3[0] = 4000
	// i := 100
	// j := i //deep copy
	// j = 300

	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
	fmt.Println("slice3", slice3)

}
