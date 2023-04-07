package models

import (
	"errors"
	"time"
)

type Todo struct {
	ID           string    `json:"id" bson:"_id"`
	Title        string    `bson:"title" json:"title"`
	Discpription string    `bson:"description" json:"description"`
	CreatedAt    time.Time `bson:"created_at" json:"created_at"`
}

func (i *Todo) Validate() map[string]string {
	var err error
	var errormessage = make(map[string]string)
	if i.Title == "" {
		err = errors.New("Required title")
		errormessage["Required_title"] = err.Error()
	}
	if i.Discpription == "" {
		err = errors.New("Required description")
		errormessage["Required_description"] = err.Error()
	}
	return errormessage
}
