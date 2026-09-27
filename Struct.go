// A struct is a custom data type that lets you group related data together.

package main

import "fmt"

type Student struct {
	Name string
	Age  int
	Roll int
}


// COMPACT VERSION
	type Address struct {
    name, street, city, state string
    Pincode int
}

func main() {

	// METHOD 1: Using struct literal

	// s1 := Student{
	// 	Name: "Rahul",
	// 	Age:  20,
	// 	Roll: 101,
	// }

	// METHOD 2: Using var keyword and then assigning values

	var s1 Student
	s1.Name = "Rahul"
	s1.Age = 20
	s1.Roll = 101

	fmt.Println(s1)
	fmt.Println("Name : ", s1.Name)
	fmt.Println("Age : ", s1.Age)
	fmt.Println("Roll : ", s1.Roll)


	// METHOD 3: Using new keyword
	s2 := new(Student)
	s2.Name = "Anjali"
	s2.Age = 19
	s2.Roll = 102

	fmt.Println(s2)
	fmt.Println("Name : ", s2.Name) 



	// COMPACT VERSION 
	var a Address
	fmt.Println(a)

	a1 := Address{"John Doe", "123 Main St", "Springfield", "IL", 62701}
	fmt.Println(a1.name)

}
