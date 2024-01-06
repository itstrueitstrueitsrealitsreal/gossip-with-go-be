package models

import "fmt"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// UserInput represents the input for creating or updating a user
type UserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (user *User) Greet() string {
	return fmt.Sprintf("Hello, I am %s", user.Name)
}
