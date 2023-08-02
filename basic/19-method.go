package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
} // struct

type Man struct {
	Name string
} // struct

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
} // method pointer, * sebagai penanda pointer (reference)

func (s student) sayHello() {
	fmt.Println("halo", s.name)
} // method (function) untuk struct student (s)

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
} // method (function) untuk struct student (s)

func main() {

	// method
	var s1 = student{"john wick", 21}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("nama panggilan :", name)

	// method pointer
	eko := Man{"Eko"}
	eko.Married()

	fmt.Println(eko.Name)

	/*
		func sayHello() { // function
		func (s student) sayHello() { // method

		func getNameAt(i int) string {
		func (s student) getNameAt(i int) string {
	*/
}
