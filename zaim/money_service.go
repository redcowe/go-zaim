package zaim

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// MoneyService handles money-related API operations.
type MoneyService struct {
	client *http.Client
}

// List retrieves money records based on the provided request parameters.
func (m *MoneyService) List(req *ListRecordsRequest) (*ListRecordsResponse, error) {
	u, _ := url.Parse("https://api.zaim.net/v2/home/money")
	params := u.Query()

	if req.CategoryId != 0 {
		params.Set("category_id", strconv.Itoa(req.CategoryId))
	}
	if req.GenreId != 0 {
		params.Set("genre_id", strconv.Itoa(req.GenreId))
	}
	if req.Mode != "" {
		params.Set("mode", string(req.Mode))
	}
	if !req.StartDate.IsZero() {
		params.Set("start_date", req.StartDate.Format("2006-01-02"))
	}
	if !req.EndDate.IsZero() {
		params.Set("end_date", req.EndDate.Format("2006-01-02"))
	}
	if req.Page != 0 {
		params.Set("page", strconv.Itoa(req.Page))
	}
	if req.Limit != 0 {
		params.Set("limit", strconv.Itoa(req.Limit))
	}
	if req.GroupBy != "" {
		params.Set("group_by", req.GroupBy)
	}
	if req.Order != "" {
		params.Set("order", string(req.Order))
	}

	u.RawQuery = params.Encode()

	resp, err := m.client.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("error making list records request: %w", err)
	}
	defer resp.Body.Close()

	var result ListRecordsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding list records response: %w", err)
	}

	return &result, nil
}

// Post creates a new payment record.
func (m *MoneyService) CreatePayment(req *PaymentRequest) (*PaymentResponse, error) {
	u, _ := url.Parse("https://api.zaim.net/v2/home/money/payment")
	values := url.Values{}
	values.Set("mapping", "1")
	values.Set("category_id", strconv.Itoa(req.CategoryId))
	values.Set("genre_id", strconv.Itoa(req.GenreId))
	values.Set("amount", strconv.Itoa(req.Amount))
	values.Set("date", req.Date.Format("2006-01-02"))

	if req.FromAccountId != 0 {
		values.Set("from_account_id", strconv.Itoa(req.FromAccountId))
	}
	if req.Comment != "" {
		values.Set("comment", req.Comment)
	}
	if req.Name != "" {
		values.Set("name", req.Name)
	}
	if req.Place != "" {
		values.Set("place", req.Place)
	}

	resp, err := m.client.PostForm(u.String(), values)
	if err != nil {
		return nil, fmt.Errorf("error making post payment request: %w", err)
	}
	defer resp.Body.Close()

	var result PaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding post payment response: %w", err)
	}

	return &result, nil
}
