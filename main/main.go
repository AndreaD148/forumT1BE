package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
)







func main() {
	//connect()
	r := Router()

	fmt.Println("Starting server on port 5000...")

	log.Fatal(http.ListenAndServe(":5000", r))
}
