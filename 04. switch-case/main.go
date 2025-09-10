package main

import "fmt"

func main() {

	//! switch statement with break :
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("it is Monday")
		break //! if we use break, then the switch statement will not execute the next case
	case "Tuesday":
		fmt.Println("it is Tuesday")
		break
	case "Wednesday":
		fmt.Println("it is Wednesday")
		break
	case "Thursday":
		fmt.Println("it is Thursday")
		break
	case "Friday":
		fmt.Println("it is Friday")
		break
	case "Saturday":
		fmt.Println("it is Saturday")
		break
	case "Sunday":
		fmt.Println("it is Sunday")
		break
	default:
		fmt.Println("it is not a day")
	}
}
