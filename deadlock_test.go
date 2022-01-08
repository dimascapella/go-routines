package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	Mutex   sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) UnLock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(user1, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock", user1.Name)
	user1.Change(-amount)

	time.Sleep(2 * time.Second)

	user2.Lock()
	fmt.Println("Lock", user2.Name)
	user2.Change(amount)

	time.Sleep(2 * time.Second)

	user1.UnLock()
	user2.UnLock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Dimas",
		Balance: 10000,
	}

	user2 := UserBalance{
		Name:    "Eka",
		Balance: 10000,
	}

	go Transfer(&user1, &user2, 500)
	go Transfer(&user2, &user1, 700)

	time.Sleep(10 * time.Second)

	fmt.Println(user1.Name, "Has Balance ", user1.Balance)
	fmt.Println(user2.Name, "Has Balance ", user2.Balance)
}
