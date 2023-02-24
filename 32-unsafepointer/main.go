package main

import (
	"fmt"
	"unsafe"
)

// in go directly pointer arthemetic is not allowed
// to perform pointer arthemetic you need to use uintptr and unsafe.Pointer
func main() {

	slice1 := []int{10, 20, 30, 40, 50}
	var addr1 uintptr = uintptr(unsafe.Pointer(&slice1[0]))

	for i := 0; i < len(slice1); i++ {
		fmt.Println(*(*int)(unsafe.Pointer(addr1)))
		addr1 = addr1 + unsafe.Sizeof(slice1[0])
	}

	slice2 := []string{"Bangalore", "hyderabad", "Trivandrum", "Chennai", "Delhi"}
	var addr2 uintptr = uintptr(unsafe.Pointer(&slice2[0]))

	for i := 0; i < len(slice2); i++ {
		fmt.Println(*(*string)(unsafe.Pointer(addr2)))
		addr2 = addr2 + unsafe.Sizeof(slice2[i+1])
	}

	// This is about slice capacity

	slice3 := []int{10, 20}
	fmt.Println("cap of slice3", cap(slice3), "address of slice3", &slice3[0])
	slice3 = append(slice3, 30, 40, 50)
	fmt.Println("cap of slice3", cap(slice3), "address of slice3", &slice3[0])
	slice3 = append(slice3, 60)
	fmt.Println("cap of slice3", cap(slice3), "address of slice3", &slice3[0])
	slice3 = append(slice3, 70)
	fmt.Println("cap of slice3", cap(slice3), "address of slice3", &slice3[0])
	// then previous allocations are garbage colleted
}
