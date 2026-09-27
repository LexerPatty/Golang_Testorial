package main

import (
	"fmt"
	"os"
)


func main() {

	// Create a new file
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
}
defer file.Close() // Ensure the file is closed after operations
fmt.Println("File created successfully")

}