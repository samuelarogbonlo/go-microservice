package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Welcome To Go Microservice Project")
	})

	http.ListenAndServe(":9090", nil)
}
