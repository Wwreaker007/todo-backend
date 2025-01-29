package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-backend/data/contracts"
	"todo-backend/types"
)

type DeleteTaskHandler struct{
	tsm 		types.TasksServiceManager
}

func NewDeleteTaskHandler(service types.TasksServiceManager) *DeleteTaskHandler{
	return &DeleteTaskHandler{
		tsm: service,
	}
}

func (h *DeleteTaskHandler) Handle(w http.ResponseWriter, r *http.Request){
	ctx, cancel := context.WithTimeout(r.Context(), 4*time.Second)
	defer cancel()

	var request contracts.DeleteTaskRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return	
	}

	response := h.tsm.DeleteTask(ctx, request)
	json.NewEncoder(w).Encode(response)
}