package types

import (
	"context"
	"todo-backend/data/common"
	"todo-backend/data/contracts"
)

type UserReposioryManager interface {
	CreateUser(context.Context, common.UserData) error
	GetUserByUserName(context.Context, string) (common.User, error)
	DeleteUser(context.Context, string) error
}

type TasksRepositoryManager interface {
	CreateTask(context.Context, common.TaskData) error
	GetTaskByTaskID(context.Context, uint64) (common.Task, error)
	GetAllTasksByUserID(context.Context, string) ([]common.Task, error)
	GetTasksByStatus(context.Context, string, common.TaskStatus) ([]common.Task, error)
	UpdateTaskByTaskID(context.Context, common.Task) error
	DeleteTaskByTaskID(context.Context, uint64) error
}

type UserServiceManager interface{
	SingUp(context.Context, contracts.CreateUserRequest) contracts.CreateUserResponse
	Login(context.Context, contracts.LoginUserRequest) contracts.LoginUserResponse
	DeleteUser(context.Context, contracts.DeleteUserRequest) contracts.DeleteUserResponse
}

type TasksServiceManager interface{
	CreateTask(context.Context, contracts.CreateTaskRequest) contracts.CreateTaskResponse
	GetAllTasks(context.Context, contracts.GetAllTaskRequest) contracts.GetAllTaskResponse
	GetTasksByStatus(context.Context, contracts.GetTaskStatusRequest) contracts.GetTaskStatusResponse
	UpdateTask(context.Context, contracts.UpdateTaskRequest) contracts.UpdateTaskResponse
	DeleteTask(context.Context, contracts.DeleteTaskRequest) contracts.DeleteTaskResponse
}