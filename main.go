package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/devlorvn/go-project/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	ph := handlers.NewProducts(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.FindMany)
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.FindMany)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.CreateProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareProductValidation)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
			os.IsExist(err)
		}
		l.Println("Starting server on port 9090")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	// signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(ctx)
}
