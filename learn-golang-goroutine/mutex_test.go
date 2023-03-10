package hello

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() (balance int) {
	account.RWMutex.RLock()
	balance = account.Balance
	account.RWMutex.RUnlock()

	return
}

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("jumlah", x)
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("jumlah")
}
