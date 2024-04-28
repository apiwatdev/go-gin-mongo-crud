package dto

type OrderCreateRequestDTO struct {
	Customer string                      `json:"customer"`
	Total    float64                     `json:"total"`
	Items    []OrderItemCreateRequestDTO `json:"items"`
}

type OrderItemCreateRequestDTO struct {
	OrderId  string  `json:"orderId"`
	Product  string  `json:"product"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
