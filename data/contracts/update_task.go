package contracts

import "todo-backend/data/common"

type UpdateTaskRequest struct{
	TaskID 		uint64 				`json:"task_id"`
	TaskName 	string 				`json:"task_name"`
	TaskDetails string 				`json:"task_details"`
	TaskStatus 	common.TaskStatus 	`json:"task_status"`
}

type UpdateTaskResponse struct{
	Status 		string 		`json:"status"`
	TaskData 	common.Task `json:"task_data"`
}