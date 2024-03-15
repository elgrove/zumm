package routes

import (
	"net/http"
	"zumm/models"

	"github.com/labstack/echo/v4"
)

func userCreateHandler(c echo.Context) error {
	user := models.CreateRandomUser()
	models.DB.Create(&user)
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
