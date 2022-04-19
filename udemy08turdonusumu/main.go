package main

import (
	"fmt"
	"strconv"
)

func main() {

	var mystring = "1"

	var x = 10

	var f float32 = 2.0

	fmt.Println(mystring, x, f)

	//string to int
	number, _ := strconv.Atoi(mystring)
	fmt.Println(number)
	fmt.Printf("%T \n", number)

	//casting
	var i int = 55
	//var f1 = float64 = i yapamassın cast etmek şart
	var t float64 = float64(i)
	var u uint = uint(f)

	fmt.Println(i, t, u)
}
