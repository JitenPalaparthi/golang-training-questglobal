package main

import (
	"errors"
	"fmt"
	"reflect"
)

// map store keys and a values
// key must unique across the map
// any comparable type can be a key == (must work) (int,string,bool etc..)
// to create a map, need to use make built in function
// can use nil for map. map is declared but not instantiated then it is nil
// map returns a reference of the map
// for map range loop is used to fetch key and value
// big(o)- o(1) // for most cases
// to delete a key from a map use delete built in function
// map values are not ordered
// when a map is declared outside of the function and called as an argument
//
//	_ then the same map reference is used. It will not have a separate copy of the map
//	_ it used the same reference of the map which is created ourside of the func call
//
// i.e map,slice,pointer, interface and chan are refernece types.. but whether
//
//	_ they are created inside stack or help is taken care by runtime.
//
// just because they are reference type , it does not mean they gonna store in heap
func main() {
	i := 100             // creates a variable called i and assign a value 100 to it
	changeIntVal(i, 200) // creates a new copy of i and assign value 200
	fmt.Println("Value of I after func call", i)

	// value types and reference types

	var mp1 map[string]any // string is key and any is any type of value can be given to the map
	//var mp2 map[string]string
	// slice=make([]int,0)
	// mp= make(map[string]any)
	// make is used to create slice, map and chan
	if mp1 == nil {
		println("mp1 map is nil")
	}
	mp1 = make(map[string]any) // instantiated a map
	if mp1 == nil {
		println("mp1 map is nil")
	} else {
		println("map is not nil")
	}

	mp1["name"] = "Quest Global"                                                 //string
	mp1["number"] = 10001                                                        //int
	mp1["isRegistered"] = true                                                   //bool
	mp1["branches"] = []string{"Bangalore", "Hyderabad", "Mumbai", "California"} //slice
	mp1["dummy"] = "xyz"
	//fmt.Println(mp1)
	// for slice range loop returns --> index,value
	// for map range loop returns --> key,value
	// for channel range loop returns --> value from channel
	for key, value := range mp1 {
		fmt.Println("Key:", key, "Value:", value, "Type", reflect.TypeOf(value))
	}
	mp1["name"] = "Quest Global inc" //
	// how to check whether a key exists or not

	value, ok := mp1["phone"] // first return is value of a key. second is whether the key exists or not
	fmt.Println("Value", value, "status", ok)
	value, ok = mp1["name"]
	fmt.Println("Value", value, "status", ok)

	delete(mp1, "dummy") // to delete an element from map
	for key, value := range mp1 {
		fmt.Println("Key:", key, "Value:", value, "Type", reflect.TypeOf(value))
	}

	changeVal(mp1, "name", "Quest Global Private Limited")
	for key, value := range mp1 {
		fmt.Println("Key:", key, "Value:", value, "Type", reflect.TypeOf(value))
	}
}

func changeIntVal(i int, v int) {
	i = v
}

func changeVal(mp map[string]any, key string, val any) error {
	if mp == nil {
		return errors.New("nil map")
	}
	mp[key] = val
	return nil
}

// func allocateMem() *int { // what ever is created inside a funn that must be
// 	// deallocated there itself..
// 	ptr := new(int)
// 	*ptr = 100
// 	return ptr
// } // Google MS.. has invested mills to solve this kinds of bugs..
