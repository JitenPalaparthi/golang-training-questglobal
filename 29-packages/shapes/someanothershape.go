package shapes

import "fmt"

// this func cannot be accessed outside of the package
func getPackageName() {
	fmt.Println("shapes")
}
