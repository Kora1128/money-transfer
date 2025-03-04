package handler

import (
	"encoding/json"
	"money-transfer/internal/service"
	"net/http"
)

type TransferHandler struct {
	transferService *service.TransferService
}

func NewTransferHandler(transferService *service.TransferService) *TransferHandler {
	return &TransferHandler{
		transferService: transferService,
	}
}

type TransferRequest struct {
	FromUserId string  `json:"from_user_id"`
	ToUserId   string  `json:"to_user_id"`
	Amount     float64 `json:"amount"`
}

func (h *TransferHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	var req TransferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.FromUserId == "" || req.ToUserId == "" {
		http.Error(w, "from_user_id and to_user_id are required", http.StatusBadRequest)
		return
	}
	if req.Amount <= 0 {
		http.Error(w, "amount must be greater than 0", http.StatusBadRequest)
		return
	}
	if err := h.transferService.Transfer(req.FromUserId, req.ToUserId, req.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "transfer successful"})
}
