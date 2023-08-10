package main

import (
	"fmt"
	"log"
	"net/http"
)

// Hello returns a friendly greeting.
func HandleGetFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	addr := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleGetFunc)
	fmt.Printf("Listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
