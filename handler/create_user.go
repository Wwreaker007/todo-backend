package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-backend/data/contracts"
	"todo-backend/types"
)

type CreateUserHandler struct{
	usm 		types.UserServiceManager
}

func NewCreateUserHandler(service types.UserServiceManager) *CreateUserHandler{
	return &CreateUserHandler{
		usm: service,
	}
}

func (h *CreateUserHandler) Handle(w http.ResponseWriter, r *http.Request){
	ctx, cancel := context.WithTimeout(r.Context(), 4*time.Second)
	defer cancel()

	var request contracts.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return	
	}

	response := h.usm.SingUp(ctx, request)
	json.NewEncoder(w).Encode(response)
}