// var i int-> i is a integer type that has 8 bytes of memory (arch64)
// 8 bytes are allocated from 50000 to 50008 (example) 0XC350 - 0XC358
// pointer is a datatype that holds the address;
// if a pointer is not assigned or pointed to any address then it is nil

package main

import "fmt"

func main() {
	var i int = 10
	var ptr1 *int // at this line whare it is pointing?
	ptr1 = &i
	if ptr1 == nil {
		fmt.Println("not pointing to any address")
	}
	fmt.Println("Address of i", &i)
	fmt.Println("Address of i thru pointer", ptr1)
	fmt.Println("Value of i", i)
	fmt.Println("Value of i through pointer", *ptr1)

	changeVal(i, 20) // func has its own call stack and have a copy of parameters
	fmt.Println("value of i after changing using a normal function", i)
	changeValPtr(ptr1, 40) // I want to change the value that is in a partucualr address
	fmt.Println("Value of i after changing through a function that has a pointer", i)

	var ptr2 **int = &ptr1
	///*(*address of ptr1)) // value in address of ptr1
	fmt.Println("i=", **ptr2) // what does it prints?

	var ptr3 ***int = &ptr2
	fmt.Println("i=", ***ptr3)
	fmt.Println(ptr3)    // address of ptr2
	fmt.Println(*ptr3)   // address of ptr1
	fmt.Println(**ptr3)  // address that ptr1 holds that is address of i
	fmt.Println(***ptr3) // the value that the above line holds

	// ptr1--? &i
	// *ptr1 -? value of i
	// does ptr1 has its own address? Yes
	// How do you get address &ptr1
	// **ptr2 --> This holds the address of another pointer

}

func changeVal(i, v int) { // it has its own copy of parameters
	i = v
}

// that is i --> i is a pointer
// what does it mean? --> it is pointed to a memory location
func changeValPtr(i *int, v int) { // it has pointer and its own copy
	*i = v
}

// slice, map , ptr, interface and chan
