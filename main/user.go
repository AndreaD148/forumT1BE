package main

import (
	"github.com/google/uuid"
)


type User struct {
	//gorm.Model
	Username string    `json: "username"`
	Password string    `json: "password"`
	City     string    `json: "city"`
	Email    string    `json: "email"`
	ID   uuid.UUID     `gorm:"primary_key; type: uuid; default: uuid_generate_v4()"`

}

//DA RICORDARE: se voglio mettere una nuova tabella nel DB devo mettere la prima lettera maiuscola (stessa cosa per i campi)