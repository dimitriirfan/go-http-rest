package handlers

import (
	"log"
	"ms-go/data"
	"net/http"
	"regexp"
	"strconv"
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

	if r.Method == http.MethodPost {
		h.addProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {
		// get id from uri
		regex := regexp.MustCompile(`/([0-9]+)`)
		group := regex.FindAllStringSubmatch(r.URL.Path, -1)

		if len(group) != 1 {
			http.Error(w, "Invalid URI", http.StatusInternalServerError)
			return
		}

		idString := group[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(w, "String conversion error", http.StatusInternalServerError)
			return
		}

		h.l.Println(id)
		h.updateProduct(id, w, r)
		return

	}

	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (h *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle GET Products")

	productLists := data.GetProducts()
	err := productLists.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (h *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle POST Products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)

	h.l.Printf("Prod: %#v", prod)

}

func (h *Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle PUT Products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}

}
