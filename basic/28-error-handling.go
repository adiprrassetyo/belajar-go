package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func IsFileExists(filename string) (bool, error) {
	file, err := os.Open(filename) // open file
	defer file.Close()
	if err != nil {
		return false, fmt.Errorf("Error in IsFileExists, err: %w", err) // wrapping error
	}

	return true, nil
}

func main() {

	// Untuk membuat error, kita dapat menggunakan dua pendekatan yaitu menggunakan errors.New dan fmt.Errorf.

	// var err error = errors.New("Ini adalah error")
	// var err error = fmt.Errorf("Ini adalah error")

	var number int
	var err error

	test = "abc"
	number, err = strconv.Atoi(test)

	if err != nil { // error handling
		// terjadi error
		fmt.Println(err)
		return
	}

	// tidak terjadi error
	fmt.Println(number)

	isExist, err := IsFileExists("file.txt")
	if err != nil {
		if errWrap := errors.Unwrap(err); errWrap != nil {
			fmt.Println(errWrap)
		}
	}
}
