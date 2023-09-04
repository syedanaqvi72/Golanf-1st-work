package main

import "fmt"

func main() {
	// Defining constants
	const daysInWeek = 7

	// Using constant

	weeks := 2
	totalDays := daysInWeek * weeks
	fmt.Printf("There are %d days in %d weeks\n", totalDays, weeks)
}
