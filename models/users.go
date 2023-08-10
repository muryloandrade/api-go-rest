package models

type User struct {
	Id    int    `json:"id"`
	Nome  string `json:"nome"`
	Cargo string `json:"role"`
}

var Users []User
