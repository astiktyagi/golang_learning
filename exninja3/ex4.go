package main

import "fmt"

func main() {

	age := 0
	for {

		if age > 26 {
			break
		}
		fmt.Println(age)
		age++
	}
}
