package routes

import (
	"fmt"
	"net/http"
	"zumm/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func discoverHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.UserClaims)
	fmt.Println(claims)
	// get user location... from where?
	// get all users from db
	// remove callinguser from the resultset
	// sort by distance from callinguser
	// put some sort of limit on the distance
	// add distance from callinguser to resultset
	// return {results : [users] }
	//
	// shouldn't I be using middleware to authenticate the token?
	//

	return c.NoContent(http.StatusOK)

}
