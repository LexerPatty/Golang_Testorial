package main

import (
	"fmt"
)

func main() {
	for i := 0 ; i < 10 ; i++{
		fmt.Println ("Numbers is : " , i)
	}

	//INFINITE LOOP WITH BREAK STATEMENT as while loop 
	counter := 0
	for {
		fmt.Println("Infinite Loop %v" , counter)
		counter++
		if counter ==5{
			break 
		}
	}

	// for- RANGE loop 
	numbers := []int {1,2,3,4,5}
	for index , value := range numbers {
		fmt.Printf("Index %d, Value: %d\n" , index, value)
	}
}