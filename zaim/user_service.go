package zaim

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// UserService handles user-related API operations.
type UserService struct {
	client *http.Client
}

// Verify verifies the user's authentication and returns user information.
func (u *UserService) Verify() (bool, *VerifyUserAuthenticationResponse) {
	resp, err := u.client.Get("https://api.zaim.net/v2/home/user/verify")
	if err != nil {
		fmt.Printf("there was an error when making verify authentication request: %s\n", err.Error())
		return false, nil
	}
	defer resp.Body.Close()

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

	return true, r
}
