package main

import "fmt"

func main() {

	r1 := new(Rect) // pointer
	r1.L = 10.23
	r1.B = 12.34

	s1 := &Square{S: 25.25} //pointer

	c1 := Cuboid{L: 10.23, B: 12.34, H: 12.56} // not a pointer

	var ish1 IShape = r1 // only methods that r1 implements for IShape are returned
	a1 := ish1.Area()
	p1 := ish1.Perimeter()
	fmt.Println("area if r1:", a1)
	fmt.Println("perimeter of r1", p1)

	Shape(r1)
	Shape(s1)
	//Shape(c1) // interface is always a reference
	Shape(&c1)
}

func Area(iarea IArea) {
	fmt.Println("Area:", iarea.Area())
}

func Perimeter(iperimeter IPerimeter) {
	fmt.Println("Perimeter:", iperimeter.Perimeter())
}

func Shape(ishape IShape) {
	Area(ishape)
	Perimeter(ishape)
}

type Rect struct {
	L, B float32
}

type Square struct {
	S float32
}

type Cuboid struct {
	L, B, H float32
}

func (r *Rect) Area() float32 {
	return r.L * r.B
}

func (r *Rect) Perimeter() float32 {
	return 2 * (r.L + r.B)
}

func (r *Rect) PrintDetails() {
	fmt.Println("Length=", r.L)
	fmt.Println("Bredth=", r.B)
}

func (s *Square) Area() float32 {
	return s.S * s.S
}

func (s *Square) Perimeter() float32 {
	return 4 * s.S
}

func (s *Square) GetDetails() float32 {
	return s.S
}

func (c *Cuboid) Area() float32 {
	return c.L * c.B * c.H
}

func (c *Cuboid) Perimeter() float32 {
	return 2 * (c.L + c.B + c.H)
}

func (c *Cuboid) GetDetails() (float32, float32, float32) {
	return c.L, c.B, c.H
}

type IShape interface {
	IArea
	IPerimeter
}
type IArea interface {
	Area() float32
}
type IPerimeter interface {
	Perimeter() float32
}
