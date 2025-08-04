package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/devlorvn/go-project/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProduct()
	d, err := json.Marshal(lp)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)

	}

	rw.Write(d)
}
