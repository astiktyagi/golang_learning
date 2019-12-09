package main

import "fmt"

func main() {

	switch {
	case false:
		fmt.Println("this wont print")
	case true:
		fmt.Println("this will print")
	}
}
