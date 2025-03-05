package service

import (
	"errors"
	"log"
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

	log.Println("Transferring", amount, "from", fromUserId, "to", toUserId)

	// To Prevent Deadlock
	firstLock, secondLock := fromAccount, toAccount
	if firstLock.UserId > secondLock.UserId {
		firstLock, secondLock = secondLock, firstLock
	}

	log.Println("acquiring lock for", firstLock.UserId)
	firstLock.Mutex.Lock()

	log.Println("acquiring lock for", secondLock.UserId)
	secondLock.Mutex.Lock()

	defer firstLock.Mutex.Unlock()
	defer secondLock.Mutex.Unlock()

	log.Println("locks acquired")

	if fromUserId == toUserId {
		return errors.New("self transfer is not allowed")
	}

	if !fromAccount.Withdraw(amount) {
		return errors.New("insufficient balance")
	}
	log.Println("withdrawn from", fromUserId)

	toAccount.Deposit(amount)
	log.Println("deposited to", toUserId)

	return nil
}
