package models

import (
	"errors"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (i *User) Validate() map[string]string {
	var err error
	var errormessage = make(map[string]string)
	if i.Username == "" {
		err = errors.New("Required username")
		errormessage["Required_username"] = err.Error()
	}
	if i.Email == "" {
		err = errors.New("Required email")
		errormessage["Required_email"] = err.Error()
	}
	if i.Password == "" {
		err = errors.New("Required password")
		errormessage["Required_password"] = err.Error()
	}
	return errormessage
}
