package users

import (
	"context"
	"database/sql"
	"log"
	"todo-backend/data/common"
)

type UserReposiory struct{
	Client 		*sql.DB
}

const (
	CREATE_USER_QUERY = "INSERT INTO users (user_name, user_password, login_status, email) VALUES ($1, $2, $3, $4)"
	GET_USER_BY_USERID_QUERY = "SELECT * FROM users WHERE user_name = $1"
	UPDATE_USER_BY_USERID = "UPDATE users SET user_password = $1, login_status = $2, email = $3 WHERE user_name = $4"
	DELETE_USER_QUERY = "DELETE FROM users WHERE user_name = $1"
)

func NewUserRepository(client *sql.DB) *UserReposiory{
	return &UserReposiory{
		Client: client,
	}
}

func (ur *UserReposiory) CreateUser(ctx context.Context, data common.UserData) error {
	_, err := ur.Client.Exec(CREATE_USER_QUERY, data.UserName, data.Password, common.USER_LOGGED_OUT, data.Email)
	if err != nil {
		log.Println("Error while creating user: ", err)
		return err
	}
	return nil
}

func (ur *UserReposiory) GetUserByUserName(ctx context.Context, username string) (common.User, error) {
	var user common.User
	err := ur.Client.QueryRow(GET_USER_BY_USERID_QUERY, username).Scan(&user.UserName, &user.Password, &user.Status, &user.Email, &user.CreatedOn, &user.UpdatedOn)
	if err != nil{
		log.Println("Error while fetching user: ", err)
		return common.User{}, err
	}
	return user, nil
}

func (us *UserReposiory) UpdateUser(ctx context.Context, data common.User) error {
	_, err := us.Client.Exec(UPDATE_USER_BY_USERID, data.Password, data.Status, data.Email, data.UserName)
	if err != nil {
		log.Println("Error while updating user : ", err)
		return err
	}
	return nil
}

func (ur *UserReposiory) DeleteUser(ctx context.Context, username string) error {
	_, err := ur.Client.Exec(DELETE_USER_QUERY, username)
	if err != nil {
		log.Println("Error while deleting user: ", err)
		return err
	}
	return nil
}