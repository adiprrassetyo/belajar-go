package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// WaitGroup adalah sebuah fitur di package sync yang bisa digunakan untuk menunggu main selesai dijalankan sebelum melanjutkan eksekusi kode berikutnya di main main
func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done() // mengurangi jumlah main yang sedang berjalan sebanyak 1 main dari WaitGroup (group)

	group.Add(1) // menambahkan jumlah main yang sedang berjalan sebanyak 1 main ke WaitGroup (group)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait() // menunggu semua main selesai dijalankan sebelum melanjutkan eksekusi kode berikutnya (pengganti time.Sleep(1 * time.Second)
	fmt.Println("Selesai")
}
