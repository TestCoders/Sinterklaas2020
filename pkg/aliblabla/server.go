package aliblabla

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
	"time"
)

func (a *Application) StartServer() {
	middleware := alice.New(a.logRequest)

	r := mux.NewRouter()
	r.HandleFunc("/cadeau/{id}", a.getProductHandler).Methods("GET")
	r.HandleFunc("/cadeau/{id}", a.purchaseProductHandler).Methods("POST")

	port := ":3003"

	srv := &http.Server{
		Handler:      middleware.Then(r),
		Addr:         port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	a.infoLog.Printf("starting aliblabla server at %v\n", port)
	a.errorLog.Fatal(srv.ListenAndServe())
}
