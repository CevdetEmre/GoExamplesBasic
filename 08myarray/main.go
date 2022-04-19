package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golangs")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("list is :", fruitList)
	fmt.Println("list len is", len(fruitList))

	var veglist = [4]string{"potato", "beans", "mushroom"}
	fmt.Println("list is : ", veglist)
	fmt.Println("list len is : ", len(veglist))
}
