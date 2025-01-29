package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-backend/data/contracts"
	"todo-backend/types"
)

type CreateTaskHanlder struct{
	tsm  	types.TasksServiceManager
}

func NewCreateTaskHanlder(service types.TasksServiceManager) *CreateTaskHanlder{
	return &CreateTaskHanlder{
		tsm: service,
	}
}

func (h *CreateTaskHanlder) Handle(w http.ResponseWriter, r *http.Request){
	ctx, cancel := context.WithTimeout(r.Context(), 4*time.Second)
	defer cancel()

	var request contracts.CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return	
	}

	response := h.tsm.CreateTask(ctx, request)
	json.NewEncoder(w).Encode(response)
}