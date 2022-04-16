package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Contact struct {
	FirstName string
	LastName  string
}

// This enables us to share data structures across our handlers
type Server struct {
	ContactsDb []Contact
}

func main() {
	server := &Server{}
	http.HandleFunc("/", server.ParentHandler)
	http.HandleFunc("/contact", server.ContactHandler)

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println(err)
	}
}

func (s *Server) ParentHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "Ack")
}

func (s *Server) ContactHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(s.ContactsDb)
	} else if r.Method == http.MethodPost {
		request := ApiRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			fmt.Println(err)
		}
		s.ContactsDb = append(s.ContactsDb, Contact{FirstName: request.FirstName, LastName: request.LastName})
	} else {
		http.Error(w, "generic error", 405)
	}
}
