package helper

import "fmt"

type Student struct {
	Name  string
	Grade int
}

var version = 1
var Application = "Belajar Golang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}
