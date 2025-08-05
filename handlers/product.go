package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/devlorvn/go-project/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetProduct(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.AddProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.UpdateProduct(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProduct()
	err := lp.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println(r.URL.Path)
	re := regexp.MustCompile(`^/product/([0-9]+)$`)
	g := re.FindAllStringSubmatch(r.URL.Path, -1)
	if len(g) != 1 {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	idString := g[0][1]

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
		return
	}

	prod := &data.Product{}

	er := prod.FromJSON(r.Body)

	if er != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	e := data.UpdateProduct(id, prod)

	if e == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if e != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(fmt.Sprintf("Product %d updated!", id)))

}
