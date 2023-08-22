package main

import (
	"fmt"
	"sync"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// sync map adalah map yang aman untuk digunakan oleh banyak main secara bersamaan tanpa perlu ada tambahan lock
func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value) // Store digunakan untuk menambahkan data ke sync.Map
}

func TestMap(t *testing.T) {
	data := &sync.Map{} // inisialisasi sync.Map
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group) // menambahkan data ke sync.Map
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
