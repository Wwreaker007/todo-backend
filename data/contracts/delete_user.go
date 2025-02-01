package contracts

type DeleteUserRequest struct{
	UserName 	string `json:"user_name"`
	Password 	string `json:"pass_word"`
}

type DeleteUserResponse struct{
	Status 		string `json:"status"`
}