package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Descrption string  `json:"description"`
	Price      float32 `json:"price"`
	SKU        string  `json:"sku"`
	CreatedOn  string  `json:"-"`
	UpdatedOn  string  `json:"-"`
	DeletedOn  string  `json:"-"`
}

type Products []Product

func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func GetProducts() Products {
	return productLists
}

var productLists = Products{
	Product{
		ID:         1,
		Name:       "Latte",
		Descrption: "Milky",
		Price:      2.45,
		SKU:        "abc323",
		CreatedOn:  time.Now().UTC().String(),
		UpdatedOn:  time.Now().UTC().String(),
	},
	Product{
		ID:         2,
		Name:       "Espresso",
		Descrption: "Milky",
		Price:      3.45,
		SKU:        "bbc324",
		CreatedOn:  time.Now().UTC().String(),
		UpdatedOn:  time.Now().UTC().String(),
	},
}
