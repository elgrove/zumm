package routes

import (
	"net/http"
	"zumm/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func loginHandler(c echo.Context) error {
	var login models.UserLogin
	err := c.Bind(&login)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	var user models.User
	result := models.DB.Take(&user, "email = ?", login.Email)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusUnauthorized)
		}
		return c.NoContent(http.StatusInternalServerError)
	}
	if user.Password != login.Password {
		return c.NoContent(http.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["age"] = user.Age
	claims["gender"] = user.Gender
	t, _ := token.SignedString(JWTSecretKey)

	return c.JSON(http.StatusOK, map[string]string{"token": t})
}
