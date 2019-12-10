package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	z := [][]string{x, y}

	for _, val := range z {
		fmt.Println(val)
		for _, v := range val {
			fmt.Println(v)
		}
		fmt.Println()
	}

}
