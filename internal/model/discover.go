package model

// DiscoverUserProfile represents another user of the app that the current user could match with.
type DiscoverUserProfile struct {
	ID             uint
	Name           string
	Age            int
	Gender         string
	DistanceFromMe float64
}

// DiscoverResponse holds a slice of DiscoverUserProfiles.
type DiscoverResponse struct {
	Results []DiscoverUserProfile
}

// DiscoverRequest represents the request to Discover other users and includes filter points
// such as desired age and gender.
type DiscoverRequest struct {
	Location      UserLocation
	DesiredGender string
	DesiredAgeMin int
	DesiredAgeMax int
}
