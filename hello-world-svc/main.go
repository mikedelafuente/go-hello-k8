package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/compute/metadata"
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

	callServerWithAuth(w, helloserver)
	callServerWithAuth(w, worldserver)
}

// func callServer(w http.ResponseWriter, serverName string) {
// 	fmt.Fprintf(w, "Calling : %v\n", serverName)
// 	response, err := http.Get(serverName)
// 	if err != nil {
// 		fmt.Fprintf(w, "The HTTP request failed with error %s\n", err)
// 	} else {
// 		data, _ := ioutil.ReadAll(response.Body)
// 		fmt.Fprintf(w, "Data: \n %v\n\n", string(data))
// 	}
// }

func callServerWithAuth(w http.ResponseWriter, serverName string) {
	fmt.Fprintf(w, "Calling : %v\n", serverName)
	tokenURL := fmt.Sprintf("/instance/service-accounts/default/identity?audience=%s", serverName)
	fmt.Fprintf(w, "Getting token : %v\n", tokenURL)

	idToken, err := metadata.Get(tokenURL)
	if err != nil {
		fmt.Fprintf(w, "metadata.Get: failed to query id_token: %+v", err)
		return
	}

	req, err := http.NewRequest("GET", serverName, nil)
	if err != nil {
		fmt.Fprintf(w, "Error creating new request: %+v", err)
		return
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", idToken))
	fmt.Fprintf(w, "Got id token: Bearer %s", idToken)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(w, "The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Fprintf(w, "Data: \n %v\n\n", string(data))
	}
}
