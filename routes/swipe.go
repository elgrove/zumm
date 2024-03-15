package routes

import (
	"errors"
	"net/http"
	"zumm/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func swipeHandler(c echo.Context) error {
	var swipe models.Swipe
	c.Bind(&swipe)
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*models.UserClaims)
	var callingUser models.User
	models.DB.Take(&callingUser, "id = ?", claims.User.ID)
	if swipe.SwiperID != callingUser.ID {
		return c.NoContent(http.StatusUnauthorized)
	}
	models.DB.Create(&swipe)
	var matchedSwipe models.Swipe
	result := models.DB.
		Where("swiper_id = ?", swipe.SwiperID).
		Where("swipee_id = ?", swipe.SwipeeID).
		Where("interested = ?", true).
		First(&matchedSwipe)
	var swipeResult models.SwipeResult
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		swipeResult = models.SwipeResult{Matched: false}

	} else {
		swipeResult = models.SwipeResult{Matched: true, MatchID: &matchedSwipe.SwipeeID}
	}

	swipeResponse := models.SwipeResponse{Results: swipeResult}
	return c.JSON(http.StatusOK, swipeResponse)
}
