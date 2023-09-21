package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Open the JSON file
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	// Decode the JSON data into the struct
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&person); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Print the parsed data
	fmt.Println("First Name:", person.FirstName)
	fmt.Println("Last Name:", person.LastName)
}
