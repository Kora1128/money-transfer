package repository

import (
	"errors"
	"money-transfer/internal/models"
	"sync"
)

type AccountRepository struct {
	mu       sync.RWMutex
	accounts map[string]*models.Account
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{
		accounts: make(map[string]*models.Account),
	}
}

func (ar *AccountRepository) GetAccount(userId string) (*models.Account, error) {
	ar.mu.RLock()
	defer ar.mu.RUnlock()
	account, exists := ar.accounts[userId]
	if !exists {
		return nil, errors.New("account not found for user: " + userId)
	}
	return account, nil
}

func (ar *AccountRepository) CreateAccount(userId string, initialBalance float64) (*models.Account, error) {
	ar.mu.Lock()
	defer ar.mu.Unlock()
	if _, exists := ar.accounts[userId]; exists {
		return nil, errors.New("account already exists for user: " + userId)
	}

	if initialBalance < 0 {
		return nil, errors.New("initial balance cannot be negative")
	}
	account := &models.Account{
		UserId:  userId,
		Balance: initialBalance,
	}
	ar.accounts[userId] = account
	return account, nil
}
