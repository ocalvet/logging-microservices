package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleHome)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "{ 'succes': 'true', 'service': 'admin' }")
}
