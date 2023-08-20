package main

import "fmt"

type Man struct {
	Name string
}

/*func (man Man) Married() {
	man.Name = "Mr. " + man.Name
} */ // ketika memakai ini, maka nilai asli tidak akan berubah, hanya di dalam method saja yang berubah ketika keluar dari method, nilai asli tidak berubah

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
} // harus memakai pointer agar bisa mengubah nilai asli

func main() {
	Adi := Man{"Adi"}
	Adi.Married()

	fmt.Println(Adi.Name) //jika tidak memakai pointer, maka hasilnya tetap Adi, jika memakai pointer, maka hasilnya Mr. Adi
}
