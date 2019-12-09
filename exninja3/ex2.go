package main

import "fmt"

func main() {

	for i := 65; i <= 65+25; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\n", i)
		}
	}
}
