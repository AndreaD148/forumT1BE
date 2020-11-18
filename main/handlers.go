package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

//da aggiungere strconv

/*
type response struct {
	ID        uint64 `json:id, omitempty`
	Message   string `json: message`
}
*/

type psw struct {
	Password string `json: "password"`
}


var db = connect()

//function to create a new User from the request
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user User
	var searchedUser User

	/*body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err.Error())
	}*/

	//fmt.Println(r.Body)
	//fmt.Println(string(body))

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		panic(err.Error())
	}

	//ora che ho le informazioni sull'utente posso stabilire se è stato già aggiunto oppure no

	db.Where("username = ?", user.Username).First(&searchedUser)

	fmt.Println(searchedUser)

	if user.Username == searchedUser.Username {
		fmt.Fprintln(w, "Esiste già un utente con questo username")
	} else {
		encryptedPassword, err := EncryptPassword(user.Password)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(encryptedPassword)

		user.Password = encryptedPassword

		//fmt.Println(user.Password)
		fmt.Fprintln(w, user)

		db.Create(&user)
	}



	//fmt.Fprintln(w, user.ID)
}

func sayHello(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Fprintln(w,"Hello. The server is listening correctly")
}


func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	return string(bytes), err
}


func testEncryption(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var userPassword psw

	err := json.NewDecoder(r.Body).Decode(&userPassword)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintln(w, userPassword)


	fmt.Println(userPassword.Password)

	//try to Encrypt
	bytes, err := bcrypt.GenerateFromPassword([]byte(userPassword.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	//fmt.Println(string(bytes))

	userPassword.Password = string(bytes)

	fmt.Println(userPassword.Password)


	/*
	err := json.NewDecoder(r.Body).Decode(&userPassword)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(userPassword.password)

	fmt.Fprintln(w, userPassword)
	*/


}
