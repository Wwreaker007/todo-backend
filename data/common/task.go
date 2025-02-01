package common

type TaskStatus string

const (
	Pending TaskStatus = "PENDING"
	InProgress TaskStatus = "IN_PROGRESS"
	Completed TaskStatus = "COMPLETED"
)

type Task struct{
	TaskData
	TaskID			uint64		`json:"task_id"`
	CreatedOn		string		`json:"created_on"`
	UpdatedOn		string		`json:"updated_on"`		
}

type TaskData struct{
	UserName		string		`json:"user_name"`
	TaskName		string		`json:"task_name"`
	TaskDetails		string		`json:"task_details"`
	TaskStatus		TaskStatus	`json:"task_status"`
}