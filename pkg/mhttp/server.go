package mhttp

import (
	"fmt"
	"log"
	"net/http"
)

type IServer interface {
	InitializeFileServer()
	InitializeHandlerFunctions()
	ListenAndServe()
}

type ServerConfig struct {
	staticDirectory string
	url             string
	port            string
}

func NewServer(staticDir string, url string, port string) ServerConfig {
	var serverConfig = ServerConfig{
		staticDirectory: staticDir,
		url:             url,
		port:            port,
	}
	return serverConfig
}

// returns a handler that serves HTTP requests with the contents of the file system rooted at root
func (fs ServerConfig) InitializeFileServer() {
	fmt.Println("Initializing handler with FileServer")
	http.Handle("/", http.FileServer(http.Dir(fs.staticDirectory)))
	fmt.Println("Handler has been initialized")
}

// starts listenning the port and serves the requests
func (fs ServerConfig) ListenAndServe() {
	fmt.Println("Listening on localhost:8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// Initializes all the handler functions
func (fs ServerConfig) InitializeHandlerFunctions() {
	staticDirectory = fs.staticDirectory
	fmt.Println("Initializing the handler functions...")
	http.HandleFunc("/", serveFile)
	http.HandleFunc("/increment", incrementCounter)
	http.HandleFunc("/hi", hiString)
	fmt.Println("Handler functions have been initialized")
}
