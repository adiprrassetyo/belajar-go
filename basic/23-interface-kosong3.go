package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	var data map[string]interface{}

	data = map[string]interface{}{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}
	fmt.Println(data)

	// casting variable interface kosong ke objek pointer
	var secret interface{} = &person{name: "wick", age: 27} //&person sebagai casting ke objek pointer, karena secret bertipe data interface{} maka bisa menampung semua tipe data termasuk pointer objek person (pointer objek person adalah *person)
	var name = secret.(*person).name                        //*person sebagai casting ke objek pointer
	fmt.Println(name)

	// kombinasi slice, map, dan interface kosong
	var person = []map[string]interface{}{
		{"name": "Wick", "age": 23},
		{"name": "Ethan", "age": 23},
		{"name": "Bourne", "age": 22},
	}

	for _, each := range person {
		fmt.Println(each["name"], "age is", each["age"])
	}
}
