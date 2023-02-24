package main

import "fmt"

type Person struct {
	No        int     // 0
	Name      string  // ""
	Email     string  // ""
	Address   *string // nil
	IsMarried bool    // false
}

// built in types    --> int, float32 etc..
// user defined types --> type that user creates..
// type inference work even for user defined types

func main() {
	var p1 Person
	fmt.Println(p1) // prints values based on type inference
	p2 := Person{}  // short hand declaration
	fmt.Println(p2)
	str := new(string)
	*str = "Bangalore,India"
	p3 := Person{No: 101, Name: "Jiten", Email: "JitenP@outlook.com", Address: str, IsMarried: true}
	fmt.Println(p3)

	p1.No = 102

	p1.Name = "Rajan"
	p1.Email = "Rajan.Singh@gmail.Com"
	p1.Address = str
	p1.IsMarried = true
	fmt.Println("Name:", p1.Name)
	fmt.Println("No:", p1.No)
	fmt.Println("Email:", p1.Email)
	fmt.Println("Address:", *(p1).Address) // can directly give p1.Addresss
}
