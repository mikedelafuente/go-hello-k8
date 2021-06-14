package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/static", StaticServer)
	http.HandleFunc("/", HelloWorldServer)
	http.ListenAndServe(":8080", nil)
}

func StaticServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, static path!")
}

func HelloWorldServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Dynamic path")
}
