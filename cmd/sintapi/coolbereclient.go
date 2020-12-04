package main

import (
	"encoding/json"
	"fmt"
	"github.com/TestCoders/Sinterklaas2020/pkg/coolbere"
	"net/http"
	"net/url"
	"time"
)

type CoolbereClient struct {
	client HTTPClient
	host   *url.URL
}

func (c *CoolbereClient) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

func (c *CoolbereClient) GetHost() *url.URL {
	return c.host
}

func (c *CoolbereClient) GetProduct(id int) (Product, error) {
	data, statusCode, err := getProduct(c, id)

	if err != nil {
		return Product{}, err
	}

	if statusCode != http.StatusOK {
		var result coolbere.ResponseError

		if err = json.Unmarshal(data, &result); err != nil {
			return Product{}, err
		}

		return Product{}, fmt.Errorf("status %v: %q", result.Error, result.Description)
	}

	var result coolbere.Response

	if err = json.Unmarshal(data, &result); err != nil {
		return Product{}, err
	}

	return Product{
		ID:     result.Product.ID,
		Name:   result.Product.Name,
		Price:  result.Product.Price,
		Source: "coolbere",
	}, nil
}

func (c *CoolbereClient) PurchaseProduct(id, quantity int) (*PurchaseResponse, *ErrorResponse) {
	b := coolbere.PurchaseBody{
		Quantity:  quantity,
		ProductID: id,
	}

	body, err := json.Marshal(b)

	if err != nil {
		return nil, &ErrorResponse{
			Status:      http.StatusInternalServerError,
			Description: "Internal server error",
		}
	}

	data, status, err := purchaseProduct(c, id, body)

	if status != http.StatusOK {
		var resp coolbere.ResponseError

		if err = json.Unmarshal(data, &resp); err != nil {
			return nil, &ErrorResponse{
				Status:      status,
				Description: err.Error(),
			}
		}
	}

	var resp coolbere.ResponsePurchase

	if err = json.Unmarshal(data, &resp); err != nil {
		return nil, &ErrorResponse{
			Status:      500,
			Description: "internal server error",
		}
	}

	return &PurchaseResponse{
		Quantity: resp.Quantity,
		ID:       resp.Product.ID,
		Name:     resp.Product.Name,
		Price:    resp.Product.Price,
		Supplier: "coolbere",
	}, nil
}

func NewCoolbereClient(client HTTPClient, host string) *CoolbereClient {
	if client == nil {
		client = &http.Client{
			Timeout: time.Second * 10,
		}
	}

	h, _ := url.Parse(host)

	return &CoolbereClient{
		client: client,
		host:   h,
	}
}
