package udtslice

import "fmt"

type AnySlice []any // User defined type

func (as AnySlice) Add() any {
	sum := 0.0
	for _, v := range as {

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
		default:
			fmt.Println("supplied a type on which cannot perform numeric addition")
		}

	}
	return sum
}
