package calc

import "fmt"

type Calculator struct {
	// Nums ...any // variadic must be a function parameter.. it cannot be used as a type for user defined type
	Nums []any
}

// Add ,Sub, Multiply
// after a careful thought , it is decided to do addtion on numracic types
func (c *Calculator) Add() any {
	sum := 0.0 // float64      10+10 => 20.0 but 10+10.5 okay to give 20.5 but not okay to give 20
	//var sum any
	for _, v := range c.Nums {
		//sum += v.(int) // v is any so need to do type assertion
		// to  perform operation .. you need to assert after checking the exact type

		// switch v.(type) {
		// case int:
		// 	sum += float64(v.(int))
		// case uint8:
		// 	sum += float64(v.(uint8))
		// case uint16:
		// 	sum += float64(v.(uint16))
		// case uint32:
		// 	sum += float64(v.(uint32))
		// case uint64:
		// 	sum += float64(v.(uint64))

		// case int8:
		// 	sum += float64(v.(int8))
		// case int16:
		// 	sum += float64(v.(int16))
		// case int32:
		// 	sum += float64(v.(int32))
		// case int64:
		// 	sum += float64(v.(int64))

		// case float32:
		// 	sum += float64(v.(float32))
		// case float64:
		// 	sum += v.(float64)
		// default:
		// 	fmt.Println("supplied a type on which cannot perform numeric addition")
		// }

		switch val := v.(type) {
		case int:
			sum += float64(val)
		case uint8:
			sum += float64(val)
		case uint16:
			sum += float64(val)
		case uint32:
			sum += float64(val)
		case uint64:
			sum += float64(val)

		case int8:
			sum += float64(val)
		case int16:
			sum += float64(val)
		case int32:
			sum += float64(val)
		case int64:
			sum += float64(val)

		case float32:
			sum += float64(val)
		case float64:
			sum += val

		// case uint8, uint16, uint32, uint64, int8, int16, int32, int64, int, float32, float64:
		// 	sum = sum + float64(val)
		default:
			fmt.Println("supplied a type on which cannot perform numeric addition")
		}

	}

	return sum
}

// what is panic ?
// What all data types you would like to perform addion
