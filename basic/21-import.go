package main

import (
	"belajar-go/basic/helper" // import package helper yang berada di folder basic (folder yang sama dengan folder main.go) // apabila ditambah . maka tidak perlu menuliskan helper. lagi
	"fmt"
)

func main() {
	helper.SayHello("Adi") // dapat diakses karena huruf awalnya besar (public)
	// helper.sayGoodbye("Adi") // error karena huruf awalnya kecil (private)
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error

	var s1 = helper.Student{"ethan", 21} // menjadi student saja jika ditambah . pada import awal
	fmt.Println("name ", s1.Name)
	fmt.Println("grade", s1.Grade)
}
