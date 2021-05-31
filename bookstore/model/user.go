package model

import "fmt"

type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

func (user *User) String() string {
	return fmt.Sprintln("ID:", user.ID, ",Username:", user.Username,
		",Password:", user.Password, ",Email:", user.Email)
}
