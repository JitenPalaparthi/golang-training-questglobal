package rect

// type rect struct { // This is restricted rect
// 	l, b float32 //lowercase are restricted fields
// }

// methods are associated with type(s)

type Rectangle struct {
	L, B float32
	A, P float32
}

// type Cube struct {
// 	L, B, H float32
// }

func Area(r Rectangle) float32 {
	return r.L * r.B
}

func Perimeter(r Rectangle) float32 {
	return 2 * (r.L + r.B)
}

// method: func (r receiver) function(param,param) returntype
func (rct *Rectangle) Area() float32 {
	rct.A = rct.L * rct.B
	return rct.A
}

func (rct *Rectangle) Perimeter() float32 {
	rct.P = 2 * (rct.L + rct.B)
	return rct.P
}

// func (c Cube) Area() float32 {
// 	return c.L * c.B * c.H
// }

// anything that stats with lowercase is restricted field..
// anything that statrs with uppercase is unrestricted field
