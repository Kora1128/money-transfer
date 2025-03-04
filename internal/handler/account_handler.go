package handler

import (
	"encoding/json"
	"money-transfer/internal/repository"
	"net/http"
)

type AccountHandler struct {
	accountRepo *repository.AccountRepository
}

func NewAccountHandler(accountRepo *repository.AccountRepository) *AccountHandler {
	return &AccountHandler{
		accountRepo: accountRepo,
	}
}

type AccountRequest struct {
	UserId  string  `json:"user_id"`
	Balance float64 `json:"balance"`
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req AccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	account, err := h.accountRepo.CreateAccount(req.UserId, req.Balance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"user_id": account.UserId, "status": "account created"})
}

func (h *AccountHandler) GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	if userId == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}
	account, err := h.accountRepo.GetAccount(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]float64{"balance": account.GetBalance()})
}
