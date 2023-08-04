package main

import (
	"fmt"
	"log"
	"os"
)

func readFile() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Println("Cannot read file")
	}
	defer file.Close()
	//some other statements
}

func main() {
	defer fmt.Println("This will be printed last") //defer adalah keyword untuk menunda eksekusi suatu statement
	fmt.Println("This will be printed first")
	fmt.Println("This will be printed second")
}
