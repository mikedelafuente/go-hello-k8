package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/static", StaticServer)
	http.HandleFunc("/", HelloWorldServer)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("hello-world-svc: listening on port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func StaticServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, static path!")
}

func HelloWorldServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Dynamic path... \n\n")
	helloserver := os.Getenv("HELLOSERVER")
	if helloserver == "" {
		helloserver = "http://hello-app:8080"
	}

	worldserver := os.Getenv("WORLDSERVER")
	if worldserver == "" {
		worldserver = "http://world-app:8080"
	}

	callServer(w, helloserver)
	callServer(w, worldserver)
}

func callServer(w http.ResponseWriter, serverName string) {
	fmt.Fprintf(w, "Calling : %v\n", serverName)
	response, err := http.Get(serverName)
	if err != nil {
		fmt.Fprintf(w, "The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Fprintf(w, "Data: \n %v\n\n", string(data))
	}
}
