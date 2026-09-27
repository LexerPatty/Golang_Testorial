/*
Real-life example :
Think of a library book 📚
You open the book
You read the book
You return the book when you leave

defer is like saying:
“No matter what happens, return the book when I’m done.”

Most common use: closing files
*/

package main

import "fmt"

func main() {
	// defer fmt.Println("World") // run last
	fmt.Println("Hello") // run first


	// order of execution:
	defer fmt.Println("1")   // run third
	defer fmt.Println("2")   // run second
	defer fmt.Println("3")       // run first

}
