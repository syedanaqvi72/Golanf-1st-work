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

func main3() {
	result := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum of numbers:", result)
}
