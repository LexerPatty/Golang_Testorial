package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	// var age int

	// fmt.Print("Enter your name: ")
	// fmt.Scanln(&name)
	// fmt.Print("Enter your age: ")
	// fmt.Scanln(&age)

// USING bufio for input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ = reader.ReadString('\n')

	fmt.Printf("Hello %s", name)
}
