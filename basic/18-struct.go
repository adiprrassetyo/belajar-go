package main

import "fmt"

type student struct {
	name  string
	grade int
}

type Customer struct {
	Name, Address string
	Age           int
}

type Person struct {
	name, job string
	age       int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name) // customer.Name = this.Name
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuuuuu from", a.Name) // a.Name = this.Name
}

func main() {
	var Adi Customer
	Adi.Name = "Adi"
	Adi.Address = "Indonesia"
	Adi.Age = 30

	Adi.sayHi("Joko")
	Adi.sayHuuu()

	fmt.Println(Adi.Name)
	fmt.Println(Adi.Address)
	fmt.Println(Adi.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Cirebon",
		Age:     35,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Jakarta", 35}
	fmt.Println(budi)

	// struct as parameter
	var s1 = student{}
	s1.name = "wick"
	s1.grade = 2

	var s2 = student{"ethan", 2}

	var s3 = student{name: "jason"}

	fmt.Println("student 1 :", s1.name)
	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)

	//anonymous struck
	nissan := struct {
		model string
		year  int
	}{model: "nissan", year: 2020}

	//slice of struck
	human := []Person{
		{name: "adi",
			job: "makan",
			age: 18,
		},
		{name: "yow",
			job: "makan",
			age: 18,
		},
	}

	//slice of anonymous stuck
	cars := []struct {
		model string
		years int
	}{
		{model: "toyota", years: 2021},
		{model: "toyoti", years: 2021},
		{model: "toyotu", years: 2021},
	}

	_, _, _ = nissan, cars, human

}
