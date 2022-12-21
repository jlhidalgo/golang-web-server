package main

import (
	"fmt"
	"log"
	"net/http"
)

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func hiString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}

func main() {

	http.HandleFunc("/", echoString)

	http.HandleFunc("/hi", hiString)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
