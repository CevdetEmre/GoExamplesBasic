package main

import "fmt"

func main() {

	//1. yontem

	var message string
	message = "merhaba go"
	fmt.Println(message)
	//2. yontem

	var message1 string = "abc"
	fmt.Println(message1)

	//3. yontem

	var message2 = "Merhaba go!"
	fmt.Println(message2)

	//4. yontem

	var a, b, c int = 3, 4, 5

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//5. yontem

	var message3 = "hello world"

	var d, e, f = 3, true, 4.5

	fmt.Println(message3, d, e, f)

	//6. yontem

	var c1 int
	var k, o string = "abc", "xyz"

	fmt.Println(c1, k, o)

	//7. yontem

	var p = 42

	var s, z = "xyz", true

	fmt.Println(p, s, z)

	//8. yontem

	u := 55

	fmt.Println(u)

	//9. yontem

	u1, z1 := "abc", true
	fmt.Println(u1, z1)

	//10. yontem

	a3 := "metin"
	b3 := 'M'
	c3 := `Metin`

	fmt.Println(a3, b3, c3)

	//11. yontem

	var myfloat32 float32 = 44.

	myComplex := complex(3, 4)

	fmt.Println(myfloat32)
	fmt.Println(myComplex)

}
