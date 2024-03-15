package routes

import (
	"net/http"
	"sort"
	"zumm/models"

	"github.com/LucaTheHacker/go-haversine"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func discoverHandler(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	var requestData models.DiscoverRequest
	c.Bind(&requestData)
	claims := token.Claims.(*models.UserClaims)
	var callingUser models.User
	models.DB.Take(&callingUser, "email = ?", claims.User.Email)
	var allOtherUsers []models.User
	models.DB.Where("ID <> ?", callingUser.ID).Find(&allOtherUsers)

	type UserDistance struct {
		User     models.User
		Distance float64
	}

	var usersWithDistance []UserDistance

	for _, user := range allOtherUsers {
		distance := haversine.Distance(
			haversine.NewCoordinates(callingUser.Location.Latitude, callingUser.Location.Longitude),
			haversine.NewCoordinates(user.Location.Latitude, user.Location.Longitude),
		).Kilometers()
		usersWithDistance = append(usersWithDistance, UserDistance{User: user, Distance: distance})
	}

	sort.Slice(usersWithDistance, func(i, j int) bool {
		return usersWithDistance[i].Distance < usersWithDistance[j].Distance
	})

	var discoverUserProfiles []models.DiscoverUserProfile
	for _, uwd := range usersWithDistance {
		discoverProfile := models.DiscoverUserProfile{
			ID:             uwd.User.ID,
			Name:           uwd.User.Name,
			Gender:         uwd.User.Gender,
			Age:            uwd.User.Age,
			DistanceFromMe: uwd.Distance,
		}
		discoverUserProfiles = append(discoverUserProfiles, discoverProfile)
	}

	return c.JSON(http.StatusOK, echo.Map{"results": discoverUserProfiles})

}
