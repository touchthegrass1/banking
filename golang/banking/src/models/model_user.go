package models

import "time"

type User struct {
	Id          int64
	Password    string
	LastLogin   time.Time
	IsSuperuser bool
	FirstName   string
	LastName    string
	IsStaff     bool
	IsActive    bool
	DateJoined  time.Time
	Phone       string
	Email       string
}

func (User) TableName() string {
	return "user"
}
