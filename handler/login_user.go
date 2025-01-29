package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-backend/data/contracts"
	"todo-backend/types"
)

type LoginUserHandler struct{
	usm 		types.UserServiceManager
}

func NewLoginUserHandler(service types.UserServiceManager) *LoginUserHandler{
	return &LoginUserHandler{
		usm: service,
	}
}

func (h *LoginUserHandler) Handle(w http.ResponseWriter, r *http.Request){
	ctx, cancel := context.WithTimeout(r.Context(), 4*time.Second)
	defer cancel()

	var request contracts.LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return	
	}

	response := h.usm.Login(ctx, request)
	json.NewEncoder(w).Encode(response)
}