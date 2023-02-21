package main

import (
	"fmt"
	"reflect"
	"strconv"
	_ "time"
)

func main() {
	s1 := "ABCD"
	println(len(s1))

	var l1 int = int(s1[0])
	fmt.Println(l1)

	var l2 int = 65

	var s2 string = string(l2)

	fmt.Println(s2) // A "65"

	// as it is convertion between string and other datatypes

	s3 := fmt.Sprint(l2)
	fmt.Println(s3, reflect.TypeOf(s3)) // "65"
	b1 := false
	s4 := fmt.Sprint(b1)
	fmt.Println(s4) // "true"

	s5 := "254"

	l3, _ := strconv.Atoi(s5) // golang func or methods can return more than one return value

	// try{
	// }catch(Exception ex){
	// }
	fmt.Println(l3)
	var l5 int32 = -23434
	s6 := strconv.Itoa(int(l5))
	fmt.Println(s6)

	s7 := strconv.FormatInt(int64(l5), 10) // only for int64 to string
	fmt.Println(s7)
	name := "Jiten"
	age := 38
	isMarried := true
	s8 := fmt.Sprint("My name is ", name, " Age is ", age, " Married ", isMarried)
	fmt.Println(s8, reflect.TypeOf(s8))
}
