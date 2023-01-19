package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}
var valid_args = []string{"handler_functions", "handler_directory"}

func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func hiString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}

func InitializeHandlerFunctions() {
	http.HandleFunc("/", serveFile)
	http.HandleFunc("/increment", incrementCounter)
	http.HandleFunc("/hi", hiString)
}