package main

import "fmt"

func main() {

	arr1 := [...]int{10, 11, 12, 13, 14, 15}
	// for range loop
	// range loop can be used for arrays, slices, maps and channels
	// for arrays and slices --> range loop returns index and value
	for i, v := range arr1 {
		fmt.Println("Index", i, "Value", v)
	}

	str1 := "Hello, 世界"
	fmt.Println("Normal loop")
	for l := 0; l < len(str1); l++ {
		fmt.Println(l, string(str1[l]))
	}
	fmt.Println("Range loop")

	for l, val := range str1 { // if the loop is on string , it returns rune
		fmt.Println(l, string(val))
	}

}
