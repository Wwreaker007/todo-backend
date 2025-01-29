package contracts

type DeleteUserRequest struct{
	UserName 	string `json:"user_name"`
}

type DeleteUserResponse struct{
	Status 		string `json:"status"`
}