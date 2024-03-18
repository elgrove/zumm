package model

// DiscoverUserProfile represents another user of the app that the current user could match with.
type DiscoverUserProfile struct {
	ID             uint    `json:"id"`
	Name           string  `json:"name"`
	Age            int     `json:"age"`
	Gender         string  `json:"gender"`
	DistanceFromMe float64 `json:"distance_from_me"`
}

// DiscoverResponse holds a slice of DiscoverUserProfiles.
type DiscoverResponse struct {
	Results []DiscoverUserProfile `json:"results"`
}

// DiscoverRequest represents the request to Discover other users and includes filter points
// such as desired age and gender.
type DiscoverRequest struct {
	Location      UserLocation `json:"location"`
	DesiredGender string       `json:"desired_gender"`
	DesiredAgeMin int          `json:"desired_age_min"`
	DesiredAgeMax int          `json:"desired_age_max"`
}
