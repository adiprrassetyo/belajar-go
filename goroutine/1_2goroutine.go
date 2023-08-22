package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2) // set jumlah core yang digunakan untuk eksekusi program

	go print(5, "halo")
	print(5, "apa kabar")

	var input string
	fmt.Scanln(&input)
}

// output 1
// 1 apa kabar
// 2 apa kabar
// 3 apa kabar
// 4 apa kabar
// 5 apa kabar
// 1 halo
// 2 halo
// 3 halo
// 4 halo
// 5 halo

// hasil ekseskusi pertama dan kedua bisa berbeda karena main berjalan secara asynchronous (tidak berurutan) dan tidak bisa diprediksi. tergantung dari scheduler yang ada di runtime go. atau bisa juga diakibatkan oleh perbedaan jumlah core yang digunakan.
