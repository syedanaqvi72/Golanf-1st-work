package main

import "fmt"

// Calculate the sum of multiple numbers
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func Argu() {
	result := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum of numbers:", result)
}

//We created a function called sum that can add up many numbers.
//In the main function, we used sum to add up 1, 2, 3, 4, and 5, and the result is 15.
