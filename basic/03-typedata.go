package main

import "fmt"

func main() {
	// tipe data numberik

	// tipe data numerik non-desimal 
		/*
		uint8  	0 ↔ 255
		uint16	0 ↔ 65535
		uint32	0 ↔ 4294967295
		uint64	0 ↔ 18446744073709551615
		uint		sama dengan uint32 atau uint64 (tergantung nilai)
		byte		sama dengan uint8
		int8		-128 ↔ 127
		int16		-32768 ↔ 32767
		int32		-2147483648 ↔ 2147483647
		int64		-9223372036854775808 ↔ 9223372036854775807
		int			sama dengan int32 atau int64 (tergantung nilai)
		rune		sama dengan int32
		*/

	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644
	
	fmt.Println("Dua : ", 2)
	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// tipe data numerik desimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber) // pembulatan 6 angka di belakang koma
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber) // pembulatan 3 angka di belakang koma

	// tipe data boolean
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)
	fmt.Println("Benar", true)

	// tipe data string
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)
	fmt.Println(len("adi"))
	fmt.Println("adi"[0])

	// tipe data string literal
	var messege = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`
	fmt.Println(messege)

	// tipe data string dengan backslash
	fmt.Println("C:\\Users\\John Wick\\Desktop\\file.exe")

	// nilai nil & zero value
	var name string
	fmt.Println(name)
		/*
		value dari string adalah "" (string kosong).
		Zero value dari bool adalah false.
		Zero value dari tipe numerik non-desimal adalah 0.
		Zero value dari tipe numerik desimal adalah 0.0.
		*/

	// constanta
	const firstName string = "john"
	const lastName = "wick"
	
	fmt.Print("halo ", firstName, "!\n")
	fmt.Print("nice to meet you ", lastName, "!\n")
}