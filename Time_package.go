/*
use it when you want to:
Get current date & time
Measure execution time
Add or subtract time (days, hours, seconds)
Format date/time
Create delays or timers


time.Now() → current time
✅ time.Sleep() → pause execution
✅ Format() → custom date/time
✅ Add() → future/past time
✅ Since() → measure duration
*/

package main

import (
	"fmt"
	"time"
)

func main() {
    now := time.Now()

    fmt.Println(now.Year())   // 2026
	fmt.Println(now.Month())  // January
	fmt.Println(now.Day())    // 23
	fmt.Println(now.Hour())   // 10
	fmt.Println(now.Minute()) // 45
	fmt.Println(now.Second()) // 30

	// FORMAT DATE AND TIME 
	fmt.Println("Formatted Date:", time.Now().Format("02-01-2006 , Monday"))
	fmt.Println("Formatted Time:", time.Now().Format("15:04:05"))

}		
