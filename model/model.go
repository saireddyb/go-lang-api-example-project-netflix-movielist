package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movies struct {
	_id    primitive.ObjectID
	Name   string
	Rating int
}
