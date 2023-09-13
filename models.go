package main

import (
	"github.com/ChukwunonsoFrank/hngx-task-two/internal/database"
	"github.com/google/uuid"
)

type Person struct {
	ID        uuid.UUID `json:"id"`
	Name      string `json:"name"`
}

func databaseUserToUser(dbUser database.Person) Person {
	return Person{
		ID: dbUser.ID,
		Name: dbUser.Name,
	}
}
