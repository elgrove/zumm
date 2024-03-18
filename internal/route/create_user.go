package route

import (
	"net/http"
	"zumm/internal/model"

	"github.com/labstack/echo/v4"
)

func UserCreateHandler(c echo.Context) error {
	user := model.CreateRandomUser()
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
