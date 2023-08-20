package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Adi Kurniawan", "Adi"))
	fmt.Println(strings.Contains("Adi Kurniawan", "Budi"))

	fmt.Println(strings.Split("Adi Kurniawna Prasetyo", " "))

	fmt.Println(strings.ToLower("Adi Kurniawan Prasetyo"))
	fmt.Println(strings.ToUpper("Adi Kurniawan Prasetyo"))
	fmt.Println(strings.ToTitle("Adi kurniawan Prasetyo"))

	fmt.Println(strings.Trim("      Adi Kurniawan     ", " "))
	fmt.Println(strings.ReplaceAll("Adi Joko Adi", "Adi", "Budi"))
}
