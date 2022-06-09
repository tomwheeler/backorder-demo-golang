package main

// call this via HTTP GET with a URL like:
//     http://localhost:8009/get-quantity-available?productId=5181250

import (
	"fmt"
	"net/http"
)

func quantityAvailableHandler(w http.ResponseWriter, r *http.Request) {
	_, ok := r.URL.Query()["productId"]
	if ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "42")
	} else {
		http.Error(w, "Missing required productId parameter.", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/get-quantity-available", quantityAvailableHandler)
	http.ListenAndServe(":8009", nil)
}
