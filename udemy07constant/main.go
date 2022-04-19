package main

import "fmt"

type Brand string

const (
	FACEBOOK  Brand = "Facebook"
	MICROSOFT Brand = "Microsoft"
	GOOGLE    Brand = "Google"
	DIJIBIL   Brand = "Dijibil"
)

func PrintBrand(aq Brand) {
	fmt.Println(aq)
}
func main() {
	PrintBrand(GOOGLE)
	PrintBrand(DIJIBIL)

	// Compile Error : undefined: VERIZON
	// PrintBrand(VERIZON)

	// Compile error : cannot use 25 (type int) as type Bran
	// PrintBrand (25)
}
