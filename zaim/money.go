package zaim

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type Mode string
type Order string

const (
	Payment  Mode = "payment"
	Income   Mode = "income"
	Transfer Mode = "transfer"
)

const (
	Id   Order = "id"
	Date Order = "date"
)

type ListRecordsRequest struct {
	CategoryId int
	GenereId   int
	Mode       Mode
	StartDate  time.Time
	EndDate    time.Time
	Page       int
	Limit      int
	GroupBy    string
	Order      Order
}

type Money struct {
	ID         int    `json:"id"`
	Mode       string `json:"mode"`
	UserID     int    `json:"user_id"`
	Date       string `json:"date"`
	CategoryID int    `json:"category_id"`
	GenreID    int    `json:"genre_id"`
	Amount     int    `json:"amount"`
	Comment    string `json:"comment"`
	Place      string `json:"place"`
	Created    string `json:"created"`
	Modified   string `json:"modified"`
}

type ListRecordsResponse struct {
	Money []Money `json:"money"`
}

func (z *Zaim) ListRecords(l *ListRecordsRequest) (*ListRecordsResponse, error) {
	u, _ := url.Parse("https://api.zaim.net/v2/home/money")
	params := u.Query()

	if l.CategoryId != 0 {
		params.Set("category_id", strconv.Itoa(l.CategoryId))
	}
	if l.GenereId != 0 {
		params.Set("genre_id", strconv.Itoa(l.GenereId))
	}
	if l.Mode != "" {
		params.Set("mode", string(l.Mode))
	}
	if !l.StartDate.IsZero() {
		params.Set("start_date", l.StartDate.Format("2006-01-02"))
	}
	if !l.EndDate.IsZero() {
		params.Set("end_date", l.EndDate.Format("2006-01-02"))
	}
	if l.Page != 0 {
		params.Set("page", strconv.Itoa(l.Page))
	}
	if l.Limit != 0 {
		params.Set("limit", strconv.Itoa(l.Limit))
	}
	if l.GroupBy != "" {
		params.Set("group_by", l.GroupBy)
	}
	if l.Order != "" {
		params.Set("order", string(l.Order))
	}

	u.RawQuery = params.Encode()

	resp, err := z.client.Get(u.String())
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
