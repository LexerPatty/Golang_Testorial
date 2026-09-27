/*
A goroutine is a lightweight thread managed by Go.
It lets your program do multiple tasks at the same time (concurrently).

Real-life example:

Imagine a restaurant
One waiter doing everything → slow
Multiple waiters handling tables → fast

Each waiter = goroutine
*/

// concurracy [multiple task at the same time (a chief in a kitchen working on several dishes )]
// vs
// parallelism [actually executing multiple tasks simuntaneously (multiple chefs working different dishes at the same time )]

package main

import (
	"fmt"
	"time"
)

func sayHello(){	
	fmt.Println("Hello from goroutine")
	time.Sleep (2000*time.Millisecond) // simulate a long task
}

func sayHi(){
	fmt.Println("Hi from goroutine")
}


func main(){
	fmt.Println("Main function started")
	go sayHello() // start a goroutine
	sayHi()    // start another goroutine


}
