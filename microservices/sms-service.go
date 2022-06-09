package main

// call this via HTTP GET with a URL like:
//     http://localhost:8008/send-text?message=Hello+there&number=2138675309

import (
	"fmt"
	"math/rand"
	"net/http"
)

func sendTextHandler(w http.ResponseWriter, r *http.Request) {
	_, ok := r.URL.Query()["message"]
	_, ok2 := r.URL.Query()["number"]
	if ok && ok2 {
		// pretend we actually sent something here
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, generateMessageId())
	} else {
		http.Error(w, "Missing required parameter.", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/send-text", sendTextHandler)
	http.ListenAndServe(":8008", nil)
}

func generateMessageId() string {
	numRandChars := 10
	randChars := make([]byte, numRandChars)
	for i := range randChars {
		allowedChars := "0123456789ABCDEF"
		randChars[i] = allowedChars[rand.Intn(len(allowedChars))]
	}
	return "SMSID" + string(randChars)
}
