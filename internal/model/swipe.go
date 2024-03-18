package model

// Swipe represents a request to the swipe endpoint and holds data about the user's swipe
// including who did the swipe, who they swiped on and their verdict on the user.
type Swipe struct {
	ID         uint `gorm:"primary_key" json:"id"`
	SwiperID   uint `json:"swiper_id"`
	SwipeeID   uint `json:"swipee_id"`
	Swiper     User `gorm:"foreignKey:SwiperID" json:"-"`
	Swipee     User `gorm:"foreignKey:SwipeeID" json:"-"`
	Interested bool `json:"interested"`
}

// SwipeResult holds data about if the swipe was reciprocated and, if so, the ID of
// the user they are matched with.
type SwipeResult struct {
	Matched bool  `json:"matched"`
	MatchID *uint `json:"match_id"`
}

// SwipeResponse represents a response from the swipe endpoint, holding a SwipeResult.
type SwipeResponse struct {
	Results SwipeResult `json:"results"`
}
