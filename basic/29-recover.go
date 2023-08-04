package main

import "fmt"

func division(first, second int) int {
	defer func() {
		if v := recover(); v != nil { // menangkap 'panic' yang terjadi
			fmt.Println(v)
		}
	}()

	return first / second // terjadi panic di sini
}

func main() {
	div := division(10, 0)

	// kode akan terus berjalan karena sudah dilakukan recover
	fmt.Println(div)
	fmt.Println("Program Selesai")
}
