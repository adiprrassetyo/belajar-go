package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// sync cond adalah sebuah lock yang bisa di lock dan unlock secara berkala perbedaannya dengan mutex adalah cond bisa di lock dan unlock berkali kali sedangkan mutex hanya bisa di lock dan unlock sekali saja dan cond bisa di lock dan unlock secara berkala dengan menggunakan signal dan broadcast

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait() // 1.menunggu signal atau broadcast
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i) // 1.menunggu signal atau broadcast
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal() // 2.mengirim signal ke salah satu main yang sedang menunggu secara berkala
		}
	}()

	//go func() {
	//	time.Sleep(1 * time.Second)
	//	cond.Broadcast() // 3.mengirim signal ke semua main yang sedang menunggu secara langsung tanpa menunggu secara berkala
	//}()

	group.Wait()
}
