package main

import "fmt"

func main() {
	greet() // calling a function
	greetName("Quest Global")
	add, sub, mul, div, mod := Arthemetic(10, 20)
	fmt.Println("Addition:", add, "Substraction:", sub, "Multiplication:", mul, "Division:", div, "Mod:", mod)
	a1, p1 := AreaAndPerimeterOfRect(12.13, 14.56)
	fmt.Println("Area a1:", a1, "Perimeter p1:", p1)
	a2, _ := AreaAndPerimeterOfRect(12.34, 12.45)
	fmt.Println("Area a2:", a2)
	_, p2 := AreaAndPerimeterOfRect(12.34, 12.45)
	fmt.Println("Perimeter p2:", p2)

}
func greet() { // function stack frame contains -> arguments, variables, return values
	str := "hello world"
	println(str)
}

func greetName(name string) {
	{
		str := "hello"
		fmt.Println(str, name)
	}
	fmt.Println("How are you", name)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) (c int) {
	c = a + b
	return c
}

func multiply(a int, b int) int {
	c := a * b
	return c
}

func Arthemetic(a, b int) (int, int, int, int, float32) { // function with multiple return parameters
	r1 := a + b
	r2 := a - b
	r3 := a * b
	r4 := a / b
	r5 := a % b
	//return a + b, a - b, a * b
	return r1, r2, r3, r4, float32(r5)
}

func AreaAndPerimeterOfRect(a, b float32) (area float32, perimeter float32) { // function with multiple return parameters
	area = a * b
	perimeter = 2 * (a + b)
	return area, perimeter
}
