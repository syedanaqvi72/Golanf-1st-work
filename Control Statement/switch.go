package main

import (
	"fmt"
)

func main2() {
	dayNumber := 4

	switch dayNumber {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day number")
	}
}

//switch expression//

func main() {
	//value := 42

	//switch value {
	//case int:
	//fmt.Println("Value is an integer")
	//case float64:
	//fmt.Println("Value is a float")
	//case string:
	//fmt.Println("Value is a string")
	//default:
	fmt.Println("Value is of unknown type ")
}
