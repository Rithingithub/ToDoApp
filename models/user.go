package models

import (
	"errors"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       string `json:"id" bson:"_id"`
	Username string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	//Todo     []primitive.ObjectID `bson:"todo,omitempty" json:"todo,omitempty"`
}

// type Todo struct {
// 	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
// 	Title string             `bson:"title" json:"title"`
// 	// Date  time.Time          `bson:"date" json:"date"`
// }

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
