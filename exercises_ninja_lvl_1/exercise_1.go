package main

import "fmt"

/*Hands-on exercise #1
Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
42
“James Bond”
true
Now print the values stored in those variables using
a single print statement
multiple print statements
*/
func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println("Printing all in same line")
	fmt.Println(x, y, z)
	fmt.Println()
	fmt.Println("Printing on separate line")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
