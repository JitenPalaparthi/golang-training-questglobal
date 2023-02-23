package main

import "fmt"

func main() {

	i := 1
loop: //label
	fmt.Println(i)
	i++
	if i <= 10 {
		goto loop
	} else {
		goto done
	}
done:
	fmt.Println("Goto done")
	// i = 1
	// goto loop
}
