package main

import "fmt"

func main() {

	if false {
		fmt.Println("true")
	} else {
		fmt.Println("else")
	}

	if false {
		fmt.Println("false")
	} else if 2 == 2 {
		fmt.Println("else if")
	}
}
