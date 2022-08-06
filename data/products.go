package data

import (
	"encoding/json"
	"fmt"
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

type Products []*Product

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}
func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func GetProducts() Products {
	return productLists
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productLists = append(productLists, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findByID(id)

	if err != nil {
		return err
	}

	p.ID = id
	productLists[pos] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findByID(id int) (*Product, int, error) {
	for i, data := range productLists {
		if data.ID == id {
			return data, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := productLists[len(productLists)-1]
	return lp.ID + 1
}

var productLists = Products{
	&Product{
		ID:         1,
		Name:       "Latte",
		Descrption: "Milky",
		Price:      2.45,
		SKU:        "abc323",
		CreatedOn:  time.Now().UTC().String(),
		UpdatedOn:  time.Now().UTC().String(),
	},
	&Product{
		ID:         2,
		Name:       "Espresso",
		Descrption: "Milky",
		Price:      3.45,
		SKU:        "bbc324",
		CreatedOn:  time.Now().UTC().String(),
		UpdatedOn:  time.Now().UTC().String(),
	},
}
