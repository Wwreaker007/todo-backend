package contracts

import "todo-backend/data/common"

type GetTaskStatusRequest struct{
	UserName 		string 				`json:"user_name"`
	Status 			common.TaskStatus 	`json:"task_status"`
}

type GetTaskStatusResponse struct{
	Status			string 				`json:"status"`
	Tasks 			[]common.Task 		`json:"tasks"`
}