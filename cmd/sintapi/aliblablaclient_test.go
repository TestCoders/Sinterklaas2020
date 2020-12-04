package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewAliblablaClient(t *testing.T) {
	c := NewAliblablaClient(nil, "http://mockserver:3001")

	assert.NotNil(t, c.client)
	assert.Equal(t, "http://mockserver:3001", c.host.String())
}

func TestAliblablaClient_GetProduct(t *testing.T) {
	want, err := ioutil.ReadFile("mocks/product_response.xml")

	if err != nil {
		t.Fatal(fmt.Sprintf("error while setting up test file: %q", err.Error()))
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(want)
	}))

	defer srv.Close()

	aliblablaClient := NewAliblablaClient(srv.Client(), srv.URL)

	got, err := aliblablaClient.GetProduct(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, got.ID)
	assert.Equal(t, "c0ffee", got.Name)
	assert.Equal(t, 13.37, got.Price)
	assert.Equal(t, "aliblabla", got.Source)
}

func TestAliblablaClient_GetProduct_emptyResponse(t *testing.T) {
	var want []byte

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(want)
	}))

	defer srv.Close()

	aliblablaClient := NewAliblablaClient(srv.Client(), srv.URL)

	_, err := aliblablaClient.GetProduct(1)

	assert.Error(t, err)
}

func TestAliblablaClient_GetProduct_badRequest(t *testing.T) {
	want, err := ioutil.ReadFile("mocks/bad_request_response.xml")

	if err != nil {
		t.Fatal(fmt.Sprintf("error while setting up test file: %q", err.Error()))
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(want)
	}))

	defer srv.Close()

	aliblablaClient := NewAliblablaClient(srv.Client(), srv.URL)

	_, err = aliblablaClient.GetProduct(1)

	assert.Error(t, err)
	assert.EqualError(t, err, "status 400: \"incorrect product id provided\"")
}

func TestAliblablaClient_GetProduct_badRequest_emptyBody(t *testing.T) {
	var want []byte

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(want)
	}))

	defer srv.Close()

	aliblablaClient := NewAliblablaClient(srv.Client(), srv.URL)

	_, err := aliblablaClient.GetProduct(1)

	assert.Error(t, err)
	assert.EqualError(t, err, "EOF")
}

func TestAliblablaClient_GetProduct_incorrectHost(t *testing.T) {
	var want []byte

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(want)
	}))

	defer srv.Close()

	aliblablaClient := NewAliblablaClient(srv.Client(), "")

	_, err := aliblablaClient.GetProduct(1)

	assert.Error(t, err)
	assert.EqualError(t, err, "Get \"/cadeau/1\": unsupported protocol scheme \"\"")
}
