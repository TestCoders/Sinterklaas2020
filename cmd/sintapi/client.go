package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// HTTPClient allows for mocking api responses
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type SintClient interface {
	Do(req *http.Request) (*http.Response, error)
	GetHost() *url.URL
	GetProduct(id int) (Product, error)
	PurchaseProduct(id, quantity int) (*PurchaseResponse, *ErrorResponse)
}

func getProduct(c SintClient, id int) ([]byte, int, error) {
	ref, err := url.Parse(fmt.Sprintf("/cadeau/%v", id))

	if err != nil {
		return nil, 0, err
	}

	u := c.GetHost().ResolveReference(ref)

	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return nil, 0, err
	}

	resp, err := c.Do(req)

	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, resp.StatusCode, err
	}

	return body, resp.StatusCode, nil
}

func purchaseProduct(c SintClient, id int, body []byte) ([]byte, int, error) {
	ref, err := url.Parse(fmt.Sprintf("/cadeau/%v", id))

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	u := c.GetHost().ResolveReference(ref)

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))

	resp, err := c.Do(req)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return respBody, http.StatusOK, nil
}
