package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", StaticServer)
	log.Printf("world-svc: listening on port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func StaticServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "World")
}
