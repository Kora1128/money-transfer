package models

import "sync"

type Account struct {
	mu      sync.RWMutex
	UserId  string
	Balance float64
}

func (a *Account) GetBalance() float64 {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.Balance
}

func (a *Account) Deposit(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
}

func (a *Account) Withdraw(amount float64) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	if amount > a.Balance {
		return false
	}
	a.Balance -= amount
	return true
}
