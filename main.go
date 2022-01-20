package main

// https://tutorialedge.net/golang/creating-restful-api-with-golang/
// https://gobyexample.com/json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Fields need to be exported, by capitalizing, in order for json module to access
// json tags used for mapping json obj field to struct field
type Ticket struct {
	Title  string `json:"title"`
	Points int    `json:"points"`
	Owner  string `json:"owner"`
}

func main() {
	fmt.Println("starting simpleapi")

	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/ticket", TicketPostHandler).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func TicketPostHandler(rw http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	fmt.Printf("Handling ticket POST\n%+v\n", string(reqBody))

	var ticket Ticket
	err := json.Unmarshal(reqBody, &ticket)

	if err != nil {
		fmt.Println("Error parsing json object:", err)
		return
	}

	fmt.Println(ticket)
}

func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Hello World!")
}
