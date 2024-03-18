package model

// Swipe represents a request to the swipe endpoint and holds data about the user's swipe
// including who did the swipe, who they swiped on and their verdict on the user.
type Swipe struct {
	ID         uint `gorm:"primary_key"`
	SwiperID   uint
	SwipeeID   uint
	Swiper     User `gorm:"foreignKey:SwiperID"`
	Swipee     User `gorm:"foreignKey:SwipeeID"`
	Interested bool
}

// SwipeResult holds data about if the swipe was reciprocated and, if so, the ID of
// the user they are matched with.
// TODO I think the matchID should actually come from a seperate table of matches
type SwipeResult struct {
	Matched bool
	MatchID *uint
}

// SwipeResponse represents a response from the swipe endpoint, holding a SwipeResult.
type SwipeResponse struct {
	Results SwipeResult
}
