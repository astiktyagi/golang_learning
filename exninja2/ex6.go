package main

import "fmt"

const (
	_  = iota
	y1 = 2019 + iota
	y2
	y3
	y4
)

func main() {

	fmt.Println(y1)
	fmt.Println(y2)
	fmt.Println(y3)
	fmt.Println(y4)
}
