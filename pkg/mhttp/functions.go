package mhttp

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}
var staticDirectory string

// Replies to the request with the contents of the named file or directory
func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, staticDirectory)
}

// Increments a counter every time the '/increment' page is requested
// additionally the counter is replied back to the web browser
func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprint(w, strconv.Itoa(counter))
	mutex.Unlock()
}

// Just replies with the word 'Hi' as the body of the page
func hiString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}
