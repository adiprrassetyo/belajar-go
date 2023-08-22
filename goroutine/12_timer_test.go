package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// timer adalah sebuah channel yang akan mengirimkan data setelah durasi tertentu
func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second) // timer akan mengirimkan data setelah 5 detik
	fmt.Println(time.Now())

	time := <-timer.C // mengambil data dari channel timer setelah 5 detik
	fmt.Println(time)
}

// versi simple dari timer diatas
func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

// lebih simple karena tidak perlu membuat channel langsung ke function menggunakan AfterFunc
func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() { // AfterFunc akan menjalankan function setelah 5 detik
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())

	group.Wait()
}
