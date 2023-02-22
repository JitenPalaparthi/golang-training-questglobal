package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice1 := make([]int, 10)
	//slice2 := make([]int, 10)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = rand.Intn(100) // to randum numbers are given to the slice
	}
	fmt.Println("Slice1", slice1)
	slice2 := slice1 // both of them are pointed to the same pointer
	slice2[0] = 200
	fmt.Println("Slice1", slice1)
	fmt.Println("Slice2", slice2)

	slice3 := slice1[3:] // &slice1[3] both of them pointed to the same memory but from 3rd elements onwareds of slice1
	//slice3[0] = 300
	fmt.Println("Slice1", slice1)
	fmt.Println("Slice2", slice2)
	fmt.Println("Slice3", slice3)

	arr1 := [3]int{10, 11, 12}
	arr2 := [3]int{}
	slice4 := arr1[:] // Slice4 is pointed to the address of arr
	arr2 = arr1       // deep copy..
	slice4[0] = 100

	fmt.Println("Arr1", arr1)
	fmt.Println("Arr2", arr2)
	fmt.Println("Slice4", slice4)

}

// [200| 43| 54| 12|67 |55 |48 | 75| 22| 39|]
// 			[|12|67 |55 |48 | 75| 22| 39|]
