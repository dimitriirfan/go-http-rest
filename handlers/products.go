package handlers

import (
	"log"
	"ms-go/data"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (h *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.getProducts(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (h *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	productLists := data.GetProducts()
	err := productLists.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
