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

func main() {

	var adi Person
	adi.Name = "Rizky"

	sayHello(adi)
}
