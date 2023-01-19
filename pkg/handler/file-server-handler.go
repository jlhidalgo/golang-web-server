package handler

import "net/http"

func InitializeHandler() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
}
