package main

import (
	"fmt"
	"io/ioutil"
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
	fmt.Fprint(w, "Dynamic path... \n\n")
	callServer(w, "http://hello-app:8080")
	callServer(w, "http://world-app:8080")
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
