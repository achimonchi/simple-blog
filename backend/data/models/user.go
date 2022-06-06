package models

import "time"

type User struct {
	Id        int
	FullName  string
	Email     string
	Password  string
	CreatedAt time.Time
}
