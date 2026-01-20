package zaim

import (
	"encoding/json"
	"fmt"
)

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

func (z *Zaim) VerifyUserAuthentication() (bool, *VerifyUserAuthenticationResponse) {
	resp, err := z.client.Get("https://api.zaim.net/v2/home/user/verify")
	if err != nil {
		fmt.Printf("there was an error when making verify authentication request: %s\n", err.Error())
		return false, nil
	}

	var v verifyUserAuthenticationResponseDto
	err = json.NewDecoder(resp.Body).Decode(&v)

	if err != nil {
		fmt.Printf("there was an error decoding %s\n", err.Error())

		return false, nil
	}

	r, err := v.ToResponse()

	if err != nil {
		fmt.Printf("there was an error parsing the DTO: %s\n", err.Error())
		return false, nil
	}

	defer resp.Body.Close()

	return true, r
}
