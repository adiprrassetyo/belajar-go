package main

import (
	"fmt"
	"sync"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce) // hanya akan dieksekusi sekali saja
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter) // ouputnya menjadi 1, bukan 100
}
