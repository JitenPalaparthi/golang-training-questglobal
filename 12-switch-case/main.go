package main

import "fmt"

func main() {
	switch char := 'A'; char {
	case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		fmt.Println(string(char), "is a vovel")
	case 'B', 'b', 'C', 'c', 'D', 'd', 'F', 'f', 'G', 'g', 'H', 'h', 'J', 'j', 'K', 'k', 'L', 'l', 'M', 'm', 'N', 'n', 'P', 'p', 'Q', 'q', 'R', 'r', 'S', 's', 'T', 't', 'V', 'v', 'W', 'w', 'X', 'x', 'Y', 'y', 'Z', 'z':
		fmt.Println(string(char), "is a consonent")
	default:
		fmt.Println(string(char), "is a special character")
	}
}
