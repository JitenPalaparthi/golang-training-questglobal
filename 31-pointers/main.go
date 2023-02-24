// int and *int --> are different
package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var i, j int = 100, 200
	var ptr1, ptr2 = &i, &j
	s1 := sum1(i, j)
	fmt.Println("Sum1:", s1)

	ptr3 := sum2(ptr1, ptr2)
	fmt.Println("Sum2:", *ptr3)

	ptr4 := sum2(&i, &j)
	fmt.Println("Sum2:", *ptr4)

	s2 := sum3(i, j)
	fmt.Println("Sum3:", s2)
	//var ptr5 *int
	//*ptr5 = 100 // it has not hold any address..You cannot assign a value to a nil pointer
	ptr5 := sum2(ptr1, ptr2)
	fmt.Println("Sum4:", *ptr5)

	ptr6 := sum2(&i, &j)
	fmt.Println("Sum4:", *ptr6)

	slice := []int{10, 20, 30, 40}

	//pointer arthemetic

	ptr7 := &slice[0]

	fmt.Println(ptr7) // 0x14000022080 +8--> 0x14000022088
	// ptr7 = ptr7 + 8
	// ptr7 = ptr7 + 8

	//var addr uintptr = 0x14000022080 // convert pointer to uintptr
	var addr uintptr = uintptr(unsafe.Pointer(ptr7))

	addr = addr + 8
	//fmt.Println(*(*int)(addr))
	fmt.Printf("0x%X", addr)
	ptr7 = (*int)(unsafe.Pointer(addr))

	fmt.Println("\n", *ptr7)

}

func sum1(i, j int) int {
	s := i + j
	return s
}

func sum2(i, j *int) *int { // this func is expecting two pointers
	var s int
	s = *i + *j
	return &s //memory leak
}

func sum3(i, j int) int {
	var ptr *int
	var s int = i + j
	ptr = &s
	return *ptr
}

func sum4(i, j *int) *int {
	s := new(int) // memory is allocated and address is assigned to s which is a pointer
	// type inference work for the value in the memory
	*s = *i + *j
	return s //memory leak
}

//uintptr
