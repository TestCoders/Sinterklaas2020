package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
)

func getCheapestProduct(products []Product) (Product, error) {
	var cheapest Product

	if products == nil || len(products) == 0 {
		return Product{}, errors.New("received empty slice of product")
	}

	for _, product := range products {
		if cheapest.Name == "" {
			cheapest = product
		} else if cheapest.Price > product.Price {
			cheapest = product
		}
	}

	return cheapest, nil
}

func (a *Application) writeErrorResponse(w http.ResponseWriter, err error, status int) {
	resp := &ErrorResponse{
		Status:      status,
		Description: err.Error(),
	}

	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	if e := a.errorLog.Output(2, trace); e != nil {
		a.errorLog.Println(err.Error())
	}

	w.WriteHeader(status)
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	err = e.Encode(resp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
