// slice_name := []datatype{values}
//
// len() function - returns the length of the slice
// cap() function - returns the capacity of the slice

package main

import "fmt"

func main() {

	// Creating slices
	myslice := []string{"Go", "Slices", "Are", "Powerful"}
	myslice2 := []int{1, 2, 3, 4, 5, 6}
	myslice3 := []int{10, 11, 12}

	// Length, capacity, and values
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
	fmt.Println(myslice)

	// Slicing a slice
	fmt.Printf("Sliced position = %v\n", myslice[2:4])

	// Using make() function
	myslice4 := make([]int, 5, 10) // length 5, capacity 10
	fmt.Println(len(myslice4))
	fmt.Println(cap(myslice4))
	fmt.Println(myslice4)

	// append() function
	myslice2 = append(myslice2, 7, 8, 9)
	fmt.Println(myslice2)

	// Append one slice to another
	myslice5 := append(myslice2, myslice3...)
	fmt.Println(myslice5)

}
