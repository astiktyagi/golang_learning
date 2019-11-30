package main

import "fmt"

var s string
var i int
var b bool
var m map[string]int

func main() {

	fmt.Printf("The zero value of %T is %v \n", s, s)
	fmt.Printf("The zero value of %T is %v \n", i, i)
	fmt.Printf("The zero value of %T is %v \n", b, b)
	fmt.Printf("The zero value of %T is %v \n", m, m)
}
