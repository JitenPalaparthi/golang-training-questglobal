package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1- lenght of an array must be evaluated at compile time
	// 2- size of an array is immutable. if initially length of an array is 3 then it remains 3
	// 3- type inference works for an array
	// 4- type of an array contains its length;
	//	  From below example arr1 and arr2 are different types
	//    type of arr1 [4]int
	//    type of arr2 [3]int

	var arr1 [4]int
	arr1[0] = 10
	arr1[1] = 11
	arr1[2] = 12
	arr1[3] = 13

	arr1[0] = 100 // mutability

	fmt.Println("arr1 type and value", reflect.TypeOf(arr1), arr1)
	arr2 := [3]int{10, 11, 12} // shorthand declaration of array with length
	fmt.Println("arr2", arr2)
	arr3 := [...]int{10, 11, 12} // shorthand declararion of array but the length is evaluated at compile time
	fmt.Println("arr3", arr3)
	var arr4 [3]int // [0,0,0]
	fmt.Println("arr4", arr4)
	//var arr5 [3]bool   // false,false,false
	//var arr6 [3]string // "","",""

	var arr7 [3]any
	arr7[0] = "Hello"
	arr7[1] = false
	arr7[2] = 100

	fmt.Println("arr7", arr7)

	// iterating array

	for i := 0; i < len(arr1); i++ {
		fmt.Println("arr1 Index:", i, "Value:", arr1[i])
	}

	sumOfArr(arr1) // [4]int []int

	var arr8 [4]int

	arr8 = arr1

	// arr9 := arr1 + arr8

	fmt.Println("arr8", arr8)

}

// func sumOfArr(arr []int) int {
// 	sum := 0
// 	for i := 0; i < len(arr); i++ {
// 		sum += arr[i]
// 	}
// 	return sum
// }

func sumOfArr(arr [4]int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
