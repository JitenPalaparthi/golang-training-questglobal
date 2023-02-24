package main

import "fmt"

func main() {

	p1 := Person{No: 101, Name: "jiten", Email: "Jitenp@Outlook.com", Addr: Address{Line1: "Door no 101", City: "Bangalore", State: "karnataka", Country: "India", ZipCode: "560086"}}
	p2 := Person{No: 101, Name: "jiten", Email: "Jitenp@Outlook.com", Addr: Address{Line1: "Door no 101", City: "Bangalore", State: "karnataka", Country: "India", ZipCode: "560086"}}

	if p1 == p2 {
		fmt.Println("Both structs are same")
	}

	fmt.Println(p1)           // The whole object
	fmt.Println(p1.Addr)      // object of Addr
	fmt.Println(p1.Addr.City) // Value of Object of Addr which is in p1 object

	e1 := Employee{}
	e1.No = 101
	e1.Name = "Jiten"
	e1.Email = "JitenP@Outlook.Com"
	e1.Line1 = "Door no 101"
	e1.Street = "Mahalakshmi Layout"
	e1.Country = "India"
	e1.State = "Karnataka"
	e1.City = "Bangalore"
	e1.ZipCode = "560086"
	e1.Status = "Active"
	e1.Address.Status = "Active"
	fmt.Println(e1)
	e1.GetAddress()

	// anonymous struct
	c1 := struct {
		Id   int
		Name string
	}{
		Name: "Quest Global",
		Id:   1001,
	}
	c2 := struct {
		Id   int
		Name string
	}{
		Name: "Quest Global",
		Id:   1001,
	}
	fmt.Println(c1 == c2) // two struct objects can be compared(equal only)..true if each filed value is equal to the other

}

// golang supprts has a relationship
type Person struct {
	No          int
	Name, Email string
	Addr        Address //composition : Person has a relationship with Addr..
}

type Address struct {
	Line1, City, Street, State, Country, ZipCode, Status string
}

func (a Address) GetAddress() {
	fmt.Println(a)
}

type Employee struct {
	No                  int
	Name, Email, Status string
	Address             // promoted field similar to kind of Inheritance
	// Address is related to person
	//Company
}

// class Address{
// 	string City,Country;
// }

// class Person:Address{
// 	string Name;
// 	int id
// }

// Person p1 = new Person();
// p1.Name
// p1.Id
// p1.City
// p1.County

// type DBFields struct {
// 	ID           int
// 	Status       string
// 	LastModified uint64
// }

// type Employee struct {
// 	DBFields
// 	Name  string
// 	EmpID int
// 	Sal   float32
// }

// type Item struct {
// 	DBFields
// 	Name     string
// 	Quantity int
// }
