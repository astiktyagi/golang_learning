package main

import "fmt"

func main() {
	favSport := "volleyball"

	switch favSport {
	case "cricket":
		fmt.Println("fav sport is cricket")
	case "hockey":
		fmt.Println("fav sport is hockey")
	case "volleyball":
		fmt.Println("fav sport is volleyball")
	}
}
