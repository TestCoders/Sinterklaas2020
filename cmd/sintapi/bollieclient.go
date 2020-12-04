package main

import (
	"encoding/json"
	"fmt"
	"github.com/TestCoders/Sinterklaas2020/pkg/bollie"
	"net/http"
	"net/url"
	"time"
)

type BollieClient struct {
	client HTTPClient
	host   *url.URL
}

func (c *BollieClient) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

func (c *BollieClient) GetHost() *url.URL {
	return c.host
}

func (c *BollieClient) GetProduct(id int) (Product, error) {
	data, statusCode, err := getProduct(c, id)

	if err != nil {
		return Product{}, err
	}

	if statusCode != http.StatusOK {
		var result bollie.ResponseError

		if err = json.Unmarshal(data, &result); err != nil {
			return Product{}, err
		}

		return Product{}, fmt.Errorf("status %v: %q", result.Error, result.Description)
	}

	var result bollie.Response

	if err = json.Unmarshal(data, &result); err != nil {
		return Product{}, err
	}

	return Product{
		ID:     result.Product.ID,
		Name:   result.Product.Name,
		Price:  result.Product.Price,
		Source: "bollie",
	}, nil
}

func (c *BollieClient) PurchaseProduct(id, quantity int) (*PurchaseResponse, *ErrorResponse) {
	b := bollie.PurchaseBody{
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
		var resp bollie.ResponseError

		if err = json.Unmarshal(data, &resp); err != nil {
			return nil, &ErrorResponse{
				Status:      status,
				Description: err.Error(),
			}
		}
	}

	var resp bollie.ResponsePurchase

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
		Supplier: "bollie",
	}, nil
}

func NewBollieClient(client HTTPClient, host string) *BollieClient {
	if client == nil {
		client = &http.Client{
			Timeout: time.Second * 10,
		}
	}

	h, _ := url.Parse(host)

	return &BollieClient{
		client: client,
		host:   h,
	}
}
