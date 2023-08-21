package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movies struct {
	Name   string `json:"name"`
	Rating int    `json:"rating"`
	// DirectorID primitive.ObjectID `bson:"directorId"`
}

type Directors struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `json:"name"`
	NationalityID primitive.ObjectID `bson:"nationalityId"`
}

type Countries struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `json:"name"`
}
