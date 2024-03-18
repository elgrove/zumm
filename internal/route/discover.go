package route

import (
	"net/http"
	"sort"
	"zumm/internal/model"

	"github.com/LucaTheHacker/go-haversine"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func DiscoverHandler(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	var requestData model.DiscoverRequest
	c.Bind(&requestData)
	claims := token.Claims.(*model.UserClaims)

	var callingUser model.User
	model.DB.Take(&callingUser, "email = ?", claims.User.Email)
	var possibleDiscoverUsers []model.User
	model.DB.
		Where("ID <> ?", callingUser.ID).
		Where("Gender = ?", requestData.DesiredGender).
		Where("Age BETWEEN ? AND ?", requestData.DesiredAgeMin, requestData.DesiredAgeMax).
		Find(&possibleDiscoverUsers)

	type UserDistance struct {
		User     model.User
		Distance float64
	}
	var usersWithDistance []UserDistance

	for _, user := range possibleDiscoverUsers {
		distance := haversine.Distance(
			haversine.NewCoordinates(callingUser.Location.Latitude, callingUser.Location.Longitude),
			haversine.NewCoordinates(user.Location.Latitude, user.Location.Longitude),
		).Kilometers()
		usersWithDistance = append(usersWithDistance, UserDistance{User: user, Distance: distance})
	}

	sort.Slice(usersWithDistance, func(i, j int) bool {
		return usersWithDistance[i].Distance < usersWithDistance[j].Distance
	})

	var discoverUserProfiles []model.DiscoverUserProfile
	for _, uwd := range usersWithDistance {
		discoverProfile := model.DiscoverUserProfile{
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
