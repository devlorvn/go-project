package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"desc"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)

	return e.Decode(p)
}

type Products []*Product

func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(p)
}

func GetProduct() Products {
	return productLists
}

func AddProduct(p *Product) {
	p.Id = getNextId()
	productLists = append(productLists, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)

	if err != nil {
		return err
	}
	p.Id = id
	productLists[pos] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productLists {
		if p.Id == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextId() int {
	lastP := productLists[len(productLists)-1]
	return lastP.Id + 1
}

var productLists = []*Product{
	{
		Id:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.34,
		SKU:         "pro001",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
	{
		Id:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       2.34,
		SKU:         "pro001",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
}
