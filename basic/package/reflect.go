package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type ContohLagi struct {
	Name        string `required:"true"`
	Description string `required:"true"`
}

// 1
func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

// 2
func validateInput(input interface{}) error {
	inputType := reflect.TypeOf(input)

	if inputType.Kind() != reflect.String {
		return fmt.Errorf("input harus berupa string, tetapi input berupa %v", inputType.Kind())
	}

	return nil
}

func main() {
	//2
	err := validateInput("hello")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("input valid")
	}

	//1
	sample := Sample{"Adi"}

	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	sample.Name = ""
	fmt.Println(IsValid(sample))

	contoh := ContohLagi{"Adi", "Oke"}
	fmt.Println(IsValid(contoh))
}
