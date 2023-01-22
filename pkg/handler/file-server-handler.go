package handler

import "net/http"

// returns a handler that serves HTTP requests with the contents of the file system rooted at root
func InitializeHandler() {
	http.Handle("/", http.FileServer(http.Dir("../../static")))
}
