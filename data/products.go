package data

import "time"

type Products struct {
	Id          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}

func GetProduct() []*Products {
	return productLists
}

var productLists = []*Products{
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
