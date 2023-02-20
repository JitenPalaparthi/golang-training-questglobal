package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num1 int8 // 1 byte
	num1 = 127
	fmt.Printf("num1=%d\n", num1)

	// type inference
	// long l; //example in clang
	// printf("%d",&l)

	//types
	// int8,int16,int32,int64,int
	// uint8,uint16,uint32,uint64
	// float32, float64

	var num2 int16 // type inference gives a default value 0
	fmt.Println("num2=", num2)

	var num3, num4, num5 int32 = 10000, 100001, 100002 // declaring multiple variables

	var (
		num6, num7 float32 = 10.11, 12.22 // declaring multiple variables as a group
		num8, num9 float64 = 10.121231231, 12.1231231231
	)

	//fmt.Println(num3, num4, num5, num6, num7, num8, num9)
	fmt.Printf("value and type of num3=%v %T\n", num3, num3)
	fmt.Printf("value and type of num4=%v %T\n", num4, num4)
	fmt.Printf("value and type of num5=%v %T\n", num5, num5)
	fmt.Printf("value and type of num6=%v %T\n", num6, num6)
	fmt.Printf("value and type of num7=%v %T\n", num7, num7)
	fmt.Printf("value and type of num8=%v %T\n", num8, num8)
	fmt.Printf("value and type of num9=%v %T\n", num9, num9)

	var num10 = 12.4 // declare by assigning value
	var num11 = 0

	fmt.Println("value of num10=", num10, "type of num10", reflect.TypeOf(num10))
	fmt.Println("value of num11=", num11, "type of num11", reflect.TypeOf(num11))

	var age1 = 38  // waste of 7 bytes..
	var age2 uint8 // saves 7 bytes..

	fmt.Println("Value of age1 and type of age1", age1, reflect.TypeOf(age1))
	fmt.Println("Value of age2 and type of age2", age2, reflect.TypeOf(age2))

}
