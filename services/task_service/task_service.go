package taskservice

import (
	"context"
	"todo-backend/data/common"
	"todo-backend/data/contracts"
	"todo-backend/types"
)

type TaskService struct{
	TaskRepo 	types.TasksRepositoryManager
}

func NewTaskService(taskRepo types.TasksRepositoryManager) *TaskService{
	return &TaskService{
		TaskRepo: taskRepo,
	}
}

func (ts *TaskService) CreateTask(ctx context.Context, r contracts.CreateTaskRequest) contracts.CreateTaskResponse {
	newTaskData := common.TaskData{
		UserName: r.UserName,
		TaskName: r.TaskName,
		TaskDetails: r.TaskDetails,
		TaskStatus: common.Pending,
	}
	err := ts.TaskRepo.CreateTask(ctx, newTaskData)
	if err != nil{
		return contracts.CreateTaskResponse{
			Status: "FAILED : " + err.Error(),
		}
	}
	return contracts.CreateTaskResponse{
		Status: "SUCCESS",
	}
}

func (ts *TaskService) GetAllTasks(ctx context.Context, r contracts.GetAllTaskRequest) contracts.GetAllTaskResponse {
	tasks, err := ts.TaskRepo.GetAllTasksByUserID(ctx, r.UserName)
	if err != nil {
		return contracts.GetAllTaskResponse{
			Status: "FAILED : " + err.Error(),
		}
	}
	return contracts.GetAllTaskResponse{
		Tasks: tasks,
		Status: "SUCCESS",
	}
}

func (ts *TaskService) GetTasksByStatus(ctx context.Context, r contracts.GetTaskStatusRequest) contracts.GetTaskStatusResponse {
	tasks, err := ts.TaskRepo.GetTasksByStatus(ctx, r.UserName, r.Status)
	if err != nil {
		return contracts.GetTaskStatusResponse{
			Status: "FAILED : " + err.Error(),
		}
	}
	return contracts.GetTaskStatusResponse{
		Status: "SUCCESS",
		Tasks: tasks,
	}
}

func (ts *TaskService) UpdateTask(ctx context.Context, r contracts.UpdateTaskRequest) contracts.UpdateTaskResponse {
	task, err := ts.TaskRepo.GetTaskByTaskID(ctx, r.TaskID)
	if err != nil {
		return contracts.UpdateTaskResponse{
			Status: "FAILED : " + err.Error(),
		}
	}

	// Update the details of the task
	task.TaskName = r.TaskName
	task.TaskDetails = r.TaskDetails
	task.TaskStatus = r.TaskStatus

	err = ts.TaskRepo.UpdateTaskByTaskID(ctx, task)
	if err != nil{
		return contracts.UpdateTaskResponse{
			Status: "FAILED : " + err.Error(),
		}
	}
	return contracts.UpdateTaskResponse{
		Status: "SUCCESS",
		TaskData: task,
	}
}

func (ts *TaskService) DeleteTask(ctx context.Context, r contracts.DeleteTaskRequest) contracts.DeleteTaskResponse {
	err := ts.TaskRepo.DeleteTaskByTaskID(ctx, r.TaskID)
	if err != nil {
		return contracts.DeleteTaskResponse{
			Status: "FAILED : " + err.Error(),
		}
	}
	return contracts.DeleteTaskResponse{
		Status: "SUCCESS",
	}
}