package main

import (
	"fmt"
	"shapes/shape/rect"
)

func main() {
	r1 := new(rect.Rect)
	r1.L = 10.10
	r1.B = 15.35
	fmt.Println(r1.Area(), r1.Perimeter())
}
