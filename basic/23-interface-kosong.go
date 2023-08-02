package main

import (
	"fmt"
)

// interface kosong
type Animal interface {
	Speak()
} // interface kosong adalah interface yang tidak memiliki deklarasi method satupun sehingga semua tipe data bisa masuk ke dalamnya (semua tipe data pasti memiliki method kosong) // interface kosong biasanya digunakan untuk menampung tipe data yang berbeda-beda (polymorphism) // interface kosong juga biasa disebut dengan empty interface // interface kosong juga biasa digunakan untuk menampung tipe data yang tidak diketahui

type Cat struct{} // Cat adalah struct kosong

func (c *Cat) Speak() {
	fmt.Println("Meow!")
} // method Speak() adalah method kosong yang mengembalikan string "Meow!", method ini merupakan method dari interface Animal sehingga Cat merupakan tipe data yang bisa masuk ke dalam interface Animal

type Dog struct{}

func (d *Dog) Speak() {
	fmt.Println("Woof!")
} // method Speak() adalah method kosong yang mengembalikan string "Woof!", method ini merupakan method dari interface Animal sehingga Dog merupakan tipe data yang bisa masuk ke dalam interface Animal

func main() {

	// interface kosong
	var animal Animal // Animal adalah interface kosong

	animal = &Cat{} // Cat adalah struct kosong yang merupakan tipe data yang bisa masuk ke dalam interface Animal
	animal.Speak()  // Output: Meow!

	animal = &Dog{}
	animal.Speak() // Output: Woof!
}
