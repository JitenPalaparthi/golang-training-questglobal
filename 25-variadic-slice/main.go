package main

import (
	"fmt"
)

func main() {

	slice1 := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println("Slice1 before delete", slice1)
	err := delete(slice1, 12)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Slice1 after delete", slice1)
	}

	if err := delete(slice1, 3); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Slice1 after delete", slice1)
	}
	//slice2:=slice1
	//slice3:= slice1[:3]
	//slice4:= slice1[3:5]
	//slice5:= slice1[3:]

}

// interfaces, slices, maps, channels,pointers can give or check nil
func delete(slice []int, i uint) error { // errors are just values in golng
	//try{
	// i is uint and len(int)
	if i >= uint(len(slice)) {
		//return errors.New("index is beyond the length of the slice")
		return fmt.Errorf("%d index is beyind the lenght of the slice", i)
	}
	slice = append(slice[0:i], slice[i+1:]...)
	// [10,11],[13,14,15,16,17]
	// slice = append(slice[0:i],13,14,15,16,17)
	//return slice
	//}catch(Exception exp){
	//	fmt.Println(exp)
	//}
	return nil
}
