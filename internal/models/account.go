package models

import "sync"

type Account struct {
	Mutex   sync.RWMutex
	UserId  string
	Balance float64
}

func (a *Account) GetBalance() float64 {
	a.Mutex.RLock()
	defer a.Mutex.RUnlock()
	return a.Balance
}

func (a *Account) Deposit(amount float64) {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()
	a.Balance += amount
}

func (a *Account) Withdraw(amount float64) bool {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()

	if amount > a.Balance {
		return false
	}
	a.Balance -= amount
	return true
}
