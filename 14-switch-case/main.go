package main

import "fmt"

func main() {
	var value any
	var f1 float32 = 12.34
	value = f1
	// fmt.Println("Enter a value")
	// fmt.Scanln(&value)
	switch value.(type) { // This is only for empty interface{} a.k.a any
	case int:
		fmt.Println("The value is of type int:", value)
	case uint8:
		fmt.Println("The value is of type uint8:", value)
	case uint16:
		fmt.Println("The value is of type uint16:", value)
	case uint32:
		fmt.Println("The value is of type uint32:", value)
	case uint64:
		fmt.Println("The value is of type uint64:", value)
	case float32:
		fmt.Println("The value is of type float32:", value)
	case float64:
		fmt.Println("The value is of type float64:", value)
	case string:
		fmt.Println("The value is of type string:", value)
	default:
		fmt.Println("This type is not used")
	}

	// val1, val2, val3, _, _ := 10, 10.5, "Hello", true, 'A'
	// fmt.Println(val1, val2, val3)
	// System.Object
	// Object obj = new Employee();
	// obj = 10
	// obj= 10.11
	// obj=false

	// 1- Generic Methods/Functions
	// 2- Dynamic Programming

}
