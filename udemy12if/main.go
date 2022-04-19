package main

import "fmt"

func main() {

	a := 10

	b := 10

	if b < a {

		fmt.Println("buyuktur")

	} else if b == a {

		fmt.Println("esittir")

	} else {
		fmt.Println("kucuktur")
	}

	foo := 1

	if foo == 1 {

		println("bar")
	}

	if foo1 := 2; foo1 == 1 {
		println("baa")

	} else {
		println("caA")
	}
}
