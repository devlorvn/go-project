package main

import (
	"log"
	"net/http"
	"os"

	"github.com/devlorvn/go-project/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	http.HandleFunc("/hello", hh.ServeHTTP)

	http.ListenAndServe(":9090", nil)
}
