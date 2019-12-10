package main

import "fmt"

func main() {
	var xi [5]int

	xi[0] = 0
	xi[1] = 1
	xi[2] = 2
	xi[3] = 3
	xi[4] = 4
	for i, v := range xi {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", xi)
}
