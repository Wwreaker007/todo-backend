package contracts

import "todo-backend/data/common"


type CreateUserRequest struct{
	UserData 	common.UserData 	`json:"user_data"`
}

type CreateUserResponse struct{
	Status 		string 		`json:"status"`
}