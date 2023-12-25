package models

import "fmt"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UserInput represents the input for creating or updating a user
type UserInput struct {
	Name string `json:"name"`
}

func (user *User) Greet() string {
	return fmt.Sprintf("Hello, I am %s", user.Name)
}
