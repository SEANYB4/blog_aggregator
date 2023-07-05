package main


import (
	"time"
	"github.com/google/uuid"
	"github.com/SEANYB4/blog_aggregator/internal/database"
)


type User struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string
}


func databaseUserToUser(dbUser database.User) User {

	return User {
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
	}
}