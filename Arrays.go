/*
var array_name = [length]datatype{values} // here length is defined
or
var array_name = [...]datatype{values} // here length is inferred
*/

package main

import (
	"fmt"
)

func main(){
	var arr1 = [5]int{10,20,30,40,50} // length defined
	arr2:= [...]string{"Go","Python","Java","C++"} // length inferred


	fmt.Println("Array 1:", arr1) 
	fmt.Println("Array 2:", arr2)


	// access element of an array 
	fmt.Println("First element of arr1:", arr1[0])

	// CHANGE ELEMENT OF AN ARRAY
	arr1[2] = 100
	fmt.Println("Modified Array 1:", arr1)
}

/* default of int = 0
bool = false 
string - empty space 

eg car price [5] string / int / bool 
	fmt.Println[price]
*/



/* DIFFERENCE BETWEEN ARRAY AND SLICE */