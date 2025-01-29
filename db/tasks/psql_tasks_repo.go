package tasks

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"todo-backend/data/common"
)

type TaskRepository struct{
	Client 			*sql.DB
}

const (
	CREATE_TASK_QUERY = "INSERT INTO tasks (user_name, task_name, task_details, task_status) VALUES ($1, $2, $3, $4)"
	GET_TASK_BY_TASK_ID_QUERY = "SELECT * FROM tasks WHERE task_id = $1"
	GET_ALL_TASKS_BY_USER_ID_QUERY = "SELECT * FROM tasks WHERE user_name = $1"
	GET_TASKS_BY_STATUS_QUERY = "SELECT * FROM tasks WHERE user_name = $1 AND task_status = $2"
	UPDATE_TASK_BY_TASK_ID_QUERY = "UPDATE tasks SET task_name = $1, task_details = $2, task_status = $3 WHERE task_id = $4"
	DELETE_TASK_BY_TASK_ID_QUERY = "DELETE FROM tasks WHERE task_id = $1"
)

func NewTaskRepository(client *sql.DB) *TaskRepository{
	return &TaskRepository{
		Client: client,
	}
}

func (tr *TaskRepository) CreateTask(ctx context.Context, data common.TaskData) error {
	_, err := tr.Client.Exec(CREATE_TASK_QUERY, data.UserName, data.TaskName, data.TaskDetails, string(data.TaskStatus))
	if err != nil {
		log.Println("Error while creating task: ", err)
		return err
	}
	return nil
}

func (tr *TaskRepository) GetTaskByTaskID(ctx context.Context, taskID uint64) (common.Task, error) {
	var task common.Task
	err := tr.Client.QueryRow(GET_TASK_BY_TASK_ID_QUERY, taskID).Scan(&task.TaskID, &task.UserName, &task.TaskName, &task.TaskDetails, &task.TaskStatus, &task.CreatedOn, &task.UpdatedOn)
	if err != nil {
		log.Println("Error while fetching task by task id: ", err)
		return common.Task{}, err
	}
	return task, nil
}

func (tr *TaskRepository) GetAllTasksByUserID(ctx context.Context, userName string) ([]common.Task, error) {
	rows, err := tr.Client.Query(GET_ALL_TASKS_BY_USER_ID_QUERY, userName)
	if err != nil {
		log.Println("Error while fetching tasks: ", err)
		return nil, err
	}

	var tasks []common.Task
	for rows.Next() {
		var task common.Task
		err := rows.Scan(&task.TaskID, &task.UserName, &task.TaskName, &task.TaskDetails, &task.TaskStatus, &task.CreatedOn, &task.UpdatedOn)
		if err != nil {
			log.Println("Error while scanning tasks: ", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (tr *TaskRepository) GetTasksByStatus(ctx context.Context, userName string, status common.TaskStatus) ([]common.Task, error) {
	rows, err := tr.Client.Query(GET_TASKS_BY_STATUS_QUERY, userName, status)
	if err != nil {
		log.Println("Error while fetching tasks by status: ", err)
		return nil, err
	}

	if(rows == nil) {
		log.Println("No tasks found with the given status : " + status)
		return nil, errors.New("no matching tasks with the passed status")
	}

	var tasks []common.Task
	for rows.Next() {
		var task common.Task
		err := rows.Scan(&task.TaskID, &task.UserName, &task.TaskName, &task.TaskDetails, &task.TaskStatus, &task.CreatedOn, &task.UpdatedOn)
		if err != nil {
			log.Println("Error while scanning tasks by status: ", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (tr *TaskRepository) UpdateTaskByTaskID(ctx context.Context, task common.Task) error {
	_, err := tr.Client.Exec(UPDATE_TASK_BY_TASK_ID_QUERY, task.TaskName, task.TaskDetails, task.TaskStatus, task.TaskID)
	if err != nil {
		log.Println("Error while updating task: ", err)
		return err
	}
	return nil
}

func (tr *TaskRepository) DeleteTaskByTaskID(ctx context.Context, taskID uint64) error {
	_, err := tr.Client.Exec(DELETE_TASK_BY_TASK_ID_QUERY, taskID)
	if err != nil {
		log.Println("Error while deleting task: ", err)
		return err
	}
	return nil
}