package main

import "fmt"

func main() {

	// INTEGER VALUES
	var x int16 = 170
	var y int16 = 83

	// FLOAT VALUES
	a := 20.45
	b := 34.89

	// COMPLEX VALUES
	var k complex128 = complex(6, 2)
	var l complex64 = complex(9, 2)

	// COMPLEX NO . TO REAL(INT) NO. and imagianary NO........................................................
	comp2 := 13 + 33i
	realPart := real(comp2)
	imaginary := imag(comp2)


	// BOOLEAN DATATYPE 
   str1 := "GeeksforGeeks"
   str2:= "geeksForgeeks"
   str3:= "GeeksforGeeks"
   result1:= str1 == str2
   result2:= str1 == str3




	// INTEGER ARITHMETIC OPERATIONS
	fmt.Printf("Addition: %d + %d = %d\n", x, y, x+y)
	fmt.Printf("Subtraction: %d - %d = %d\n", x, y, x-y)
	fmt.Printf("Multiplication: %d * %d = %d\n", x, y, x*y)
	fmt.Printf("Division: %d / %d = %d\n", x, y, x/y)
	fmt.Printf("Remainder: %d %% %d = %d\n", x, y, x%y)

	// FLOATING-POINT OPERATIONS
	c := b - a
	fmt.Printf("\nFloating-point subtraction: %g - %g = %f\n", b, a, c)
	fmt.Printf("Floating-point addition: %g + %g = %f\n", a, b, a+b)
	fmt.Printf("Type of c: %T\n", c)

	// COMPLEX NUMBER OPERATIONS
	// fmt.Printf("\nComplex addition: %v + %v = %v\n", k, l, k+l)
	// fmt.Printf("Complex subtraction: %v - %v = %v\n", k, l, k-l)
	// fmt.Printf("Complex multiplication: %v * %v = %v\n", k, l, k*l)
	// fmt.Printf("Complex division: %v / %v = %v\n", k, l, k/l)

	fmt.Println(k)
   	fmt.Println(l)
	fmt.Printf("The type of a is %T and "+
            "the type of b is %T", k, l)


	fmt.Println("Real part of comp2:", realPart)
	fmt.Println("Imaginary part of comp2:", imaginary)


	fmt.Println(result1)
	fmt.Println(result2)

}
