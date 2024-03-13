package main

import (
	"net/http"
	"zumm/models"

	"github.com/labstack/echo/v4"
)

func setupRouter() *echo.Echo {
	e := echo.New()

	configureMiddleware(e)
	configureRoutes(e)
	return e
}

func configureRoutes(e *echo.Echo) {
	e.GET("/", helloWorldHandler)
	e.GET("/user/create", userCreateHandler)
}

func configureMiddleware(e *echo.Echo) {
}

func helloWorldHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"hello": "world"})
}

func userCreateHandler(c echo.Context) error {
	user := models.CreateRandomUser()
	models.DB.Create(&user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":       user.ID,
		"email":    user.Email,
		"password": user.Password,
		"name":     user.Name,
		"gender":   user.Gender,
		"age":      user.Age,
	})
}

func main() {
	e := setupRouter()
	e.Logger.Fatal(e.Start(":8080"))

}
