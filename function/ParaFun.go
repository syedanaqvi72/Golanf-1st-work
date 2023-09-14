package main

import "fmt"

// Add two numbers and return the result
func add(a, b int) int {
	return a + b
}

func main4() {
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)
}
