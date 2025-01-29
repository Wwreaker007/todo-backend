package contracts

type LoginUserRequest struct{
	UserName 	string `json:"user_name"`
	Password 	string `json:"pass_word"`
}

type LoginUserResponse struct{
	Status 		string `json:"status"`
}