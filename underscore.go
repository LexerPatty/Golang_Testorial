// error handeling for underscore package

package main

import (
	"fmt"
)

func divide(a,b float64) (float64 , error){
	if b==0{
		return 0, fmt.Errorf("Division by zero is not allowed")
	}
	return a / b, nil
}

func main(){
	fmt.Println("Error handling for underscore package") 
	ans, _ := divide(10,0)

	// ans, err := divide(10,0)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }
	fmt.Println("Division Result: ", ans)
}