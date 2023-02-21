package main

import "fmt"

func main() {
	//var v1 = 4
	// constant: this should be evaluated at compiletime
	//v1 = 8
	const (
		pi        = 3.14          // evaluated at comp time
		max_procs = 8             // evaluated at comp time
		eval1     = max_procs * 8 // evaluated at comp time;bcz maxc_procs is constant and 8 is also a constant
		//eval2     = max_procs * v1 // Not ok. This is evaluated at runtime
		isOkay = true || false // evaluated at comp time
	)

	const DAY1 string = "Sunday"
	fmt.Printf("pi=%.2f\n", pi) // to format and print only 2 decimals
	println(pi, max_procs, eval1, isOkay, DAY1)

}
