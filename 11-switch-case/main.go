package main

func main() {

	var day uint8 = 10

	// if day == 1 {
	// 	println("Sunday")
	// } else if day == 2 {
	// 	println("Monday")
	// }

	switch day {
	case 1:
		println("Sunday")
		// there is no need of break statment.. because after a case is executed
		// automatically it breaks.
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("no day")
	}
}
