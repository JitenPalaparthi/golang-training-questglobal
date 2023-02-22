package main

import "fmt"

func main() {
	c := 1
	// for i := 0; i < 4; i++ { //outer loop
	// 	for j := 1; ; /*j<=4*/ j++ { // inner loop
	// 		if j > 4 {
	// 			break
	// 		}
	// 		fmt.Println("count=", c, "i=", i, "j=", j)
	// 		c++
	// 	}
	// }

	c = 1
outer:
	for i := 0; i < 4; i++ { //outer loop
		for j := 1; j <= 4; j++ { // inner loop
			if j == i {
				break outer
			}
			fmt.Println("count=", c, "i=", i, "j=", j)
			c++
		}
	}
}
