package main

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func (a *Application) routes() http.Handler {
	middleware := alice.New(a.recoverPanic, a.logRequest)
	r := mux.NewRouter()
	r.HandleFunc("/purchase/{id}", a.purchaseProductHandler)
	return middleware.Then(r)
}
