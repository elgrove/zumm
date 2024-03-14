package models

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	Age       int
	Gender    string
	Location  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateRandomUser() User {
	user := User{
		Name:     gofakeit.Name(),
		Age:      gofakeit.Number(18, 100),
		Gender:   gofakeit.Gender(),
		Location: fmt.Sprintf("%f,%f", gofakeit.Latitude(), gofakeit.Longitude()),
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, false, false, false, false, 12),
	}
	return user
}

type UserLogin struct {
	Email    string
	Password string
}

type TokenResponse struct {
	Token string
}
