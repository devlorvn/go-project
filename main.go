package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Ping Pong")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ooops", http.StatusBadRequest)
			return
		}
		log.Printf("Data %s\n", d)
		fmt.Fprintf(w, "Hello %s", d)
	})

	http.ListenAndServe(":9090", nil)
}
