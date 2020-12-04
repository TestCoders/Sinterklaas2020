package bollie

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
	"os"
	"testing"
)

func readFile(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

func TestApplication_getProductHandler(t *testing.T) {
	testProductResponse, err := readFile("mocks/product_response.json")

	if err != nil {
		t.Fatal("unexpected error setting up test request")
	}

	app := &Application{
		errorLog: log.New(os.Stdout, "ERROR\t", 0),
		infoLog:  log.New(os.Stdout, "INFO\t", 0),
		products: &inMemoryProductRepository{
			Products: map[int]product{
				1: {
					ID:    1,
					Name:  "c0ffee",
					Price: 13.37,
				},
			},
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/cadeau/{id}", app.getProductHandler)
	r, _ := http.NewRequest(http.MethodGet, "/cadeau/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	result := w.Result()

	body, err := ioutil.ReadAll(result.Body)

	assert.NoError(t, err)
	assert.Equal(t, 200, result.StatusCode)
	assert.Equal(t, testProductResponse, body)
}

func TestApplication_getProductHandler_productNotFound(t *testing.T) {
	testResponse, err := readFile("mocks/product_not_found_response.json")

	if err != nil {
		t.Fatal("unexpected error setting up test request")
	}

	app := &Application{
		errorLog: log.New(os.Stdout, "ERROR\t", 0),
		infoLog:  log.New(os.Stdout, "INFO\t", 0),
		products: &inMemoryProductRepository{
			Products: map[int]product{
				1: {
					ID:    1,
					Name:  "c0ffee",
					Price: 13.37,
				},
			},
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/cadeau/{id}", app.getProductHandler)
	r, _ := http.NewRequest(http.MethodGet, "/cadeau/2", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	result := w.Result()

	body, err := ioutil.ReadAll(result.Body)

	assert.NoError(t, err)
	assert.Equal(t, 404, result.StatusCode)
	assert.Equal(t, testResponse, body)
}

func TestApplication_getProductHandler_badRequest(t *testing.T) {
	testResponse, err := readFile("mocks/bad_request_response.json")

	if err != nil {
		t.Fatal("unexpected error setting up test request")
	}

	app := &Application{
		errorLog: log.New(os.Stdout, "ERROR\t", 0),
		infoLog:  log.New(os.Stdout, "INFO\t", 0),
		products: &inMemoryProductRepository{
			Products: map[int]product{
				1: {
					ID:    1,
					Name:  "c0ffee",
					Price: 13.37,
				},
			},
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/cadeau/{id}", app.getProductHandler)
	r, _ := http.NewRequest(http.MethodGet, "/cadeau/bami", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	result := w.Result()

	body, err := ioutil.ReadAll(result.Body)

	assert.NoError(t, err)
	assert.Equal(t, 400, result.StatusCode)
	assert.Equal(t, testResponse, body)
}

func TestApplication_purchaseProductHandler(t *testing.T) {
	testResponse, err := readFile("mocks/purchase_product_response.json")

	if err != nil {
		t.Fatal("unexpected error setting up test request")
	}

	app := &Application{
		errorLog: log.New(os.Stdout, "ERROR\t", 0),
		infoLog:  log.New(os.Stdout, "INFO\t", 0),
		products: &inMemoryProductRepository{
			Products: map[int]product{
				1: {
					ID:    1,
					Name:  "c0ffee",
					Price: 13.37,
				},
			},
		},
	}

	reqBody := &PurchaseBody{
		Quantity:  1,
		ProductID: 1,
	}

	data, err := json.Marshal(reqBody)

	if err != nil {
		t.Fatal(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/cadeau/{id}", app.purchaseProductHandler)
	r, _ := http.NewRequest(http.MethodGet, "/cadeau/1", bytes.NewBuffer(data))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	result := w.Result()

	body, err := ioutil.ReadAll(result.Body)

	assert.NoError(t, err)
	assert.Equal(t, 200, result.StatusCode)
	fmt.Println(string(body))
	fmt.Println(string(testResponse))
	assert.Equal(t, testResponse, body)
}

func TestApplication_purchaseProductHandler_productNotFound(t *testing.T) {
	testResponse, err := readFile("mocks/product_not_found_response.json")

	if err != nil {
		t.Fatal("unexpected error setting up test request")
	}

	app := &Application{
		errorLog: log.New(os.Stdout, "ERROR\t", 0),
		infoLog:  log.New(os.Stdout, "INFO\t", 0),
		products: &inMemoryProductRepository{
			Products: map[int]product{
				1: {
					ID:    1,
					Name:  "c0ffee",
					Price: 13.37,
				},
			},
		},
	}

	reqBody := &PurchaseBody{
		Quantity:  1,
		ProductID: 1,
	}

	data, err := json.Marshal(reqBody)

	if err != nil {
		t.Fatal(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/cadeau/{id}", app.purchaseProductHandler)
	r, _ := http.NewRequest(http.MethodGet, "/cadeau/2", bytes.NewBuffer(data))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	result := w.Result()

	body, err := ioutil.ReadAll(result.Body)

	assert.NoError(t, err)
	assert.Equal(t, 404, result.StatusCode)
	fmt.Println(string(body))
	assert.Equal(t, testResponse, body)
}

func TestApplication_purchaseProductHandler_badRequest(t *testing.T) {
	testResponse, err := readFile("mocks/bad_request_response.json")

	if err != nil {
		t.Fatal("unexpected error setting up test request")
	}

	app := &Application{
		errorLog: log.New(os.Stdout, "ERROR\t", 0),
		infoLog:  log.New(os.Stdout, "INFO\t", 0),
		products: &inMemoryProductRepository{
			Products: map[int]product{
				1: {
					ID:    1,
					Name:  "c0ffee",
					Price: 13.37,
				},
			},
		},
	}

	reqBody := &PurchaseBody{
		Quantity:  1,
		ProductID: 1,
	}

	data, err := json.Marshal(reqBody)

	if err != nil {
		t.Fatal(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/cadeau/{id}", app.purchaseProductHandler)
	r, _ := http.NewRequest(http.MethodGet, "/cadeau/bami", bytes.NewBuffer(data))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	result := w.Result()

	body, err := ioutil.ReadAll(result.Body)

	assert.NoError(t, err)
	assert.Equal(t, 400, result.StatusCode)
	fmt.Println(string(body))
	assert.Equal(t, testResponse, body)
}
