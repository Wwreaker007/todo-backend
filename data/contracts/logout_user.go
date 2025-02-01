package contracts

type LogoutUserRequest struct{
	UserName		string			`json:"user_name"`
}

type LogoutUserResponse struct{
	Status			string			`json:"status"`
}