package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/devlorvn/go-project/data"
)

type KeyProduct struct{}

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func getProductId(r *http.Request) int {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err)
	}
	return id
}

func (p *Products) FindMany(rw http.ResponseWriter, r *http.Request) {

	prods := data.GetProducts()
	err := data.ToJson(prods, rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) FindOne(rw http.ResponseWriter, r *http.Request) {
	id := getProductId(r)

	prod, err := data.GetProductById(id)

	if err != nil {
		http.Error(rw, "error", http.StatusNotFound)
		return
	}

	err = data.ToJson(prod, rw)

	if err != nil {

	}
}

func (p *Products) CreateProduct(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	e := data.UpdateProduct(prod)

	if e == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if e != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := data.FromJson(prod, r.Body)

		if err != nil {
			p.l.Println("[Error] deserializing product", err)
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
