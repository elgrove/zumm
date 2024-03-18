package route

import (
	"net/http"
	"zumm/internal/model"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/labstack/echo/v4"
)

// CreateRandomUser returns a User model filled with random data provided by the gofakeit
// package. The location is roughly within the M25/Greater London area.
func CreateRandomUser() model.User {
	lat, _ := gofakeit.LatitudeInRange(51.245, 51.759)
	long, _ := gofakeit.LongitudeInRange(-0.302, 0.285)
	location := model.UserLocation{Latitude: lat, Longitude: long}
	user := model.User{
		Name:     gofakeit.Name(),
		Age:      gofakeit.Number(18, 100),
		Gender:   gofakeit.Gender(),
		Location: location,
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, false, false, false, false, 12),
	}
	return user
}

// UserCreateHandler provides a HTTP interface to create a random user using the
// CreateRandomUser function and return it in JSON format.
func UserCreateHandler(c echo.Context) error {
	user := CreateRandomUser()
	model.DB.Create(&user)
	// TODO return {result : {user}}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":       user.ID,
		"email":    user.Email,
		"password": user.Password,
		"name":     user.Name,
		"gender":   user.Gender,
		"age":      user.Age,
	})
}
