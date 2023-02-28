package main

import "fmt"

// panic is a built in function
func main() {
	fname := new(string)
	*fname = "Jiten"
	lname := new(string)
	*lname = "P"
	defer recoverMe()
	defer func() {
		fmt.Println("Seems there is a panic at line number 17... but I can be executed even when panics")
	}()
	func() {
		func() {
			defer recoverMe()
			defer fmt.Println("There is a panic still I can do this")
			defer fmt.Println(getFullName(fname, nil))
		}()
		fmt.Println(getFullName(fname, lname))
	}()

	num := 0
	fmt.Println(100 / num)
	//
}

func recoverMe() {
	if recover() != nil {
		// what ever the mitigation that is required during panic .. write here
		fmt.Println("this call is recovered")
	}
}

// deallock and panic.
func getFullName(firstname, lastname *string) string {
	if firstname == nil || lastname == nil {
		panic("firstname or lastname is nil")
	}
	str := *firstname + " " + *lastname
	return str
}
