package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-backend/data/contracts"
	"todo-backend/types"
)

type GetTaskByStatusHandler struct{
	tsm 		types.TasksServiceManager
}

func NewGetTaskByStatusHandler(service types.TasksServiceManager) *GetTaskByStatusHandler{
	return &GetTaskByStatusHandler{
		tsm: service,
	}
}

func (h *GetTaskByStatusHandler) Handle(w http.ResponseWriter, r *http.Request){
	ctx, cancel := context.WithTimeout(r.Context(), 4*time.Second)
	defer cancel()

	var request contracts.GetTaskStatusRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := h.tsm.GetTasksByStatus(ctx, request)
	json.NewEncoder(w).Encode(response)
}