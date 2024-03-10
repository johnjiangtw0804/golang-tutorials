package main

// StorageType variable type defines available storage types
// 0 is memory and etc...
type StorageType int

// StorageType option
const Memory StorageType = 0

// type interface for different database to implement
type Storage interface {
	GetUsers() ([]User, error)
	GetUser(id string) (User, error)
	CreateUser(User) error
	DeleteUser(id string) error
}

// Assigning the storage space according to user's input
func NewStorage(storageType StorageType) (Storage, error) {
	var database Storage
	var err error

	switch storageType {
	case Memory:
		database = NewMemStorage()
	}

	return database, err
}
