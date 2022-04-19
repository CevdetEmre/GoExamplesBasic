package main

import "fmt"

const LoginToken string = "asdasdasd" //Public

func main() {

	//string
	var username string = "test"
	fmt.Println(username)
	fmt.Printf("variable: %T \n ", username)
	//bool
	var username1 bool = true
	fmt.Println(username1)
	fmt.Printf("variable: %T \n ", username1)
	//int
	var username2 int = 16
	fmt.Println(username2)
	fmt.Printf("variable: %T \n ", username2)
	//float
	var username3 float32 = 16.45
	fmt.Println(username3)
	fmt.Printf("variable: %T \n ", username3)
	//complex
	var username4 complex64 = 16e-12
	fmt.Println(username4)
	fmt.Printf("variable: %T \n ", username4)

	//default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable: %T \n ", anotherVariable)

	//implicit type

	var website = "web"
	fmt.Println(website)
	fmt.Printf("variable: %T \n ", website)

	// no var style

	noVarStyle := 300

	fmt.Println(noVarStyle)

	fmt.Println(LoginToken)
}
