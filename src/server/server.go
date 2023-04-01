package server

import (
	"fmt"
	"log"
	"net/http"
)

func Serve() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

func handler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Hello!")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
