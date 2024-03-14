package main

import (
	"net/http"
	"time"
	"zumm/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var jwtSecretKey = []byte("tottenhamhotspurfootballclub")

func setupRouter() *echo.Echo {
	e := echo.New()

	configureMiddleware(e)
	configureRoutes(e)
	return e
}

func configureRoutes(e *echo.Echo) {
	e.GET("/", helloWorldHandler)
	e.GET("/user/create", userCreateHandler)
	e.POST("/login", LoginHandler)
}

func configureMiddleware(e *echo.Echo) {
	// TODO you are here
	// add a /login route posting email and pass and returning nothing
	// then implement issueing jwt tokens
	// then implement middleware and secure hello world with jwt

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

func LoginHandler(c echo.Context) error {
	var login models.UserLogin
	err := c.Bind(&login)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	var user models.User
	result := models.DB.Take(&user, "email = ?", login.Email)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Email or password incorrect"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}
	if user.Password != login.Password {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Email or password incorrect"})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["age"] = user.Age
	claims["gender"] = user.Gender
	claims["exp"] = time.Now().AddDate(0, 3, 0).Unix() // 90 days
	t, _ := token.SignedString(jwtSecretKey)

	return c.JSON(http.StatusOK, map[string]string{"token": t})
}

func main() {
	e := setupRouter()
	e.Logger.Fatal(e.Start(":8080"))

}
