package main

import "fmt"

func main() {
	var arr0 [2][2]int
	arr0[0][0] = 1
	arr0[0][1] = 2
	arr0[1][0] = 3
	arr0[1][1] = 4

	arr1 := [2][2]int{{1, 2}, {3, 4}}

	for _, v1 := range arr1 {
		//fmt.Println(i1, v1)
		for _, v2 := range v1 {
			fmt.Print(v2, " ")
		}
		fmt.Println()
	}

	//arr2 := [...][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}

}
