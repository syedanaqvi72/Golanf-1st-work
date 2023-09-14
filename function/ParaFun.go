package main

import "fmt"

// Add two numbers and return the result
func add(a, b int) int {
	return a + b
}

func Para() {
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)
}

//We have a function called add that takes two numbers and gives us their sum.
//In our main function, we use add to add 5 and 3 together, and we get the result, which is 8.
