package main

import (
	"fmt"
	"strings"
)

// fungsi dengan parameter biasa dan variadic
func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

func calculate(numbers ...int) float64 { // variadic function dengan parameter numbers bertipe data slice int dan mengembalikan nilai float64 (rata-rata)
	var total int = 0
	for _, number := range numbers {
		total += number
	} // total = total + number

	var avg = float64(total) / float64(len(numbers)) // menghitung rata-rata
	return avg
}

func main() {
	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3) // memanggil function calculate dengan parameter numbers berisi 10 angka
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)    // mengubah nilai avg menjadi string dengan 2 angka di belakang koma
	fmt.Println(msg)

	// pengisian parameter fungsi dengan slice
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var avg1 = calculate(numbers...)
	var msg1 = fmt.Sprintf("Rata-rata : %.2f", avg1)

	fmt.Println(msg1)

	// variadic function
	var hobbies = []string{"sleeping", "eating"}
	yourHobbies("wick", hobbies...)
}
