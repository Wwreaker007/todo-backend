package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-backend/data/contracts"
	"todo-backend/types"
)

type GetAllTaskHandler struct{
	tsm 		types.TasksServiceManager
}

func NewGetAllTaskHandler(service types.TasksServiceManager) *GetAllTaskHandler{
	return &GetAllTaskHandler{
		tsm: service,
	}
}

func (h *GetAllTaskHandler) Handle(w http.ResponseWriter, r *http.Request){
	ctx, cancel := context.WithTimeout(r.Context(), 4*time.Second)
	defer cancel()

	var request contracts.GetAllTaskRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := h.tsm.GetAllTasks(ctx, request)
	json.NewEncoder(w).Encode(response)
}