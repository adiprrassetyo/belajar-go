package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// pool adalah tempat untuk menyimpan data sementara yang bisa digunakan kembali
func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	} // pool kosong karena belum ada data yang disimpan di pool jika kosong maka akan mengembalikan nilai default atau nilai yang di set di New yaitu "New"

	pool.Put("Adi") // data yang disimpan di pool
	pool.Put("Syahputra")
	pool.Put("Prasetyo")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get() // data yang diambil dari pool
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data) // data yang dikembalikan ke pool setelah digunakan karena bisa digunakan kembali oleh main lain
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
