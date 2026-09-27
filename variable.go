/*var variablename type = value*/

package main

import (
	"fmt"
)

func main() {
	var student1 string = "John" // type is string
	var student2 = "Jane"        // type is inferred
	x := 2                       // type is inferred
	var age int = 20          // type is int

	var a, y, z int = 1, 2, 3
	a, b := 10, 20


	const pi = 3.14  // cannot change
	var radius = 5  // can change
	area := pi * float64(radius*radius)

	fmt.Println(area)


	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println(age)
	fmt.Println(a, y, z)
	fmt.Println(a, b)


	/* printf
	Here we will use two formatting verbs:

%v is used to print the value of the arguments
%T is used to print the type of the arguments
*/

	fmt.Printf("Value of student1 is : %v, and Type: %T\n , Whose area value is %v and type is %T\n", student1, student1, area, pi)

	fmt.Printf("%#v\n", student2) // %#v prints the value in Go-syntax representation
}
