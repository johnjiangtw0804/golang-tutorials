package main

import (
	"errors"
	"log"
)

type StorageMemory struct {
	users []User
	ids   map[string]bool
}

func NewMemStorage() *StorageMemory {
	return &StorageMemory{make([]User, 0), make(map[string]bool)}
}

func (s *StorageMemory) GetUsers() ([]User, error) {
	// Get the users
	for _, user := range s.users {
		log.Println("handlers.go: Returned User: ", user.ID, " ", user.AGE)
	}
	if len(s.users) != 0 {
		return s.users, nil
	}

	// Cant get the users
	log.Println("handlers.go: No users in memory storage")
	return s.users, errors.New("no data")
}

func (s *StorageMemory) GetUser(id string) (User, error) {
	// Get the user
	for _, user := range s.users {
		if user.ID == id {
			log.Println("handlers.go: Returned user with id: ", id)
			return user, nil
		}
	}

	// Cant get the user
	log.Println("handlers.go: User with id: ", id, " not found")
	return User{ID: "", AGE: 0}, errors.New("no data")
}

func (s *StorageMemory) CreateUser(user User) error {
	// Cant create the user
	if len(user.ID) == 0 {
		log.Println("handlers.go: User input empty ID: ", user.ID)
		return errors.New("empty string")
	}
	if user.AGE < 0 {
		log.Println("handlers.go: User input negative age: ", user.AGE)
		return errors.New("negative age")
	}
	_, exist := s.ids[user.ID]

	if !exist {
		// Create the user
		s.users = append(s.users, user)
		// Set ID to exist
		s.ids[user.ID] = true
		log.Println("handlers.go: User with id: ", user.ID, " created")
		return nil
	} else {
		log.Println("handlers.go: User input duplicate ID: ", user.ID)
		return errors.New("duplicate ID")
	}
}

func (s *StorageMemory) DeleteUser(id string) error {
	// Delete the user
	for index, user := range s.users {
		if user.ID == id {
			// remove id from the map and user list
			s.users = append(s.users[:index], s.users[index+1:]...)
			delete(s.ids, id)
			log.Println("handlers.go: User with id: ", id, " deleted")
			return nil
		}
	}

	// Cant delete the user
	log.Println("handlers.go: User with id: ", id, " not found")
	return errors.New("no data")
}
