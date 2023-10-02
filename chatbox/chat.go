package main

import (
	"bufio"
	"fmt"
)

func main() {
	fmt.Println("Chatbot in golang")
	fmt.Println("Commands")
	fmt.Println("_ 'Hello' to greet the chatbot")
	fmt.Println("_' How are you?' to ask about chatbot")
	fmt.Println("_ ' exit' to exit from chatbot")

	reader := bufio.NewReader(os.stdin)

	for{
		fmt.Println("You")
		userInput,-:=reader.ReadString()
	}
}
