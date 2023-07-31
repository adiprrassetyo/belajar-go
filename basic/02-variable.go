package main

import "fmt"

func main() {

	// deklarasi variabel beserta tipe data dan menggunakan var
  var firstName string = "john"
	firstName = "johan" // variabel firstName dapat diubah nilainya
  var lastName string
  lastName = "wick"

	// deklarasi variabel tanpa tipe data
	country := "Indonesia"
	fmt.Println(country)

	lastName = "ethan"
	lastName = "bourne"

	// deklarasi multiple variabel
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	fmt.Println(first, second, third)

	// deklarasi variable underscore
	// digunakan untuk menampung nilai yang tidak digunakan
	_ = "belajar Golang"
	_ = "Golang itu mudah"
	nama, _ := "john", "wick"
	fmt.Println(nama)

	// deklarasi variable menggunakan keyword new
	name := new(string)
	fmt.Println(name)   // 0x20818a220
	fmt.Println(*name)  // ""

	// deklarasi variable menggunakan keyword make
	// digunakan untuk membuat slice, map, dan channel
	// slice
	kota := make([]string, 2)
	kota[0] = "Jakarta"
	kota[1] = "Bandung"
	fmt.Println(kota)


	//data type conversion
	var age int8 = 22
	var age64 int64 = int64(age)
	var h = firstName[0]
	var h8 = string(h)

	fmt.Println(age, age64, h, h8)


  fmt.Printf("halo %s %s!\n", firstName, lastName) // %s => string
	fmt.Printf("halo john wick!\n")
	fmt.Println("halo", firstName, lastName + "!")
}