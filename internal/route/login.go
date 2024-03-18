package route

import (
	"net/http"
	"zumm/internal/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func LoginHandler(c echo.Context) error {
	var login model.LoginRequest
	err := c.Bind(&login)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	var user model.User
	result := model.DB.Take(&user, "email = ?", login.Email)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusUnauthorized)
		}
		return c.NoContent(http.StatusInternalServerError)
	}
	if user.Password != login.Password {
		return c.NoContent(http.StatusUnauthorized)
	}

	claims := model.UserClaims{jwt.RegisteredClaims{}, user}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString(JWTSecretKey)

	response := model.LoginResponse{Token: t}
	return c.JSON(http.StatusOK, response)
}
