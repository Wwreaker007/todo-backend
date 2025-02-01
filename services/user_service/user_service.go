package userservice

import (
	"context"
	"todo-backend/data/common"
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

	if user.Status == common.USER_LOGGED_IN {
		return contracts.LoginUserResponse{
			Status: "FAILED : User already logged in",
		}
	}
	user.Status = common.USER_LOGGED_IN
	err = us.UserRepo.UpdateUser(ctx, user)
	if err != nil {
		return contracts.LoginUserResponse{
			Status: "FAILED : " + err.Error(),
		}
	}
	return contracts.LoginUserResponse{
		Status: "SUCCESS",
	}
}

func (us *UserService) Logout(ctx context.Context, r contracts.LogoutUserRequest) contracts.LogoutUserResponse{
	user, err := us.UserRepo.GetUserByUserName(ctx, r.UserName)
	if err != nil {
		return contracts.LogoutUserResponse{
			Status: "FAILED : " + err.Error(),
		}
	}

	// Logout validation login
	if user.Status == common.USER_LOGGED_OUT {
		return contracts.LogoutUserResponse{
			Status: "FAILED : User already logged out",
		}
	}

	// Update the status of the user and update in the DB
	user.Status = common.USER_LOGGED_OUT
	err = us.UserRepo.UpdateUser(ctx, user)
	if err != nil {
		return contracts.LogoutUserResponse{
			Status : "FAILED : " + err.Error(),
		}
	}

	return contracts.LogoutUserResponse{
		Status: "SUCCESS : User sucessfully logged out",
	}
}

func (us *UserService) DeleteUser(ctx context.Context, r contracts.DeleteUserRequest) contracts.DeleteUserResponse {
	user, err := us.UserRepo.GetUserByUserName(ctx, r.UserName)
	if err != nil {
		return contracts.DeleteUserResponse{
			Status: "FAILED : " + err.Error(),
		}
	}

	// Only delete logged in users
	if user.Status == common.USER_LOGGED_OUT {
		return contracts.DeleteUserResponse{
			Status: "FAILED : Please login first",
		}
	}

	// Validation of password before deleting the user
	if user.Password != r.Password {
		return contracts.DeleteUserResponse{
			Status: "FAILED : Password does not match",
		}
	}

	err = us.UserRepo.DeleteUser(ctx, r.UserName)
	if err != nil{
		return contracts.DeleteUserResponse{
			Status: "FAILED : " + err.Error(),
		}
	}
	return contracts.DeleteUserResponse{
		Status: "SUCCESS",
	}
}