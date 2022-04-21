package main

import "errors"

type Db interface {
	Get(string) (interface{}, error)
	GetAll() []interface{}
	Add(interface{}) bool
}

type Contact struct {
	FirstName string
	LastName  string
	Id        string
}

type ContactsDb struct {
	contacts map[string]Contact
}

// NewContactsDb needs to exist because we have to initialize the map manually
// TODO look up why map isn't initialized the same way a slice is when you attempt to append to it
func NewContactsDb() *ContactsDb {
	return &ContactsDb{
		contacts: make(map[string]Contact),
	}
}

// TODO decide whether or not a pointer is actually better in this situation
func (c *ContactsDb) Get(id string) (Contact, error) {
	if val, ok := c.contacts[id]; ok {
		return val, nil
	}
	return Contact{}, errors.New("id not found")
}

func (c *ContactsDb) GetAll() []Contact {
	var contacts []Contact
	if len(c.contacts) > 0 {
		for _, contact := range c.contacts {
			contacts = append(contacts, contact)
		}
	}
	return contacts
}

// TODO needs error handling
func (c *ContactsDb) Add(contact Contact) bool {
	c.contacts[contact.Id] = contact
	return true
}

func DbHasData(database Db) bool {
	if len(database.GetAll()) > 0 {
		return true
	}
	return false
}
