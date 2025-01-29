package contracts

type DeleteTaskRequest struct{
	TaskID 		uint64 `json:"task_id"`
}

type DeleteTaskResponse struct{
	Status 		string `json:"status"`
}