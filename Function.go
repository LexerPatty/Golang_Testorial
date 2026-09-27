package main

import "fmt"


func simple (){
	fmt.Println("This is a simple function.")
}

func add (a,b int ) int {
	return a+b
}

func multi (a,b int) (result int) {
	result = a*b
	return result 
}


func main(){
	simple()
	sum := add(5,10)
	product := multi(5,10)
	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)
}
