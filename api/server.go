package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"sync"
)

type ApiRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Contact struct {
	FirstName string
	LastName  string
	Id        string
}

// This enables us to share data structures across our handlers
type Server struct {
	ContactsDb   []Contact
	ContactsLock sync.Mutex
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
		s.ContactsLock.Lock()
		defer s.ContactsLock.Unlock()
		json.NewEncoder(w).Encode(s.ContactsDb)
	} else if r.Method == http.MethodPost {
		request := ApiRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			fmt.Println(err)
		}

		id := uuid.NewString()
		contact := &Contact{Id: id, FirstName: request.FirstName, LastName: request.LastName}

		s.ContactsLock.Lock()
		defer s.ContactsLock.Unlock()

		s.ContactsDb = append(s.ContactsDb, *contact)
	} else {
		http.Error(w, "generic error", 405)
	}
}
