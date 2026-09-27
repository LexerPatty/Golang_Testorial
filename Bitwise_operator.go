// Go program to illustrate the
// use of bitwise operators
package main

import "fmt"
 
func main() {
   p:= 34
   q:= 20
    
   // & (bitwise AND)
   result1:= p & q
   fmt.Printf("Result of p & q = %d", result1)
    
   // | (bitwise OR)
   result2:= p | q
   fmt.Printf("\nResult of p | q = %d", result2)
    
   // ^ (bitwise XOR)
   result3:= p ^ q
   fmt.Printf("\nResult of p ^ q = %d", result3)
    
   // << (left shift)
   result4:= p << 1
   fmt.Printf("\nResult of p << 1 = %d", result4)
    
   // >> (right shift)
   result5:= p >> 1
   fmt.Printf("\nResult of p >> 1 = %d", result5)
    
   // &^ (AND NOT)
   result6:= p &^ q
   fmt.Printf("\nResult of p &^ q = %d", result6)
    
    
}



/*
“=”(Simple Assignment): This is the simplest assignment operator. This operator is used to assign the value on the right to the variable on the left.
“+=”(Add Assignment): This operator is a combination of ‘+’ and ‘=’ operators. This operator first adds the current value of the variable on left to the value on the right and then assigns the result to the variable on the left.
“-=”(Subtract Assignment): This operator is a combination of ‘-‘ and ‘=’ operators. This operator first subtracts the current value of the variable on left from the value on the right and then assigns the result to the variable on the left.
“*=”(Multiply Assignment): This operator is a combination of ‘*’ and ‘=’ operators. This operator first multiplies the current value of the variable on left to the value on the right and then assigns the result to the variable on the left.
“/=”(Division Assignment): This operator is a combination of ‘/’ and ‘=’ operators. This operator first divides the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.
“%=”(Modulus Assignment): This operator is a combination of ‘%’ and ‘=’ operators. This operator first modulo the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.
“&=”(Bitwise AND Assignment): This operator is a combination of ‘&’ and ‘=’ operators. This operator first “Bitwise AND” the current value of the variable on the left by the value on the right and then assigns the result to the variable on the left.
“^=”(Bitwise Exclusive OR): This operator is a combination of ‘^’ and ‘=’ operators. This operator first “Bitwise Exclusive OR” the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.
“|=”(Bitwise Inclusive OR): This operator is a combination of ‘|’ and ‘=’ operators. This operator first “Bitwise Inclusive OR” the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.
"<<="(Left shift AND assignment operator): This operator is a combination of ‘<<’ and ‘=’ operators. This operator first “Left shift AND” the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.
">>="(Right shift AND assignment operator): This operator is a combination of ‘>>’ and ‘=’ operators. This operator first “Right shift AND” the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.
*/