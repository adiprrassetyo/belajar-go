package goroutine

import (
	"fmt"
	"testing"
	"time"
)

// race condition adalah kondisi dimana 2 atau lebih goroutine mengakses data yang sama pada waktu yang bersamaan menyebabkan data tersebut timpa menimpa satu sama lain dan itu dihitung sebagai 1 aksi
func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}
