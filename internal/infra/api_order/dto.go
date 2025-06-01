package api_order

// OrderDTO representa o Data Transfer Object para Order (usado para transporte/serialização).
type OrderDTO struct {
	ID         string `json:"id"`
	Status     string `json:"status"`
	ProductKey string `json:"product_key"`
}

type OrdersResponse struct {
	Data       []OrderDTO `json:"data"`
	TotalPages int        `json:"total_pages"`
}
