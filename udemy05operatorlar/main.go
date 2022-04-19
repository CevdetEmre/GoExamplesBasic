package main

import "fmt"

func main() {
	fmt.Println("")

	a := 10
	b := 20

	total := a + b
	fmt.Println(total)

	total = total - 5
	fmt.Println(total)

	total *= 20
	fmt.Println(total)

	total /= 10
	fmt.Println(total)

	total = a % b
	fmt.Println(total)

	total++
	fmt.Println(total)

}
