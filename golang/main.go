package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	healthResponse = "{ status : up }"
)

func simpleHandler(res http.ResponseWriter, req *http.Request) {
	_, err := res.Write([]byte(healthResponse))
	if err != nil {
		fmt.Println("Failed to write response. Error : ", err)
	}
}

func main() {
	http.HandleFunc("/health", simpleHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start server. Error : ", err)
	}
}
