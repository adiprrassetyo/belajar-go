package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// penerapan function
func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

// function dengan return value
func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

// penggunaan keyword return untuk menghentikan proses function
func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

// function dengan multiple return value
func calculate(d float64) (float64, float64) { //(area float64, circumference float64) bisa seperti ini juga
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference // hanya tersisa return saja karena sudah dideklarasikan di awal function
}

// main function
func main() {

	// penerapan function
	var names = []string{"John", "Wick"}
	printMessage("halo", names)

	// function dengan return value
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	// penggunaan keyword return untuk menghentikan proses function
	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)

	// function dengan multiple return value
	var diameter float64 = 15                     // jari-jari lingkaran = 15
	var area, circumference = calculate(diameter) // panggil function calculate

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}
