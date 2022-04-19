package main

import (
	"fmt"
	"time"
)

func main() {

	//Date() methodu ile kendi tarih verimizi olusturuyoruz.

	t := time.Date(2016, time.November, 10, 20, 0, 0, 0, time.UTC)

	// t adıyla tarih tipinde oluşturulan veriyi string tipinde ekrana yazıdırıyoruz.

	fmt.Printf("Go lauch at %s\n", t)

	fmt.Println("-------")

	//time.Now() ile şu anın zaman bilgisini alıyoruz.

	now := time.Now()

	fmt.Printf("The time now is %s\n", now)

	fmt.Println("-----------")

	//ilk başta oluşturdugumuz t adındaki zaman bilgisinden Ay, gun ve haftanın gunu bilgilerini ekrana yazıyoruz.

	fmt.Println("The month is", t.Month())
	fmt.Println("The day is", t.Day())
	fmt.Println("The week is", t.Weekday())

	fmt.Println("---------")

	//Tarihe 1 gün ekle

	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v, %v, %v \n",
		tomorrow.Weekday(),
		tomorrow.Month(),
		tomorrow.Day(),
		tomorrow.Year())

	fmt.Println("-----------")

	longFormat := "Monday, Junuary 2, 2006"

	fmt.Println("Tomorrow is ", tomorrow.Format(longFormat))

	shortFormat := "1/2/06"

	fmt.Printf("Tomorrow is", tomorrow.Format(shortFormat))

}
