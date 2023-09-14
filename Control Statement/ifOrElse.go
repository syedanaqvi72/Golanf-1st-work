package main

import "fmt"

func control() {
	// Declare the student's score
	score := 60

	// Check if the student has passed or failed
	if score >= 40 {
		fmt.Println("Student has passed!")
	} else {
		fmt.Println("Student has failed.")
	}

	//To compare the value of two values, you can use six operators :

	//Comparison Operators
	//Operator	Signification
	//==	equal
	//!=	not equal
	//>	greater
	//>=	greater or equal
	//<	less
	//<=	less or equal//

	// Equal operator (==)
	num1 := 10
	num2 := 10

	if num1 == num2 {
		fmt.Println("num1 is equal to num2")
	} else {
		fmt.Println("num1 is not equal to num2")
	}

	// Not equal operator (!=)
	age1 := 25
	age2 := 30

	if age1 != age2 {
		fmt.Println("age1 is not equal to age2")
	} else {
		fmt.Println("age1 is equal to age2")
	}

	// Greater than operator (>)
	marks1 := 80
	marks2 := 60

	if marks1 > marks2 {
		fmt.Println("marks1 is greater than marks2")
	} else {
		fmt.Println("marks1 is not greater than marks2")
	}

	// Greater than or equal operator (>=)
	quantity1 := 100
	quantity2 := 100

	if quantity1 >= quantity2 {
		fmt.Println("quantity1 is greater than or equal to quantity2")
	} else {
		fmt.Println("quantity1 is not greater than or equal to quantity2")
	}

	// Less than operator (<)
	price1 := 25.0
	price2 := 30.0

	if price1 < price2 {
		fmt.Println("price1 is less than price2")
	} else {
		fmt.Println("price1 is not less than price2")
	}

	// Less than or equal operator (<=)
	score1 := 85
	score2 := 90

	if score1 <= score2 {
		fmt.Println("score1 is less than or equal to score2")
	} else {
		fmt.Println("score1 is not less than or equal to score2")
	}
}
