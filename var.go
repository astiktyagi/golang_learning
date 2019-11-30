package main

import "fmt"

/* false for booleans, 0 for numeric types,
"" for strings, and
nil for pointers, functions, interfaces, slices, channels, and maps. */

func main() {

	/* short declaration operator
	declares and assigns the vairable*/
	x := 42
	fmt.Println(x)

	/* var keyword
	can be declared outside function body */
	var y = 50
	fmt.Println(y)

	/* declare a variable with identifier
	assigns the ZERO value of type to the value*/
	var z []int

	fmt.Println(z)
}
