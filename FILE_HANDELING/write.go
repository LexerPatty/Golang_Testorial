package main

import (
	"fmt"
	"os"
)


func main(){
	// craete a file for writing
	file, err := os.Create("write.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Write some content to the file
	content := "Hello, this is a test file created in Go!"

	
	_, err = file.WriteString(content)
	fmt.Println("File created and written successfully.")


	// APPEND TO THE FILE
	f, err := os.OpenFile("write.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	f.WriteString("\nThis line is appended to the file.")


}


