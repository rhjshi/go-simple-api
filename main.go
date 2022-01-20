package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Ticket struct {
	title  string
	points int
	owner  string
}

func main() {
	fmt.Println("Hello World!")

	r := mux.NewRouter()

	r.HandleFunc("/ticket", TicketPostHandler).Methods("POST")
}

func TicketPostHandler(rw http.ResponseWriter, req *http.Request) {

}
