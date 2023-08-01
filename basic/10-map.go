package main

import "fmt"

func main() {

	var chicken map[string]int // deklarasi map dengan key string dan value int
	chicken = map[string]int{} // inisialisasi map dengan key string dan value int

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	// inisialisasi nilai map

	// cara horizontal
	var chicken1 = map[string]int{
		"januari":  50,
		"februari": 40,
	}

	fmt.Println("januari", chicken1["januari"]) // januari 50

	// iterasi item map
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range chicken2 {
		fmt.Println(key, "  \t:", val)
	}

	// delete item map

	var chicken3 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	fmt.Println(len(chicken3)) // 4
	fmt.Println(chicken3)      // map[april:67 februari:40 januari:50 maret:34]
	delete(chicken3, "januari")
	fmt.Println(chicken3)      // map[april:67 februari:40 maret:34]
	fmt.Println(len(chicken3)) // 3

	// cek item map
	var chicken4 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	var value, isExist = chicken4["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	// kombinasi slice dan map
	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	type myMap = map[string]string

	person := myMap{"name": "reza"}
	var num uint8 = 1 //sane
	var num2 byte = 1 //same

	fmt.Printf("%T, %T, %v", num, num2, person)

}
