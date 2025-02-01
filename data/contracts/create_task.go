package contracts

type CreateTaskRequest struct{
	UserName 		string `json:"user_name"`
	TaskName 		string `json:"task_name"`
	TaskDetails 	string `json:"task_details"`
}

type CreateTaskResponse struct{
	Status			string 		`json:"status"`
}