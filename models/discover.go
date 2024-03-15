package models

type DiscoverUserProfile struct {
	ID             uint
	Name           string
	Age            int
	Gender         string
	DistanceFromMe float64
}

type DiscoverResponse struct {
	Results []DiscoverUserProfile
}

type DiscoverRequest struct {
	Location      UserLocation
	DesiredGender string
	DesiredAgeMin int
	DesiredAgeMax int
}
