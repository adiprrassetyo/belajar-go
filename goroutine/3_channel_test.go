package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// 1.channel
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // jangan lupa close channel, jika tidak maka akan terjadi memory leak, karena channel tidak akan dihapus dari memory, dan akan terus menunggu data yang dikirimkan// defer hanya sebagai jaminan, jika ada error, maka channel akan tetap di close, jika tidak ada error, maka channel akan tetap di close, defer akan selalu dijalankan

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Adi Syahputra Prasetyo" // send data to channel
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <-channel // receive data from channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// 2
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Adi Syahputra Prasetyo"
}

// 2.channel as parameter
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Adi Syahputra Prasetyo"
} // channel hanya bisa menerima data

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
} // channel hanya bisa mengirim data
// 3
func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// 4. buffered channel (channel yang bisa menampung data lebih dari 1) dengan cara menambahkan parameter kedua pada make(chan string, 3)
	channel := make(chan string, 3) // buffered channel dengan kapasitas 3 (PENYIMPANAN CADANGAN)
	defer close(channel)

	go func() {
		channel <- "Adi"
		channel <- "Syahputra"
		channel <- "Prasetyo"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

// 5. range channel digunakan untuk mempermudah pengambilan data dari channel yang berulang kali (looping) sampai channel di close (channel tidak menerima data lagi) daripada melakukan pengecekan secara manual
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

// 6. select channel digunakan untuk mengecek data yang dikirimkan dari beberapa channel sekaligus, jika ada data yang dikirimkan dari salah satu channel, maka akan diambil data tersebut dan akan di proses sesuai dengan case yang ada di select channel tersebut (mirip seperti switch case) tetapi select channel tidak bisa digunakan untuk mengecek data dari channel yang tidak memiliki data (channel yang tidak memiliki data akan di skip) random
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		} // select channel tidak memiliki default case, jika tidak ada data yang dikirimkan dari channel, maka select channel akan menunggu data dari channel tersebut (tidak akan di skip) dan akan terus menunggu data dari channel tersebut

		if counter == 2 {
			break
		}
	}
}

// 6. default select channel 
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data") // default case akan dijalankan jika tidak ada data yang dikirimkan dari channel (tidak akan menunggu data dari channel tersebut) dan akan di skip sampai ada data yang dikirimkan dari channel
		}

		if counter == 2 {
			break
		}
	}

	// output :
	// Menunggu Data
	// Menunggu Data
	// Data dari Channel 1 Adi Syahputra Prasetyo
	// Data dari Channel 2 Adi Syahputra Prasetyo
}
