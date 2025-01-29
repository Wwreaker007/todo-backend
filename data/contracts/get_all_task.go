package contracts

import "todo-backend/data/common"

type GetAllTaskRequest struct{
	UserName 		string 			`json:"user_name"`
}

type GetAllTaskResponse struct{
	Status			string 			`json:"status"`
	Tasks 			[]common.Task 	`json:"tasks"`
}