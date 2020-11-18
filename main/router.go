package main

import "github.com/gorilla/mux"

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/test", CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/", sayHello).Methods("GET", "OPTIONS")
	router.HandleFunc("/testEnc", testEncryption).Methods("GET", "OPTIONS")

	return router

}
