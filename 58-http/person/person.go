package person

import "fmt"

type Employee struct {
	Id    int
	Name  string
	Email string
}

// init is called automatically.
// there can be any number of init functions
func init() {
	fmt.Println("Person Init is called-1")
}

func init() {
	fmt.Println("Person Init is called-2")
}

func init() {
	fmt.Println("Person Init is called-3")
}
