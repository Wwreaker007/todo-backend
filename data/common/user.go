package common

type User struct{
	UserData
	CreatedOn 	string `json:"created_on"`
}

type UserData struct{
	UserName 	string `json:"user_name"`
	Password 	string `json:"pass_word"`
	Email 		string `json:"email"`
}