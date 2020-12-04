package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func readFile(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

func TestApplication_purchaseProductHandler(t *testing.T) {
	want, err := readFile("mocks/confirmation_response.json")

	if err != nil {
		t.Fatal(fmt.Sprintf("error while setting up test file: %q", err.Error()))
	}

	wantProductResponse, err := ioutil.ReadFile("mocks/product_response.json")

	if err != nil {
		t.Fatal(fmt.Sprintf("error while setting up test file: %q", err.Error()))
	}

	app := &Application{}
	logger := &bytes.Buffer{}
	app.errorLog = log.New(logger, "ERROR:\t", 0)
	app.infoLog = log.New(logger, "INFO:\t", 0)
	app.sources = map[string]SintClient{}

	router := mux.NewRouter()
	router.HandleFunc("/purchase/{id}", app.purchaseProductHandler).Methods("POST")
	router.HandleFunc("/cadeau/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(wantProductResponse)
	}).Methods("GET")

	srv := httptest.NewServer(router)
	defer srv.Close()

	app.sources["bollie"] = NewBollieClient(srv.Client(), srv.URL)

	u, err := url.Parse(srv.URL + "/purchase/1")

	if err != nil {
		t.Fatal(err)
	}

	purchaseRequest := &PurchaseRequest{
		Quantity: 1,
	}

	data, err := json.Marshal(purchaseRequest)

	if err != nil {
		t.Fatal(err.Error())
	}

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(data))

	assert.NoError(t, err)
	got, err := srv.Client().Do(req)
	assert.NoError(t, err)
	body, err := ioutil.ReadAll(got.Body)
	assert.NoError(t, err)
	assert.Equal(t, string(want), string(body))
}

func TestApplication_purchaseProductHandler_purchaseError(t *testing.T) {
	wantProductResponse, err := ioutil.ReadFile("mocks/product_response.json")

	if err != nil {
		t.Fatal(fmt.Sprintf("error while setting up test file: %q", err.Error()))
	}

	app := &Application{}
	logger := &bytes.Buffer{}
	app.errorLog = log.New(logger, "ERROR:\t", 0)
	app.infoLog = log.New(logger, "INFO:\t", 0)
	app.sources = map[string]SintClient{}

	router := mux.NewRouter()
	router.HandleFunc("/purchase/{id}", app.purchaseProductHandler)
	router.HandleFunc("/cadeau/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(wantProductResponse)
	}).Methods("GET")
	router.HandleFunc("/cadeau/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
	}).Methods("POST")

	srv := httptest.NewServer(router)
	defer srv.Close()

	app.sources["bollie"] = NewBollieClient(srv.Client(), srv.URL)

	u, err := url.Parse(srv.URL + "/purchase/1")

	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", u.String(), nil)

	assert.NoError(t, err)
	got, err := srv.Client().Do(req)
	assert.NoError(t, err)
	body, err := ioutil.ReadAll(got.Body)
	assert.NoError(t, err)
	assert.Equal(t, string(body), "")
}
