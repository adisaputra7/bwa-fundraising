package user

import "time"

type User struct {
	ID             int
	Name           string
	Email          string
	Occupation     string
	PasswordHash   string
	Token 		   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt 	   time.Time
}