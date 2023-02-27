package main

import "fmt"

func main() {

	// print 1-100 even numbers
	for i := 1; i <= 100; i++ {
		if i%2 == 1 {
			continue // the execition does not go the next line.. it goes to the next iteration
		}
		fmt.Print(i, " ")
	}

	// print values those are divisible by 3 ..1-100
	fmt.Println()
	for i := 1; i <= 100; i++ {
		if i%3 != 0 {
			continue
		}
		fmt.Print(i, " ")
	}

	// differnece between break and continue is
	// break breaks the whole loop and terminates the loop
	// continue skips that particular iteration and continue with the next iteration
}
