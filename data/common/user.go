package common

type LOGIN_STATUS string

const (
	USER_LOGGED_IN 	LOGIN_STATUS = "LOGGED_IN"
	USER_LOGGED_OUT LOGIN_STATUS = "LOGGED_OUT"
)

type User struct{
	UserData
	Status		LOGIN_STATUS	`json:"login_status"`
	CreatedOn 	string 			`json:"created_on"`
	UpdatedOn	string 			`json:"updated_on"`
}

type UserData struct{
	UserName 	string 			`json:"user_name"`
	Password 	string 			`json:"pass_word"`
	Email 		string 			`json:"email"`
}