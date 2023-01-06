package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

func initializeHandlerFunctions() {
	http.HandleFunc("/", serveFile)
	http.HandleFunc("/increment", incrementCounter)
	http.HandleFunc("/hi", hiString)
}

func initializeHandler() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
}

func printUsage() {
	fmt.Println("Usage: ./golang-web-server option")
	fmt.Println("  options:")
	fmt.Println("    handler_functions")
	fmt.Println("    handler_directory")
	fmt.Println("  example:")
	fmt.Println("./golang-web-server handler_directory")
}

func contains(options []string, str string) bool {
	for _, option := range options {
		if option == str {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]
	if len(args) <= 0 || !contains(valid_args, args[0]) {
		printUsage()
	} else {

		switch args[0] {
		case "handler_functions":
			initializeHandlerFunctions()
		case "handler_directory":
			initializeHandler()
		}

		log.Fatal(http.ListenAndServe(":8081", nil))
	}
}
