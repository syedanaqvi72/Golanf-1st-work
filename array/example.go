package main

import "fmt"

func main() {

	// var VariableName =[length] type {value}
	//var arr1 = [4]int{1, 2, 3, 4}
	//arr2:=[3]int{100, 200, 300,}
	//fmt.Println(arr1)
	//fmt.Println(arr2)
	var arr3 = [...]int{100, 200, 300}
	fmt.Println(arr3)
	fmt.Println(len(arr3))
	var names = [3]string{"A", "B", "C"}
	fmt.Println(names)
	names[1] = "test"
	fmt.Println(names[1])
	fmt.Println(names)

	// Declare and initialize an array of integers with a fixed size of 5
	var numbers [5]int

	// Assign values to the elements of the array
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	// Access and print the elements of the array
	fmt.Println("Array elements:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Element %d: %d\n", i, numbers[i])
	}
}
