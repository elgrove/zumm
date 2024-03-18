package model

import (
	"time"
)

// UserLocation represents a user location in lat/long co-ordinates.
type UserLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// User represents a user of the application and holds data about them as well as their
// activity on the application.
type User struct {
	ID             uint `gorm:"primary_key"`
	Name           string
	Age            int
	Gender         string
	Location       UserLocation `gorm:"embedded"`
	Email          string
	Password       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	SwipesMade     []Swipe `gorm:"foreignKey:SwiperID"`
	SwipesReceived []Swipe `gorm:"foreignKey:SwipeeID"`
}
