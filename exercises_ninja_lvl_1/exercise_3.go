package main

import "fmt"

/*
Hands-on exercise #3
Using the code from the previous exercise,
At the package level scope, assign the following values to the three variables
for p assign 42
for q assign “James Bond”
for r assign true
in func main
use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
print out the value stored by variable “s”

*/

var p = 42
var q = "James Bond"
var r = true

func main() {

	s := fmt.Sprintf("%v %v %v", p, q, r)
	fmt.Println(s)

}
