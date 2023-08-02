package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (a Customer) sayHuu() {
	fmt.Println("Huu from", a.Name)
}

func main() {
	var rizky Customer

	rizky.Name = "Rizky"
	rizky.Address = "Indonesia"
	rizky.Age = 20

	rizky.sayHi("dimas")
	rizky.sayHuu()
}
