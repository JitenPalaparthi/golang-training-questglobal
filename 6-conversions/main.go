package main

import "fmt"

// a special kind of variable which is called emtpy interface interface{}

//type integer = int

func main() {
	//var i1 interface{}
	var i1 any

	//fmt.Println(10, true, "Hai", 12312.31321, "hey"[0])

	i1 = 38 // int
	fmt.Println(i1)

	// var num1 int = int(i1) //type casting cannot be done on empty interface{} or any
	// instead you need to type assertion
	var num1 int = i1.(int) // any.(T)

	// not okay : var num2 int64 = i1.(int64)
	var num2 int64 = int64(i1.(int))
	fmt.Println(num2)
	fmt.Println(num1)

	i1 = "Quest Global" //string
	fmt.Println(i1)

	i1 = true //bool
	fmt.Println(i1)

	var b1 bool = i1.(bool)
	fmt.Println(b1)
	i1 = 143.456 //float
	fmt.Println(i1)
	//var i2 integer

}

// Object obj = 10
// obj="hey"
// obj=true
// obj= new Person()
