package data

import (
	"fmt"
)

var ErrProductNotFound = fmt.Errorf("Product not found")

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"desc"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
}

type Products []*Product

func GetProducts() Products {
	return productLists
}

func GetProductById(id int) (*Product, error) {
	i := findIndexByProductId(id)

	if i == -1 {
		return nil, ErrProductNotFound
	}

	return productLists[i], nil
}

func AddProduct(p Product) {
	p.Id = getNextId()
	productLists = append(productLists, &p)
}

func UpdateProduct(p Product) error {
	i := findIndexByProductId(p.Id)
	if i == -1 {
		return ErrProductNotFound
	}

	productLists[i] = &p
	return nil
}

func getNextId() int {
	lastP := productLists[len(productLists)-1]
	return lastP.Id + 1
}

func DeleteProduct(id int) error {
	i := findIndexByProductId(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productLists = append(productLists[:i], productLists[i+1])

	return nil
}

func findIndexByProductId(id int) int {
	for i, p := range productLists {
		if p.Id == id {
			return i
		}
	}
	return -1
}

var productLists = []*Product{
	{
		Id:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.34,
		SKU:         "pro001",
	},
	{
		Id:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       2.34,
		SKU:         "pro001",
	},
}
