package route

import (
	"errors"
	"net/http"
	"zumm/internal/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SwipeHandler(c echo.Context) error {
	var swipe model.Swipe
	c.Bind(&swipe)
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.UserClaims)
	var callingUser model.User
	model.DB.Take(&callingUser, "id = ?", claims.User.ID)
	if swipe.SwiperID != callingUser.ID {
		return c.NoContent(http.StatusBadRequest)
	}
	model.DB.Create(&swipe)
	var matchedSwipe model.Swipe
	result := model.DB.
		Where("swiper_id = ?", swipe.SwiperID).
		Where("swipee_id = ?", swipe.SwipeeID).
		Where("interested = ?", true).
		First(&matchedSwipe)
	var swipeResult model.SwipeResult
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		swipeResult = model.SwipeResult{Matched: false}

	} else {
		swipeResult = model.SwipeResult{Matched: true, MatchID: &matchedSwipe.SwipeeID}
	}

	swipeResponse := model.SwipeResponse{Results: swipeResult}
	return c.JSON(http.StatusOK, swipeResponse)
}
