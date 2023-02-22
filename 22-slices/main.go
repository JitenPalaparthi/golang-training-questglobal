// slice in golang is a dynamic array.
// slice can grow at runtime where as array cannot grow at runtime.
// slice is mutable(The size of the slice)
// nil can be used to check whether slice is nil or not
// to instantiate a slice , need to use make built in function
// type inference work for slices
// can assisgn values to the slice as the same way as arrays
// to add a new element(one or more) to the slice , can use a built in function called append()

package main

import "fmt"

func main() {

	var slice1 []int // slice is declared but not instantiated
	fmt.Println(cap(slice1))

	// nil is used to check wheter a varialbe is nil or not.. usually all pointers, slices, maps, chans,interfaces we can check nil

	if slice1 == nil {
		fmt.Println("slice1 is nil")
	}

	slice1 = make([]int, 5) // allocating memory

	fmt.Println(slice1)
	slice1[0] = 10
	slice1[1] = 11

	fmt.Println(slice1)

	slice2 := []int{10, 11, 12, 13, 14} // shorthand declarion
	fmt.Println(slice2, "length=", len(slice2), "capacity", cap(slice2))
	fmt.Println("Before append", &slice2[0])
	//slice2[5] = 15
	slice2 = append(slice2, 15, 16, 17)
	//fmt.Println(&slice2[0])
	fmt.Println(slice2, "length=", len(slice2), "capacity", cap(slice2))
	fmt.Println("After append", &slice2[0])
	slice2 = append(slice2, 18, 19)
	fmt.Println(slice2, "length=", len(slice2), "capacity", cap(slice2))
	fmt.Println("After 2nd append", &slice2[0])

	slice2 = append(slice2, 20, 21)
	fmt.Println(slice2, "length=", len(slice2), "capacity", cap(slice2))
	fmt.Println("After 3nd append", &slice2[0])
	// how to add a new element to the slice

	for i, v := range slice2 {
		fmt.Println("index", i, "Value", v)
	}

	fmt.Println("Sum of slice1", sumOfSlice(slice1))
	fmt.Println("Sum of slice2", sumOfSlice(slice2))

	arr1 := [3]int{10, 11, 12}
	fmt.Println("Sum of arr1", sumOfSlice(arr1[:])) // can convert an array to a slice

	fmt.Println("All values of slice2:", slice2[:])
	fmt.Println("0-5 of slice2:", slice2[:5])   // begin to 5th
	fmt.Println("5-end of slice2:", slice2[5:]) // 5 to the end
	fmt.Println("3-8 of slice2:", slice2[3:8])  // 3rd to 8th element

	//[10 11 0 0 0]
	// StringBuilder sb1 = new StringBuilder("hello") // len: 5 cap: 10
	// sb1.Add("World")
	// sb1.Add("How are you doing")
}

func sumOfSlice(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
