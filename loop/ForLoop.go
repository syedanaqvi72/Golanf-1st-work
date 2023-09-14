package main

import "fmt"

func Lop() {
	numbers := []int{1, 2, 3, 4, 5}
	// Assuming you have a slice of numbers

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Element %d: %d\n", i, numbers[i])
	}
}

//"Element %d: %d\n": This is a format string that specifies how the text should be formatted. %d is a placeholder for numbers. The first %d will be replaced with the value of i, and the second %d will be replaced with the value of numbers[i]. \n is a special character that represents a new line, so each message is printed on a new line.
