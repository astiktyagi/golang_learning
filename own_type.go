package main

import "fmt"

var x int

type hotdog int

var a hotdog

func main() {

	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// conversion, not casting
	x = int(a)
	fmt.Println(x)
	fmt.Printf("%T\n", x)

}
