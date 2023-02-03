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

	if err := validateFolder(fs.staticDirectory); err != nil {
		return errors.New("couldn't validate static folder, error: " + err.Error())
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

	if err := validateFolder(fs.staticDirectory); err != nil {
		return errors.New("couldn't validate static folder, error: " + err.Error())
	}

	http.HandleFunc("/", serveFile)
	http.HandleFunc("/increment", incrementCounter)
	http.HandleFunc("/hi", hiString)
	fmt.Println("Handler functions have been initialized")
	return nil
}

func validateFolder(folderPath string) error {
	fileInfo, err := os.Stat(folderPath)

	if err != nil {
		return err
	}

	if fileInfo != nil && !fileInfo.IsDir() {
		return fmt.Errorf("%s : is not a directory", folderPath)
	}

	return nil
}
