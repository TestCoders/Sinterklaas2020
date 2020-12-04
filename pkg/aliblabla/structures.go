package aliblabla

type Response struct {
	Product responseProduct `xml:"product"`
}

type responseProduct struct {
	ID    int     `xml:"id,attr"`
	Price float64 `xml:"price"`
	Name  string  `xml:"name"`
}

type ResponseError struct {
	Error       string `xml:"error,attr"`
	Description string `xml:"description"`
}

type ResponsePurchase struct {
	Status   string  `xml:"status"`
	Quantity int     `xml:"quantity"`
	Product  product `xml:"product"`
}

type PurchaseBody struct {
	Quantity  int `xml:"quantity"`
	ProductID int `xml:"product_id"`
}
