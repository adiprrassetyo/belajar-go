package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Adi Syahputra", "Adi"))
	fmt.Println(strings.Contains("Adi Syahputra", "Budi"))

	fmt.Println(strings.Split("Adi Kurniawna Prasetyo", " "))

	fmt.Println(strings.ToLower("Adi Syahputra Prasetyo"))
	fmt.Println(strings.ToUpper("Adi Syahputra Prasetyo"))
	fmt.Println(strings.ToTitle("Adi Syahputra Prasetyo"))

	fmt.Println(strings.Trim("      Adi Syahputra     ", " "))
	fmt.Println(strings.ReplaceAll("Adi Joko Adi", "Adi", "Budi"))
}
