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

	var (
		age1 = 38 // waste of 7 bytes..
		age2 uint8
	) // saves 7 bytes..

	fmt.Println("Value of age1 and type of age1", age1, reflect.TypeOf(age1))
	fmt.Println("Value of age2 and type of age2", age2, reflect.TypeOf(age2))

	num12 := 12 // variable shorthand declaration. There is no var, and there is no var type

	fmt.Println("value of num12=", num12, "type of num12", reflect.TypeOf(num12))

	// bool
	// string
	// rune := uint32 -> unicode chars
	// byte
	// complex

	// var b bool
	// var b = true
	b, c := true, false // multiple shorthand variable declaration

	// type inference; default value
	// bool : false
	// numberic : 0
	// string: ""
	// byte: 0
	// rune: 0
	// uintptr : large integer can store address

	fmt.Println(b && c)               // true and false
	fmt.Println(false && (true && b)) // true and false
	fmt.Println(b || c)               // true or false
	fmt.Println(false || c)           // true and false

	var name = "Quest Global" // 12 bytes
	fmt.Println(name)
	name = "Quest Global Private Limited"
	// string is immutable
	// string is immutable from your runtime perspective not from developers perspective

	greet := "Hello 世界世" // 15 bytes bsed on the character the size if taken
	fmt.Println(greet)
	fmt.Println(len(name), len(greet)) // len is a builtin function
	// a := 10
	// a = 20

	var char rune = 44444 // 궜
	var bt byte = 65
	fmt.Println(string(char)) // type casting ..
	fmt.Println(string(bt))   // type casting ..

	// complex type
	// complex64, complex128
	c1 := complex(12, 10.4567)
	// 12+10.4567i
	var r1, i1 float32 = 12, 10.4567
	c2 := complex(r1, i1)

	fmt.Println("Value and type of c1 ", c1, reflect.TypeOf(c1))
	fmt.Println("Value and type of c2 ", c2, reflect.TypeOf(c2))

	c3 := 12 + 12.34i // direct complex number declartion; Shortcut without builtin function

	c4 := c1 + c3
	fmt.Println("Value and type of c4 ", c4, reflect.TypeOf(c3))
	//c5 := c1 + c2 // error as mismatched types of complex128 and complex64

	fmt.Println("Complex Sub:", c1-c3)
	fmt.Println("Complex Mul:", c1*c3)
	fmt.Println("Complex Div:", c1/c3)

}
