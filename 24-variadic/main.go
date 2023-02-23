package main

import "fmt"

// variadic parameter is used to pass any number of arguments. As long as there is a spece for function stack frame
// can also pass zero arguments
// eclipse symbol is used to define variadic paramter.
// variadic parameter must be the last parameter
// arrays[after converting as slices] and slices can be passed as variadic arguments.
func main() {
	//fmt.Println()
	// myPrint("Hello")
	// myPrint(100)
	// myPrint(true)
	//myPrint(100, "hello")
	//fmt.Println("Hello", 100, true, 100.34)
	s1 := sumOf()
	fmt.Println("s1=", s1)
	s2 := sumOf(10, 20)
	fmt.Println("s2=", s2)
	s3 := sumOf(10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
	fmt.Println("s3=", s3)

	arr1 := [4]int{10, 20, 30, 40}
	s4 := sumOf(arr1[:]...)
	fmt.Println("s4=", s4)

	slice1 := []int{10, 20, 30, 40, 50, 60, 70, 80}
	//s5 := sumOf(slice1) // expected argument is variadic argument but passes as slice
	s5 := sumOf(slice1...) // slice... converts from slice to variadic argument
	fmt.Println("s5=", s5)
}

func sumOf(nums ...int) int { // variadic parameter
	var sum int
	for _, v := range nums {
		sum = sum + v //sum += v
	}
	return sum
}

// func arthemeticOf(nums ...int,operation string,) int { Not Ok because nums must be last paramter
func arthemeticOf(operation string, nums ...int) int {
	// todo
	return 0
}

// normal function can pass only one argument
func myPrint(a any) {
	fmt.Println(a)
}
