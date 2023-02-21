package main

import "fmt"

// conversions or type casting in go
// a type cannot be automatically converted to another type
// to do conversion, you need to do type casting or type assertion

func main() {
	var age uint8 = 38 // 1 byte

	// error: var num1 uint64 = age // This is possible in other programming languages

	// age 10011000
	// num1 00000000 00000000 00000000 00000000 0000000 00000000 00000000 00000000

	var num1 uint64 = uint64(age) // (uint64)age
	var num2 int64 = int64(num1)  // even uint64 to int64 need to type caste

	flt1 := 12.345 // float32, float64
	var flt2 float32 = 12.456
	var num3 int = int(flt1)
	fmt.Println(age, num1, num2, num3)

	fmt.Println(uint8(flt1) % uint8(flt2))
	// arthemetic operations :- both the operands must be of same type and also the result of that operation must be of same type

	a := 10   // int
	b := 10.1 //float64
	c := float64(a) + b
	fmt.Println(c)

	// num4 00000000
	// num5 00000000 10111011 11011110 11111000

	var num4 uint16

	var num5 uint32 = 12312312

	num4 = uint16(num5)

	fmt.Println(num4)
}
