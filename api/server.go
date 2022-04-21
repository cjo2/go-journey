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

// This enables us to share data structures across our handlers
type Server struct {
	ContactsDb   *ContactsDb
	ContactsLock sync.Mutex
}

func main() {
	server := &Server{}
	server.ContactsDb = NewContactsDb()
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
		contacts := s.ContactsDb.GetAll()
		json.NewEncoder(w).Encode(contacts)
	} else if r.Method == http.MethodPost {
		request := ApiRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			fmt.Println(err)
		}

		id := uuid.NewString()
		fmt.Println(request)
		contact := Contact{Id: id, FirstName: request.FirstName, LastName: request.LastName}
		s.ContactsDb.Add(contact)
	} else {
		http.Error(w, "generic error", 405)
	}
}
