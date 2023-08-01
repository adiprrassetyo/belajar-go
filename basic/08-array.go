package main

import "fmt"

func main() {

	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	// deklarasi array langsung isi element
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruits)) // len() untuk menghitung jumlah element
	fmt.Println("Isi semua element \t", fruits)

	// deklarasi array tanpa jumlah element
	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers) //\t untuk tab
	fmt.Println("jumlah elemen \t:", len(numbers))

	// deklarasi array multidimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// perulangan array dengan for
	var fruits1 = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits1); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits1[i])
	} // %d untuk integer, %s untuk string

	// perulangan array dengan for range
	var fruits3 = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruits3 {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	// variable	underscore

	var fruits4 = [4]string{"apple", "grape", "banana", "melon"}

	for _, fruit := range fruits4 {
		fmt.Printf("nama buah : %s\n", fruit) // _ untuk mengabaikan index, seperti i di for biasa diatas
	}

	// array multidimensi dengan perulangan for range
	var fruits5 = [4][2]string{
		{"apple", "grape"},
		{"watermelon", "pinnaple"},
		{"orange", "mango"},
		{"grape", "papaya"},
	}

	for _, fruit := range fruits5 {
		fmt.Println(fruit[0], fruit[1]) // _ untuk mengabaikan index, seperti i di for biasa diatas
	}

	// array dengan make
	var fruits6 = make([]string, 2)
	fruits6[0] = "apple"
	fruits6[1] = "manggo"

	fmt.Println(fruits6)
}
