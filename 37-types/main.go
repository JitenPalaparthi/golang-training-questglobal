package main

import "fmt"

// Methods can be attached to types

type myInt int

func main() {
	var m1 myInt = 100
	str1 := m1.ToString()
	fmt.Println(str1)

	var i int = 300
	str2 := myInt(i).ToString()
	fmt.Println(str2)

	mp1 := make(map[string]string)
	mp1["bangalore-3"] = "56080"
	mp1["bangalore-1"] = "56080"
	mp1["bangalore-2"] = "56082"

	str3 := myMap(mp1).ToString()
	fmt.Println(str3)
	fmt.Println("keys of a map:", myMap(mp1).GetKeys())

}

// Method
func (m myInt) ToString() string {
	return fmt.Sprint(m)
}

// This is a ToString function
func ToString(i int) string {
	return fmt.Sprint(i)
}

type myMap map[string]string

func (m myMap) ToString() string {
	str := ""
	for k, v := range m {
		str = str + k + ":" + v + ","
	}
	if len(str) > 1 {
		return str[:len(str)-1]
	}
	return str
}

func (m myMap) GetKeys() []string {
	slice := make([]string, len(m))
	c := 0
	for k := range m {
		slice[c] = k
		c++
	}
	return slice
}
