package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (a *Application) purchaseProductHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		a.writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if len(data) == 0 {
		a.writeErrorResponse(w, errors.New("empty request body"), http.StatusBadRequest)
		return
	}

	var purchaseRequest PurchaseRequest

	if err = json.Unmarshal(data, &purchaseRequest); err != nil {
		a.writeErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	quantity := purchaseRequest.Quantity

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	var products []Product
	for _, source := range a.sources {
		p, err := source.GetProduct(id)

		if err != nil {
			a.writeErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		products = append(products, p)
	}

	product, err := getCheapestProduct(products)

	if err != nil {
		a.writeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// Purchase product
	source := product.Source
	client := a.sources[source]

	resp, errResponse := client.PurchaseProduct(product.ID, quantity)

	if errResponse != nil {
		w.WriteHeader(errResponse.Status)
		e := json.NewEncoder(w)
		e.SetIndent("", "  ")
		e.Encode(errResponse)
		return
	}

	// Successfully purchased product
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	if err = e.Encode(resp); err != nil {
		a.writeErrorResponse(w, err, http.StatusInternalServerError)
	}
}
