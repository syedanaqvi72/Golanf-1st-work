package main

import "fmt"

func Closures() {
	// Define and immediately invoke an anonymous function
	result := func(a, b int) int {
		return a * b
	}(5, 3)

	fmt.Println("5 * 3 =", result)
}

//We made a "secret" function that multiplies two numbers, but it doesn't have a name.
//Right away, we use this "secret" function to multiply 5 by 3, and the result is 15.
