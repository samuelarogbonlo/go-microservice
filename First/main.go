package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//function handler and path("") displays input
func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Building For Next Billion Users")
		// read the body
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading body", err)

			http.Error(rw, "Unable to read request body", http.StatusBadRequest)
			return
		}
		// write the response
		fmt.Fprintf(rw, "Hello %s", b)
	})

	// any other request will be handled by this function
	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World - Running Hello Handler")

	})

	// Listen for connections on all ip addresses (0.0.0.0)
	// port 9090
	log.Println("Starting Server")
	err := http.ListenAndServe(":9090", nil)
	log.Fatal(err)
}
