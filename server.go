package main

import (
	"database/sql"
	"net/http"
	"todo-backend/db/tasks"
	"todo-backend/db/users"
	"todo-backend/handler"
	taskservice "todo-backend/services/task_service"
	userservice "todo-backend/services/user_service"
)

const (
	CREATE_USER = "/create/user"
	LOGIN_USER = "/login/user"
	DELETE_USER = "/delete/user"
	LOGOUT_USER = "/logout/user"

	CREATE_TASK = "/create/task"
	GET_ALL_TASKS = "/get/tasks/all"
	GET_TASKS_BY_STATUS = "/get/tasks/status"
	UPDATE_TASK = "/update/task"
	DELETE_TASK = "/delete/task"
)

type Server struct{
	Path 		string
	Port 		string
	DbClient 		*sql.DB
}

func NewServer(path string, port string, client *sql.DB) *Server{
	return &Server{
		Path: path,
		Port: port,
		DbClient: client,
	}
}

func (s *Server) ServiceStartup() error{
	serveMux := http.NewServeMux()

	// Assign the dependencies to be passed to the handler
	userRepo := users.NewUserRepository(s.DbClient)
	taskRepo := tasks.NewTaskRepository(s.DbClient)
	userService := userservice.NewUserService(userRepo)
	taskService := taskservice.NewTaskService(taskRepo)

	// Registering the handler with the serveMux (User Service)
	serveMux.HandleFunc(CREATE_USER, handler.NewCreateUserHandler(userService).Handle)
	serveMux.HandleFunc(LOGIN_USER, handler.NewLoginUserHandler(userService).Handle)
	serveMux.HandleFunc(DELETE_USER, handler.NewDeleteUserHandler(userService).Handle)
	serveMux.HandleFunc(LOGOUT_USER, handler.NewLogoutUserHandler(userService).Handle)

	// Registering the handler with the serveMux (Task Service)
	serveMux.HandleFunc(CREATE_TASK, handler.NewCreateTaskHanlder(taskService).Handle)
	serveMux.HandleFunc(GET_ALL_TASKS, handler.NewGetAllTaskHandler(taskService).Handle)
	serveMux.HandleFunc(GET_TASKS_BY_STATUS, handler.NewGetTaskByStatusHandler(taskService).Handle)
	serveMux.HandleFunc(UPDATE_TASK, handler.NewUpdateTaskHandler(taskService).Handle)
	serveMux.HandleFunc(DELETE_TASK, handler.NewDeleteTaskHandler(taskService).Handle)

	// Start the server
	return http.ListenAndServe(s.Path + s.Port, serveMux)
}