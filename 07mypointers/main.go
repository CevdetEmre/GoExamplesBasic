package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	//var ptr *int

	//fmt.Println("value of pointers", ptr)

	mynumber := 23

	var ptr = &mynumber

	fmt.Println("value of actual pointer is ", ptr)
	fmt.Println("value of actual pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("new value is : ", mynumber)
}
