package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewBollieClient(t *testing.T) {
	c := NewBollieClient(nil, "http://mockserver:3001")

	assert.NotNil(t, c.client)
	assert.Equal(t, "http://mockserver:3001", c.host.String())
}

func TestBollieClient_GetProduct(t *testing.T) {
	want, err := ioutil.ReadFile("mocks/product_response.json")

	if err != nil {
		t.Fatal(fmt.Sprintf("error while setting up test file: %q", err.Error()))
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(want)
	}))

	defer srv.Close()

	bollieClient := NewBollieClient(srv.Client(), srv.URL)

	got, err := bollieClient.GetProduct(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, got.ID)
	assert.Equal(t, "c0ffee", got.Name)
	assert.Equal(t, 13.37, got.Price)
	assert.Equal(t, "bollie", got.Source)
}

func TestBollieClient_GetProduct_emptyResponse(t *testing.T) {
	var want []byte

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(want)
	}))

	defer srv.Close()

	bollieClient := NewBollieClient(srv.Client(), srv.URL)

	_, err := bollieClient.GetProduct(1)

	assert.Error(t, err)
}

func TestBollieClient_GetProduct_badRequest(t *testing.T) {
	want, err := ioutil.ReadFile("mocks/bad_request_response.json")

	if err != nil {
		t.Fatal(fmt.Sprintf("error while setting up test file: %q", err.Error()))
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(want)
	}))

	defer srv.Close()

	bollieClient := NewBollieClient(srv.Client(), srv.URL)

	_, err = bollieClient.GetProduct(1)

	assert.Error(t, err)
	assert.EqualError(t, err, "status 400: \"incorrect product id provided\"")
}

func TestBollieClient_GetProduct_badRequest_emptyBody(t *testing.T) {
	var want []byte

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(want)
	}))

	defer srv.Close()

	bollieClient := NewBollieClient(srv.Client(), srv.URL)

	_, err := bollieClient.GetProduct(1)

	assert.Error(t, err)
	assert.EqualError(t, err, "unexpected end of JSON input")
}

func TestBollieClient_GetProduct_incorrectHost(t *testing.T) {
	var want []byte

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(want)
	}))

	defer srv.Close()

	bollieClient := NewBollieClient(srv.Client(), "")

	_, err := bollieClient.GetProduct(1)

	assert.Error(t, err)
	assert.EqualError(t, err, "Get \"/cadeau/1\": unsupported protocol scheme \"\"")
}
