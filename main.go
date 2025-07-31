package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Ping Pong")
	})
	http.ListenAndServe(":9090", nil)
}
