package main

import "fmt"

func main() {
	// Creating a map
	myMap := make(map[string]int)

	// Inserting entries
	myMap["apple"] = 5
	myMap["banana"] = 3
	myMap["cherry"] = 8

	// Retrieving entries
	value, exists := myMap["banana"]
	if exists {
		fmt.Println("Banana count:", value)
	} else {
		fmt.Println("Banana not found in the map")
	}
}

//types  of creating map
//// Using make function
//myMap := make(map[KeyType]ValueType)

// Using map literal
//myMap := map[KeyType]ValueType{}

// Inserting entries
//myMap[key] = value

// Retrieving entries
//value := myMap[key]

//Exit value
//value, exists := myMap[key]
