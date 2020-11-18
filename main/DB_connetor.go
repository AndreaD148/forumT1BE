package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func connect() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Some error occurred! " + err.Error())
	}

	databaseName := os.Getenv("DATABASE_NAME")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	hostname := os.Getenv("HOST_NAME")


	dbUri := fmt.Sprintf("host=%s port=5432 user=%s dbname=%s sslmode=disable password=%s",
		hostname, username, databaseName, password) //Build connection string
	fmt.Println(dbUri)

	db, err := gorm.Open("postgres", dbUri)

	if err != nil {
		log.Panic(err.Error())
	}

	//defer db.Close()
	db.AutoMigrate(&User{})

	return db

}
