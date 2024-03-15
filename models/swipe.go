package models

type Swipe struct {
	ID         uint `gorm:"primary_key"`
	SwiperID   uint
	SwipeeID   uint
	Swiper     User `gorm:"foreignKey:SwiperID"`
	Swipee     User `gorm:"foreignKey:SwipeeID"`
	Interested bool
}
type SwipeResult struct {
	Matched bool
	MatchID *uint
}

type SwipeResponse struct {
	Results SwipeResult
}
