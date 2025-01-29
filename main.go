package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	PATH = ""
	PORT = ":9000"
)

func GetPostgressDBConnector() (*sql.DB, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PSWD"), os.Getenv("DB_NAME"))

	// Get Postgres DB connection
	db, err := sql.Open("postgres", url)
	if(err != nil) {
		fmt.Println("unable to connect to postgresDB : ", err.Error(), url)
		return nil, err
	}

	// Ping to check the connection with the DB
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to the database: " + err.Error())
		return nil, err
	}
	fmt.Println("SUCCESFULLY CONNECTED TO DB")
	return db, nil
}

func main(){
	client, err := GetPostgressDBConnector()
	if err != nil {
		log.Fatalln("Error connecting to the database: ", err)
	}

	// Create a new server
	server := NewServer(PATH, PORT, client)

	// Service startup
	err = server.ServiceStartup()
	if err != nil {
		log.Fatalln("Error starting the server: ", err)
	}
}