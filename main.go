package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Ticket struct {
	title  string
	points int
	owner  string
}

func main() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/ticket", TicketPostHandler).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func TicketPostHandler(rw http.ResponseWriter, req *http.Request) {

}

func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Hello World!")
}
