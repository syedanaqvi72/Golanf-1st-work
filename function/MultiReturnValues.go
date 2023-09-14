package main

import "fmt"

// Divide two numbers and return both the quotient and remainder
func divide(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

func Multi() {
	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)
}

//We have a function called divide that divides two numbers and tells us both the answer and the leftover part.
//In our main function, we use divide to divide 10 by 3, and it tells us that the quotient (result of division) is 3, and the remainder (leftover) is 1.
