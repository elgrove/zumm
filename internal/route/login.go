package route

import (
	"fmt"
	"net/http"
	"zumm/internal/middleware"
	"zumm/internal/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// LoginHandler provides a HTTP interface to authenticate a user, taking email and password
// and returning a JWT which can be used to authenticate on protected routes.
func LoginHandler(c echo.Context) error {
	var login model.LoginRequest
	err := c.Bind(&login)
	if err != nil {
		middleware.Logger.Error("/login payload does not conform to LoginRequest")
		return c.String(http.StatusBadRequest, "bad request")
	}
	var user model.User
	result := model.DB.Take(&user, "email = ?", login.Email)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			middleware.Logger.Error(fmt.Sprintf("/login email address %s not found", login.Email))
			return c.NoContent(http.StatusUnauthorized)
		}
		return c.NoContent(http.StatusInternalServerError)
	}
	if user.Password != login.Password {
		middleware.Logger.Error("/login password does not match")
		return c.NoContent(http.StatusUnauthorized)
	}

	claims := model.UserClaims{jwt.RegisteredClaims{}, user}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString(middleware.JWTokenSecretKey)

	response := model.LoginResponse{Token: t}
	middleware.Logger.Debug(fmt.Sprintf("/login successful for user %d", user.ID))
	return c.JSON(http.StatusOK, response)
}
