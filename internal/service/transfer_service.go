package service

import (
	"errors"
	"money-transfer/internal/repository"
)

type TransferService struct {
	repo *repository.AccountRepository
}

func NewTransferService(repo *repository.AccountRepository) *TransferService {
	return &TransferService{
		repo: repo,
	}
}

func (ts *TransferService) Transfer(fromUserId, toUserId string, amount float64) error {

	fromAccount, err := ts.repo.GetAccount(fromUserId)
	if err != nil {
		return err
	}

	toAccount, err := ts.repo.GetAccount(toUserId)
	if err != nil {
		return err
	}

	// To Prevent Deadlock
	firstLock, secondLock := fromAccount, toAccount
	if firstLock.UserId > secondLock.UserId {
		firstLock, secondLock = secondLock, firstLock
	}

	firstLock.Mutex.Lock()
	secondLock.Mutex.Lock()
	defer firstLock.Mutex.Unlock()
	defer secondLock.Mutex.Unlock()

	if fromUserId == toUserId {
		return errors.New("self transfer is not allowed")
	}

	if !fromAccount.Withdraw(amount) {
		return errors.New("insufficient balance")
	}

	toAccount.Deposit(amount)
	return nil
}
