package main

import (
	"fmt"
	"testing"
	"time"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ticker digunakan untuk mengirim data secara berkala setiap interval waktu yang ditentukan
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second) // ticker akan mengirimkan data setiap 1 detik

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}() // setelah 5 detik, ticker akan berhenti

	for time := range ticker.C {
		fmt.Println(time)
	}
}

// tick adalah versi simple dari ticker diatas karena tidak perlu membuat channel terlebih dahulu
func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}
