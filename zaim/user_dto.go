package zaim

type verifyUserAuthenticationResponseDto struct {
	Me        me    `json:"me"`
	Requested int64 `json:"requested"`
}

type me struct {
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

func (v *verifyUserAuthenticationResponseDto) ToResponse() (*VerifyUserAuthenticationResponse, error) {
	return &VerifyUserAuthenticationResponse{
		ID:              v.Me.ID,
		Login:           v.Me.Login,
		Name:            v.Me.Name,
		InputCount:      v.Me.InputCount,
		DayCount:        v.Me.DayCount,
		RepeatCount:     v.Me.RepeatCount,
		Day:             v.Me.Day,
		Week:            v.Me.Week,
		Month:           v.Me.Month,
		CurrencyCode:    v.Me.CurrencyCode,
		ProfileImageUrl: v.Me.ProfileImageUrl,
		CoverImageUrl:   v.Me.CoverImageUrl,
		ProfileModified: v.Me.ProfileModified,
	}, nil
}
