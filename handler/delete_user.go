package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-backend/data/contracts"
	"todo-backend/types"
)

type DeleteUserHandler struct{
	usm 		types.UserServiceManager
}

func NewDeleteUserHandler(service types.UserServiceManager) *DeleteUserHandler{
	return &DeleteUserHandler{
		usm: service,
	}
}

func (h *DeleteUserHandler) Handle(w http.ResponseWriter, r *http.Request){
	ctx, cancel := context.WithTimeout(r.Context(), 4*time.Second)
	defer cancel()

	var request contracts.DeleteUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return	
	}

	response := h.usm.DeleteUser(ctx, request)
	json.NewEncoder(w).Encode(response)
}