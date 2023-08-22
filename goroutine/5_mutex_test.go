package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mutex adalah mutual exclusion, artinya hanya satu main yang bisa mengakses data pada satu waktu, solusi atas race condition
// 1
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock() // lock
				x = x + 1
				mutex.Unlock() // unlock
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

// RWMutex adalah mutex yang bisa dibaca oleh banyak main, tapi hanya bisa ditulis oleh satu main
// 2
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock() // lock
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock() // unlock
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock() // read lock
	balance := account.Balance
	account.RWMutex.RUnlock() // read unlock
	return balance
}

// RWMutex cocok digunakan untuk kasus dimana ada banyak main yang membaca data, tapi jarang ada main yang menulis data
func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance", account.GetBalance())
	// output:
	// secara berurut 1 sampai 10000
	// Total Balance 10000
	// jika tidak menggunakan RWMutex, outputnya tidak berurut
	// dan total ballance nya bisa berkurang dari 10000
}

// 2
type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name) // deadlock terjadi disini karena user2 sudah di lock oleh main lain (Transfer) sebelumnya dan tidak di unlock lagi oleh main tersebut sehingga main ini tidak bisa melanjutkan prosesnya (Lock user2) dan terjadi deadlock (main ini tidak bisa di terminate) sehingga program tidak bisa berjalan sampai selesai (tidak bisa mencetak hasil akhir) dan harus di terminate secara paksa (ctrl + c) dan akan muncul error "signal: killed" di terminal (jika dijalankan di terminal) atau "signal: process terminated" di console (jika dijalankan di vscode)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Adi",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Budi",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(10 * time.Second)

	fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.Balance)
}
