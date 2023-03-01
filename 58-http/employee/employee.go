package employee

import "fmt"

type Employee struct {
	Id    int    `json:"id"`    // id
	Name  string `json:"name"`  // name
	Email string `json:"email"` // email
}

// type Person struct {
// 	Name, Email string
// }

// It works similar to constructor but not exactly a constructor
func New(id int, name, email string) *Employee {
	e := new(Employee)
	e.Id = id
	e.Name = name
	e.Email = email
	return e
}

// func NewPerson(name, email string) *Person {

// 	return nil
// }

// init is called automatically.
// there can be any number of init functions
func init() {
	fmt.Println("Employee Init is called-1")
}

func init() {
	fmt.Println("Employee Init is called-2")
}

func init() {
	fmt.Println("Employee Init is called-3")
}
