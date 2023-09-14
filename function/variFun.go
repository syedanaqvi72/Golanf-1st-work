package main

import "fmt"

// Define a function type
type operation func(int, int) int

// Add two numbers
func Addi(a, b int) int {
	return a + b
}

// Subtract two numbers
func subtract(a, b int) int {
	return a - b
}

func Vari() {
	// Declare a variable of type 'operation' and assign a function to it
	var op operation = add
	result := op(7, 3)
	fmt.Println("7 + 3 =", result)

	op = subtract
	result = op(7, 3)
	fmt.Println("7 - 3 =", result)
}

//We created two functions, add and subtract, that can add and subtract numbers.
//We also made a special kind of variable called op that can be any of these functions.
//We first set op to be the add function, so when we use op(7, 3), it's like using add(7, 3), which is 10.
//Then we change op to be the subtract function, so op(7, 3) becomes subtract(7, 3), which is 4.
