package bollie

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (a *Application) getProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		a.Error(err)
		errResp := &ResponseError{
			Error:       "400",
			Description: "incorrect product id provided",
		}
		w.WriteHeader(http.StatusBadRequest)
		e := json.NewEncoder(w)
		e.SetIndent("", "  ")
		err = e.Encode(errResp)
		return
	}

	p, found := a.products.Get(id)

	if !found {
		errResp := &ResponseError{
			Error:       "404",
			Description: "product not found",
		}
		w.WriteHeader(http.StatusNotFound)
		e := json.NewEncoder(w)
		e.SetIndent("", "  ")
		err = e.Encode(errResp)
		return
	}

	resp := &Response{responseProduct{
		ID:    p.ID,
		Price: p.Price,
		Name:  p.Name,
	}}

	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	err = e.Encode(resp)

	if err != nil {
		a.Error(err)
	}
}

func (a *Application) purchaseProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		a.writeError(w, http.StatusBadRequest, errors.New("nil body received"))
	}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		a.writeError(w, http.StatusBadRequest, err)
		return
	}

	if len(data) == 0 {
		a.writeError(w, http.StatusBadRequest, errors.New("empty request body"))
	}

	var reqBody PurchaseBody

	if err = json.Unmarshal(data, &reqBody); err != nil {
		a.writeError(w, http.StatusBadRequest, err)
		return
	}

	quantity := reqBody.Quantity
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		a.writeError(w, http.StatusBadRequest, errors.New("incorrect product id provided"))
		return
	}

	p, found := a.products.Get(id)

	if !found {
		a.writeError(w, http.StatusNotFound, errors.New("product not found"))
		return
	}

	resp := &ResponsePurchase{
		Status:   "successfully purchased",
		Quantity: quantity,
		Product:  p,
	}

	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	err = e.Encode(resp)

	if err != nil {
		a.Error(err)
	}
}
