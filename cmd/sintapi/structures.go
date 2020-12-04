package main

type ErrorResponse struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
}

type PurchaseResponse struct {
	Quantity int     `json:"quantity"`
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Supplier string  `json:"supplier"`
}

type PurchaseRequest struct {
	Quantity int `json:"quantity"`
}
