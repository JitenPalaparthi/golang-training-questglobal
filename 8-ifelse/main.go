package main

import "fmt"

func main() {

	age, gender := 18, 'f' //rune --> int32 used to store chars
	// M = 77 and F=70
	// m=109 and f=102
	//fmt.Println(reflect.TypeOf(age), age, reflect.TypeOf(gender), gender)
	if age >= 18 && (gender == 70 || gender == 'f') {
		fmt.Println("She is elibible for marriage")
	} else if age >= 21 && (gender == 'M' || gender == 109) {
		fmt.Println("He is eligible for marriage")
	} else {
		fmt.Println("wrong data;either age or gender is given wrong")
	}
}
