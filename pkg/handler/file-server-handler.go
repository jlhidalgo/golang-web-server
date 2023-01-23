package handler

import (
	"fmt"
	"net/http"
)

// returns a handler that serves HTTP requests with the contents of the file system rooted at root
func InitializeHandler() {
	fmt.Println("Initializing handler with FileServer")
	http.Handle("/", http.FileServer(http.Dir("../../static")))
	fmt.Println("Handler has been initialized")
}
