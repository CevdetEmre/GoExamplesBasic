package main

import "fmt"

func main() {
	foo := -1

	switch {

	case foo == 1:
		fmt.Println("one")

	case foo == 2:
		fmt.Println("2dir")
	case foo > 2:
		fmt.Println("buyuk")
	default:
		fmt.Println("asdasdsd")
	}

}
