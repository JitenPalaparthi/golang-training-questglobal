package main

// a package name should be a valid identifier
// usually pacakge names start with lowercase
// the name of the directory is the name of the package
// a single package can be split into many files
// there are no access specifiers or modifiers in go.. // There is no public/private/protected etc..
// inside package any thing starts with UpperCase is public(Unrestricted). Anything starts with lowerCase is private(restricted)
// go mod init <name> is the root package name w.r.t the project
// "fmt"
// "math/rand"
// "os"

import (
	"fmt"
	"packagedemo/shapes"
	r "packagedemo/shapes/rect" // using alias name
	"packagedemo/shapes/square"
)

func main() {
	shapes.GetPackageName()
	// shapes.getPackageName() it is restricted func
	a1 := r.Area(10.34, 12.45)
	p1 := r.Perimeter(10.34, 12.45)
	fmt.Println("Area of Rect", a1, "perimeter of Rect", p1)

	a2 := square.Area(25.25)
	p2 := square.Perimeter(25.25)
	fmt.Println("Area of Square", a2, "perimeter of Square", p2)
}

// what ever the name of go mod init is given , that is the root package name
// why do we write packages? for modularity and reusability and for encapusaltion
