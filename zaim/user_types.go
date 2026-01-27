package zaim

// VerifyUserAuthenticationResponse contains user data from authentication verification.
type VerifyUserAuthenticationResponse struct {
	ID              int    `json:"id,omitempty"`
	Login           string `json:"login,omitempty"`
	Name            string `json:"name,omitempty"`
	InputCount      int    `json:"input_count,omitempty"`
	DayCount        int    `json:"day_count,omitempty"`
	RepeatCount     int    `json:"repeat_count,omitempty"`
	Day             int    `json:"day,omitempty"`
	Week            int    `json:"week,omitempty"`
	Month           int    `json:"month,omitempty"`
	CurrencyCode    string `json:"currency_code,omitempty"`
	ProfileImageUrl string `json:"profile_image_url,omitempty"`
	CoverImageUrl   string `json:"cover_image_url,omitempty"`
	ProfileModified string `json:"profile_modified,omitempty"`
}
