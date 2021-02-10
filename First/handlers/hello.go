package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

//hello handler into its own struct
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

//add interface method - satisfies http handler interfcae
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Building For Next Billion Users")
	// read the body
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops Unable to read request body", http.StatusBadRequest)
		return
	}
	// write the response
	fmt.Fprintf(rw, "Hello %s", d)
}
