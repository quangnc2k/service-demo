package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Staff struct {
	ID          bson.ObjectId `json:"id"`
	CreatedAt   time.Time     `json:"created_at"`
	Name        string        `json:"name"`
	DateOfBirth time.Time     `json:"date_of_birth"`
	Address     string        `json:"address"`
	Teams       []Team        `json:"teams"`
}
