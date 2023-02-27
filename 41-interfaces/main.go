package main

import "fmt"

type Stringer interface {
	ToString() string // definition
}

func main() {
	var mi myInt = 100

	mm := make(myMap)
	mm["company"] = "Quest Global"
	mm["address"] = "Bangalore"

	e1 := &Employee{Id: 101, Name: "Jiten", Email: "Jitenp@outlook.com"}

	//r1 := &Rect{L: 10.45, B: 12.56}

	slice := []any{mi, mm, e1}

	for _, v := range slice {
		var istr Stringer
		istr = v.(Stringer) // if it satisfies then only you can convert to interface
		fmt.Println(istr.ToString())
	}

	// if the type implements a function/ method that has the same signature of interface
	// it automatically satisfies that interface
}

type myInt int //1- User defined int

func (mi myInt) ToString() string {
	return fmt.Sprintln(mi)
}

type myMap map[string]any //2- User defined map

func (mm myMap) ToString() string {
	str := ""
	for key, value := range mm {
		str = str + fmt.Sprintf("key:%v value:%v,", key, value) //str + "key:" + key + " value:" + value + ","
	}
	return str[0 : len(str)-1]
}

type Employee struct { // 3- User defined struct
	Id    int
	Name  string
	Email string
}

func (e *Employee) ToString() string {
	return fmt.Sprintf("Id:%d,Name:%s,Email:%s", e.Id, e.Name, e.Email)
}

type Rect struct { // 4- User defined struct
	L, B float32
}

func (r *Rect) Area() float32 {
	return r.L * r.B
}
