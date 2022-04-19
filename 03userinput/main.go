package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for our Pizza: ")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thans for reading, ", input)
	fmt.Printf("type: %T", input)
}
