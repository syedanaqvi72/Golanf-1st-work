package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	var person Person

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&person); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("First Name:", person.FirstName)
	fmt.Println("Last Name:", person.LastName)
}
