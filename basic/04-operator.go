package main

import "fmt"

func main() {

	// operator aritmatika
	var value = (((2 + 6) % 3) * 4 - 2) / 3

	// operator perbandingan
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual) // nilai 2 (true)
		/*% d digunakan untuk mencetak nilai variabel value 
		% t digunakan untuk mencetak nilai variabel isEqual*/

	// operator logika
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight) // left && right (false)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight) // left || right (true)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse) // !left (true)


}