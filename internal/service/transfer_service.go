package service

import (
	"errors"
	"money-transfer/internal/repository"
	"sync"
)

type TransferService struct {
	repo *repository.AccountRepository
	mu   sync.Mutex
}

func NewTransferService(repo *repository.AccountRepository) *TransferService {
	return &TransferService{
		repo: repo,
	}
}

func (ts *TransferService) Transfer(fromUserId, toUserId string, amount float64) error {

	ts.mu.Lock()
	defer ts.mu.Unlock()

	fromAccount, err := ts.repo.GetAccount(fromUserId)
	if err != nil {
		return err
	}

	toAccount, err := ts.repo.GetAccount(toUserId)
	if err != nil {
		return err
	}

	if fromUserId == toUserId {
		return errors.New("self transfer is not allowed")
	}

	if !fromAccount.Withdraw(amount) {
		return errors.New("insufficient balance")
	}

	toAccount.Deposit(amount)
	return nil
}
