package main

import "fmt"

func main() {
	// Declare a number
	number := -5

	// Check if the number is positive
	if number > 0 {
		fmt.Println("The number is positive.")
	}

	fmt.Println("End of program.")

	num1, num2, num3 := 12, 25, 8

	maximum := num1
	if num2 > maximum {
		maximum = num2
	}
	if num3 > maximum {
		maximum = num3
	}

	fmt.Printf("The maximum number is %d\n", maximum)
}
