package models

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type UserLocation struct {
	Latitude  float64
	Longitude float64
}

type User struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	Age       int
	Gender    string
	Location  UserLocation `gorm:"embedded"`
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateRandomUser() User {
	// roughly within Greater London/M25
	lat, _ := gofakeit.LatitudeInRange(51.245, 51.759)
	long, _ := gofakeit.LongitudeInRange(-0.302, 0.285)
	location := UserLocation{Latitude: lat, Longitude: long}
	user := User{
		Name:     gofakeit.Name(),
		Age:      gofakeit.Number(18, 100),
		Gender:   gofakeit.Gender(),
		Location: location,
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, false, false, false, false, 12),
	}
	return user
}
