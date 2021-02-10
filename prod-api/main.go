package main

import (
	"log"
	"net/http"
	"os"

	"github.com/samuelarogbonlo/Go-Microservice/prod-api/handlers"
)

//function handler and path("") displays input
func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l) //reference to the new handler
	gh := handlers.Goodbye(l)

	//new servemux and implements handler interface
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{}

	// Listen for connections on all ip addresses (0.0.0.0)
	// port 9090
	log.Println("Starting Server")
	err := http.ListenAndServe(":9090", nil)
	log.Fatal(err)
}
