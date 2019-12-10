package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}

	for i, v := range xi {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", xi)
}
