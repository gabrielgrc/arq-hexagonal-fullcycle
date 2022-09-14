package handler

import (
	"encoding/json"
	"github.com/codeedu/go-hexagonal/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle("/product/{id}/enable", n.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("GET", "OPTIONS")
	r.Handle("/product/{id}/disable", n.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("GET", "OPTIONS")
	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")
	r.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(product)
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

