package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Open an existing file
	file, err := os.ReadFile("test.txt") // Use ReadFile to read the content of the file

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(file)) // Print the file content in a single line 
	fmt.Println("File opened successfully\n")
	// No need to close the file when using os.ReadFile



	// read file line by line
	fmt.Println("\n Reading file line by line..................\n")
	file2, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file2.Close()

	scanner := bufio.NewScanner(file2) // Create a scanner to read the file line by line
	for scanner.Scan() { // Scan each line
		fmt.Println(scanner.Text()) // Print each line
	}


	// DELETE A FILE
	// err = os.Remove("write.txt")

	// CHECK IF A FILE EXISTS
	_, err = os.Stat("test.txt")
	if os.IsNotExist(err) {
		fmt.Println("File does not exist.")
	} else {
		fmt.Println("File exists.")
	}

}
