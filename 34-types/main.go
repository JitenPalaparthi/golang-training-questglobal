package main

import (
	"fmt"
	"udt/shape/rect"
)

func main() {

	//r1 := &rect.Rectangle{L: 10.45, B: 15.67}
	r1 := new(rect.Rectangle) // this is a pointer
	r1.L = 10.45
	r1.B = 15.67
	a1 := rect.Area(*r1)
	p1 := rect.Perimeter(*r1)

	fmt.Println("Area of Rect:", a1)
	fmt.Println("perimeter of Rect", p1)

	a2 := r1.Area()      // accessing the method that is associated with a Rect
	p2 := r1.Perimeter() // ""

	fmt.Println("Area of Rect:", a2)
	fmt.Println("perimeter of Rect", p2)

	fmt.Println("Area", r1.A)

	// r2 := rect.Rectangle{L: 10.34, B: 15.65}
	// var r2ptr *rect.Rectangle
	// r2ptr = &r2
}
