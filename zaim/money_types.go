package zaim

import "time"

// Mode represents the type of money record
type Mode string

// Order represents the sort order for listing records
type Order string

const (
	Payment  Mode = "payment"
	Income   Mode = "income"
	Transfer Mode = "transfer"
)

const (
	OrderById   Order = "id"
	OrderByDate Order = "date"
)

// ListRecordsRequest contains parameters for listing money records.
type ListRecordsRequest struct {
	CategoryId int       // Optional: filter by category
	GenreId    int       // Optional: filter by genre
	Mode       Mode      // Optional: filter by mode (payment/income/transfer)
	StartDate  time.Time // Optional: filter records from this date
	EndDate    time.Time // Optional: filter records until this date
	Page       int       // Optional: page number for pagination
	Limit      int       // Optional: max records to return
	GroupBy    string    // Optional: group results
	Order      Order     // Optional: sort order (id/date)
}

// ListRecordsResponse is the response from listing money records.
type ListRecordsResponse struct {
	Money []Money `json:"money"`
}

// NewPaymentRequest creates a new PaymentRequest with required fields.
func NewPaymentRequest(categoryId, genreId, amount int, date time.Time) *PaymentRequest {
	return &PaymentRequest{
		CategoryId: categoryId,
		GenreId:    genreId,
		Amount:     amount,
		Date:       date,
	}
}

// PaymentRequest contains parameters for posting a payment.
type PaymentRequest struct {
	CategoryId    int       // Required
	GenreId       int       // Required
	Amount        int       // Required
	Date          time.Time // Required
	FromAccountId int       // Optional: account ID to debit
	Comment       string    // Optional: memo/note
	Name          string    // Optional: item name
	Place         string    // Optional: location
}

// PaymentResponse is the response from posting a payment.
type PaymentResponse struct {
	Stamps       string       `json:"stamps"`
	Banners      []string     `json:"banners"`
	PaymentMoney PaymentMoney `json:"money"`
	Place        Place        `json:"place,omitempty"`
	User         User         `json:"user"`
	Requested    int          `json:"requested"`
}

// User contains user metadata in responses.
type User struct {
	InputCount   int    `json:"input_count"`
	RepeatCount  int    `json:"repeat_count"`
	DayCount     int    `json:"day_count"`
	DateModified string `json:"date_modified"`
}

// Place represents a location/merchant.
type Place struct {
	ID                int    `json:"id"`
	UserID            int    `json:"user_id"`
	GenreID           int    `json:"genre_id"`
	CategoryID        int    `json:"category_id"`
	AccountID         int    `json:"account_id"`
	TransferAccountID int    `json:"transfer_account_id"`
	Mode              string `json:"mode"`
	PlaceUID          string `json:"place_uid"`
	Service           string `json:"service"`
	Name              string `json:"name"`
	OriginalName      string `json:"original_name"`
	Tel               string `json:"tel"`
	Count             int    `json:"count"`
	PlacePatternID    int    `json:"place_pattern_id"`
	CalcFlag          int    `json:"calc_flag"`
	EditFlag          int    `json:"edit_flag"`
	Active            int    `json:"active"`
	Modified          string `json:"modified"`
	Created           string `json:"created"`
}

// Money represents a money record.
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

// PaymentMoney contains payment details in responses.
type PaymentMoney struct {
	ID       int    `json:"id"`
	PlaceUid string `json:"place_uid,omitempty"`
	Modified string `json:"modified"`
}
