package main

import (
	"net/http"
	"net/url"
)

type MockSintClient struct {
	Product
	Host          *url.URL
	Response      *http.Response
	GetError      error
	PurchaseError error
	ResponseError error
}

func (m *MockSintClient) Do(_ *http.Request) (*http.Response, error) {
	return m.Response, m.ResponseError
}

func (m *MockSintClient) GetHost() *url.URL {
	return m.Host
}

func (m *MockSintClient) GetProduct(id int) (Product, error) {
	return m.Product, m.GetError
}

func (m *MockSintClient) PurchaseProduct(id int) error {
	return m.PurchaseError
}
