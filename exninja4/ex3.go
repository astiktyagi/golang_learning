package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(xi[:5])
	fmt.Println(xi[5:])
	fmt.Println(xi[3:7])
	fmt.Println(xi[2:6])
}
