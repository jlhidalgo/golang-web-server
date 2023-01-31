package mhttp

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type IServer interface {
	InitializeFileServer() error
	InitializeHandlerFunctions() error
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
func (fs ServerConfig) InitializeFileServer() error {
	fmt.Println("Initializing handler with FileServer")

	if !folderExists(fs.staticDirectory) {
		return errors.New("'static' folder does not exist: " + fs.staticDirectory)
	}

	http.Handle("/", http.FileServer(http.Dir(fs.staticDirectory)))
	fmt.Println("Handler has been initialized")
	return nil
}

// starts listenning the port and serves the requests
func (fs ServerConfig) ListenAndServe() {
	serverUrl := fs.url + ":" + fs.port
	fmt.Printf("Listening on %s...\n", serverUrl)
	log.Fatal(http.ListenAndServe(serverUrl, nil))
}

// Initializes all the handler functions
func (fs ServerConfig) InitializeHandlerFunctions() error {
	staticDirectory = fs.staticDirectory
	fmt.Println("Initializing the handler functions...")

	if !folderExists(fs.staticDirectory) {
		return errors.New("'static' folder does not exist: " + fs.staticDirectory)
	}

	http.HandleFunc("/", serveFile)
	http.HandleFunc("/increment", incrementCounter)
	http.HandleFunc("/hi", hiString)
	fmt.Println("Handler functions have been initialized")
	return nil
}

func folderExists(folderPath string) bool {
	fileInfo, err := os.Stat(folderPath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	if fileInfo != nil {
		return fileInfo.IsDir()
	}
	return false
}
