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
}
