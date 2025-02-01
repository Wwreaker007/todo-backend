package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-backend/data/contracts"
	"todo-backend/types"
)

type LogoutUserHandler struct{
	usm 		types.UserServiceManager
}

func NewLogoutUserHandler(service types.UserServiceManager) *LogoutUserHandler {
	return &LogoutUserHandler{
		usm : service,
	}
}

func (h *LogoutUserHandler) Handle(w http.ResponseWriter, r *http.Request){
	ctx, cancel := context.WithTimeout(r.Context(), 4*time.Second)
	defer cancel()

	var request contracts.LogoutUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := h.usm.Logout(ctx, request)
	json.NewEncoder(w).Encode(response)
}