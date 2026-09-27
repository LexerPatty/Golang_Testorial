/*
A pointer is a variable that stores the memory address of another variable.

Normal variable → stores value
Pointer → stores address of value

Think of a house

House = value
House address = pointer

If you know the address, you can reach and modify the house.

*/

package main

import "fmt"

func main() {
	x := 10
	y := x

	y = 20

	fmt.Println("Using Normal Variable .......")
	fmt.Println(x) // 10
	fmt.Println(y) // 20

	// THESE IS THE EXAMPLE OF NO POINTERS

	c := 10
	p := &c  // pointer to x

	*p = 20   // change value using pointer

	fmt.Println("Using Pointers .......")
	fmt.Println(c)
	fmt.Println(p)	// address of c
	fmt.Println(*p) // value at address p


	/*

	What’s happening?.........
	&x → address of x
	p → stores that address
	*p → value at that address

	*/


	// check default value of pointer
	var ptr *int
	fmt.Println("Default value of pointer is :", ptr) // nil


	// pointer using function
}


