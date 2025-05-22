package models

// Order represents a domain entity to be reprocessed.
type Order struct {
	ID         string `json:"id"`
	Status     string `json:"status"`
	ProductKey string `json:"product_key"`
	// You can add more fields such as Timestamp, ErrorCode, etc.
}

// OrdersResponse represents the structure of the API response.
type OrdersResponse struct {
	Data       []Order `json:"data"`
	TotalPages int      `json:"total_pages"`
}