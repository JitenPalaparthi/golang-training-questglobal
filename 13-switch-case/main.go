package main

import "fmt"

func main() {
	// disclaimer : This works only for upper case vovels and consonents..
	// any lower case it is tread as another char..
	char := 'A'
	// empty switch
	switch {
	case char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U':
		fmt.Println(string(char), "is a vovel")
	case (char >= 66 && char <= 68) || (char >= 70 && char <= 72) || (char >= 74 && char <= 78) || (char >= 80 && char <= 84) || (char >= 86 && char <= 90):
		fmt.Println(string(char), "is a consonent")
	default:
		fmt.Println(string(char), "lowecase vovel or  a consonent or a special char or a digit")
	}
}
