package bollie

type Response struct {
	Product responseProduct `json:"product"`
}

type responseProduct struct {
	ID    int     `json:"id"`
	Price float64 `json:"price"`
	Name  string  `json:"name"`
}

type ResponseError struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}

type ResponsePurchase struct {
	Status   string  `json:"status"`
	Quantity int     `json:"quantity"`
	Product  product `json:"product"`
}

type PurchaseBody struct {
	Quantity  int `json:"quantity"`
	ProductID int `json:"product_id"`
}
