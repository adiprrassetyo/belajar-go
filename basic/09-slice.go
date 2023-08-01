package main

import "fmt"

func main() {

	/*var fruitsA = []string{"apple", "grape"}     // slice // jumlah elemen tidak ditentukan
	var fruitsB = [2]string{"banana", "melon"}   	// array
	var fruitsC = [...]string{"papaya", "grape"} 	// array */

	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0]) // "apple"

	// hubungan slice dan array & operasi slice
	var fruits1 = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruits1[0:2] // dimulai dari index 0 sampai index 2 (tidak termasuk index 2)

	fmt.Println(newFruits) // ["apple", "grape"]

	// fungsi len() dan cap()
	var fruits3 = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits3)) // len: 4
	fmt.Println(cap(fruits3)) // cap: 4

	var aFruits3 = fruits3[0:3]
	fmt.Println(len(aFruits3)) // len: 3
	fmt.Println(cap(aFruits3)) // cap: 4

	var bFruits3 = fruits3[1:4]
	fmt.Println(len(bFruits3)) // len: 3
	fmt.Println(cap(bFruits3)) // cap: 3

	// fungsi append()
	var fruits4 = []string{"apple", "grape", "banana"}
	var cFruits4 = append(fruits4, "papaya")

	fmt.Println(fruits4)  // ["apple", "grape", "banana"]
	fmt.Println(cFruits4) // ["apple", "grape", "banana", "papaya"]

	// fungsi copy()
	var fruits5 = []string{"apple", "grape", "banana"}
	var dFruits5 = make([]string, 2)

	copy(dFruits5, fruits5)

	fmt.Println(fruits5)  // ["apple", "grape", "banana"]
	fmt.Println(dFruits5) // ["apple", "grape"]

	dst := []string{"potato", "potato", "potato"}
	src := []string{"watermelon", "pinnaple"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n)   // 2

	// fungsi append() dan copy()
	var fruits6 = []string{"apple", "grape", "banana"}
	var eFruits6 = []string{"papaya", "pinnaple"}

	var fFruits6 = append(fruits6, eFruits6...)

	fmt.Println(fruits6)  // ["apple", "grape", "banana"]
	fmt.Println(fFruits6) // ["apple", "grape", "banana", "papaya", "pinnaple"]

	var gFruits6 = make([]string, 2)
	var hFruits6 = fruits6[0:2]

	var iFruits6 = copy(gFruits6, hFruits6)

	fmt.Println(gFruits6) // ["apple", "grape"]
	fmt.Println(hFruits6) // ["apple", "grape"]
	fmt.Println(iFruits6) // 2

	// pengaksesan elemen slice dengan 3 index

	var fruits7 = []string{"apple", "grape", "banana", "melon"}
	var newFruits7 = fruits7[0:2:2]

	fmt.Println(newFruits7)      // ["apple", "grape"]
	fmt.Println(len(newFruits7)) // 2
	fmt.Println(cap(newFruits7)) // 2

	var newFruits8 = fruits7[1:3:3]

	fmt.Println(newFruits8)      // ["grape", "banana"]
	fmt.Println(len(newFruits8)) // 2
	fmt.Println(cap(newFruits8)) // 2

	var newFruits9 = fruits7[1:4:4]

	fmt.Println(newFruits9)      // ["grape", "banana", "melon"]
	fmt.Println(len(newFruits9)) // 3
	fmt.Println(cap(newFruits9)) // 3
}
