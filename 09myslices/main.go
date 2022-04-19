package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Wlcome to slices")

	var fruitlist = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T \n", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Banana")

	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:3])
	fmt.Println(fruitlist)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 456
	highScores[3] = 867
	//highScores[4] = 777

	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
}
