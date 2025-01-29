package userservice

import (
	"context"
	"todo-backend/data/contracts"
	"todo-backend/types"
)

type UserService struct{
	UserRepo 	types.UserReposioryManager
}

func NewUserService(userRepo types.UserReposioryManager) *UserService{
	return &UserService{
		UserRepo: userRepo,
	}
}

func (us *UserService) SingUp(ctx context.Context, r contracts.CreateUserRequest) contracts.CreateUserResponse {
	err := us.UserRepo.CreateUser(ctx, r.UserData)
	if err != nil{
		return contracts.CreateUserResponse{
			Status: "FAILED : " + err.Error(),
		}
	}
	return contracts.CreateUserResponse{
		Status: "SUCCESS",
	}
}

func (us *UserService) Login(ctx context.Context, r contracts.LoginUserRequest) contracts.LoginUserResponse {
	user, err := us.UserRepo.GetUserByUserName(ctx, r.UserName)
	if err != nil{
		return contracts.LoginUserResponse{
			Status: "FAILED : " + err.Error(),
		}
	}

	// Password validation logic
	if user.Password != r.Password{
		return contracts.LoginUserResponse{
			Status: "FAILED : Password does not match",
		}
	}
	return contracts.LoginUserResponse{
		Status: "SUCCESS",
	}
}

func (us *UserService) DeleteUser(ctx context.Context, r contracts.DeleteUserRequest) contracts.DeleteUserResponse {
	err := us.UserRepo.DeleteUser(ctx, r.UserName)
	if err != nil{
		return contracts.DeleteUserResponse{
			Status: "FAILED : " + err.Error(),
		}
	}
	return contracts.DeleteUserResponse{
		Status: "SUCCESS",
	}
}