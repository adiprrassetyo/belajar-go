package main

import "fmt"

type Hasname interface {
	GetName() string
}

func sayHello(hasname Hasname) {
	fmt.Println("Hello", hasname.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {

	var adi Person
	adi.Name = "Rizky"

	sayHello(adi)

	cat := Animal{
		Name: "Meow",
	}
	sayHello(cat)
}
