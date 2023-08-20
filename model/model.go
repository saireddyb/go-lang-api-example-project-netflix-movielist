package model

type Movies struct {
	Id       int
	Name     string
	Rating   int
	Director *Director
}

type Director struct {
	Id          int
	Name        string
	Nationality Nationality
}

type Nationality struct {
	Country string
}
