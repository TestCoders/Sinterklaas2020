package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/TestCoders/Sinterklaas2020/pkg/aliblabla"
	"net/http"
	"net/url"
	"time"
)

type AliblablaClient struct {
	client HTTPClient
	host   *url.URL
}

func (c *AliblablaClient) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

func (c *AliblablaClient) GetHost() *url.URL {
	return c.host
}

func (c *AliblablaClient) GetProduct(id int) (Product, error) {
	data, statusCode, err := getProduct(c, id)

	if err != nil {
		return Product{}, err
	}

	if statusCode != http.StatusOK {
		var result aliblabla.ResponseError

		if err = xml.Unmarshal(data, &result); err != nil {
			return Product{}, err
		}

		return Product{}, fmt.Errorf("status %v: %q", result.Error, result.Description)
	}

	var result aliblabla.Response

	if err = xml.Unmarshal(data, &result); err != nil {
		return Product{}, err
	}

	return Product{
		ID:     result.Product.ID,
		Name:   result.Product.Name,
		Price:  result.Product.Price,
		Source: "aliblabla",
	}, nil
}

func (c *AliblablaClient) PurchaseProduct(id, quantity int) (*PurchaseResponse, *ErrorResponse) {
	b := aliblabla.PurchaseBody{
		Quantity:  quantity,
		ProductID: id,
	}

	body, err := xml.Marshal(b)

	if err != nil {
		return nil, &ErrorResponse{
			Status:      http.StatusInternalServerError,
			Description: "Internal server error",
		}
	}

	data, status, err := purchaseProduct(c, id, body)

	if status != http.StatusOK {
		var resp aliblabla.ResponseError

		if err = json.Unmarshal(data, &resp); err != nil {
			return nil, &ErrorResponse{
				Status:      status,
				Description: err.Error(),
			}
		}
	}

	var resp aliblabla.ResponsePurchase

	if err = xml.Unmarshal(data, &resp); err != nil {
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
		Supplier: "aliblabla",
	}, nil
}

func NewAliblablaClient(client HTTPClient, host string) *AliblablaClient {
	if client == nil {
		client = &http.Client{
			Timeout: time.Second * 10,
		}
	}

	h, _ := url.Parse(host)

	return &AliblablaClient{
		client: client,
		host:   h,
	}
}
