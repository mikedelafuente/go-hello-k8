package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", StaticServer)
	http.ListenAndServe(":8080", nil)
}

func StaticServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
