package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (int, int) { // fungsi didalam variable (anonymous function)
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)                                  // memanggil fungsi didalam variable
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max) // mencetak nilai min dan max, %v untuk mencetak nilai apapun

	// immediately-invoked function expression (IIFE)
	var numbers1 = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers1 {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println("original number :", numbers1)
	fmt.Println("filtered number :", newNumbers)

	// simple
	sayHi := func(name string) {
		fmt.Printf("Hi %s\n", name)
	}

	sayHi("Adi")

	var sum = func(a, b int) (result int) {
		result = a + b
		return
	}

	fmt.Println(sum(1, 2))

	//
	name := "Adi"
	counter := 0

	increment := func() {
		name := "Budi"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	} // fungsi didalam berisi variable yang sama dengan variable diluar, maka variable didalam fungsi akan mengambil nilai dari variable diluar fungsi tersebut

	increment()
	increment()

	fmt.Println(counter) // 2
	fmt.Println(name)    // Adi

}
