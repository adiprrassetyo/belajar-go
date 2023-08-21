package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

// 1. output akan berbeda-beda setiap kali dijalankan karena goroutine berjalan secara asynchronous (tidak berurutan) dan tidak bisa diprediksi. tergantung dari scheduler yang ada di runtime go.
func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
} // output: testing 100 ribu goroutine dengan menggunakan 2 core, membutuhkan waktu 5 detik. jika menggunakan 1 core, membutuhkan waktu 10 detik.
