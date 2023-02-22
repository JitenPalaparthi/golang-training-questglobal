package main

import "fmt"

func main() {

	//for(int i=0;)

	// for loop
	// for i := 1; i <= 10; i++ { //untile the condition is true
	// 	fmt.Println(i)
	// }

	//str := "Hello, 世界"
	str := "Hello World"
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i])) // rune or int32
	}

	i := 1
	// everlasting loop but can break it using a condition inside it
	for {
		fmt.Println(i)
		if i == 100 {
			break
		}
		i++
	}

	i = 100

	for i <= 200 { // kind of a while loop
		fmt.Println(i)
		i++
	}

	for ; i <= 300; i++ { // no init but condition and incrementer
		fmt.Println(i)
	}

	// multiple init

	for i, j := 10, 1; j <= 10 && i <= 19; i, j = i+1, j+1 {
		fmt.Println(i, j, i*j)
	}

}

// for
// goto
