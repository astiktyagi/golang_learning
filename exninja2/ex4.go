package main

import "fmt"

func main() {
	n := 42
	fmt.Printf("%d\t%b\t%x\n", n, n, n)

	n = n << 1
	fmt.Printf("%d\t%b\t%x\n", n, n, n)
}
