# Mini Web Server in Go

A very simple implementation of a **Web Server** in **Go** programming language.

## Description

This project implements a very basic Web Server using nothing else than the [Go Standard Library](https://pkg.go.dev/std). This was created for educational purposes in order to explore the basic functions that are provided by the standard library for *networking*, specifically the [net/http](https://pkg.go.dev/net/http@go1.19.5#pkg-overview) package. The server is capable of performing these operations: 
* listen to a specific port for incoming HTTP requests, and 
* process the requests and then reply with the corresponding HTTP responses.

The server can be launched in two different ways and this will determine how it will process the requests:
* handler functions: the application registers individual functions to process the requests depending on specific patterns.
* FileServer handler: Serves HTTP requests with the contents of the file system rooted at root.

This repository also includes a [Postman collection](/tests/golang-mini-web-server.postman_collection.json) that can be used for testing the functionality of the Web server.

## Getting started

### Dependencies

This is required for compiling the application:
* Go version 1.18.x or higher

This is optional, install it if you want to run the tests
* Postman

### Installing

1. Clone this repository
2. Make the repository folder your working directory
3. Initialize the module for the application by running these commands from the terminal (give a proper name to your module)
    ```bash
    go mod init module_name
    go mod vendor
    ```
4. Update the package reference in the *import* section inside the `cmd/web-server-golang-web-server.go` file with the *module_name* that you used in the previous step:
    ```golang
    "module_name/pkg/handler"
    ```
5. Compile the application to ensure everything works fine:
    ```bash
    cd cmd/web-server
    go build golang-web-server.go
    ```
    This should generate a binary filed named `golang-web-server` in the current working directory.

### Executing program

* Run the `golang-web-server` binary file directly:
    ```bash
    ./golang-web-server
    ```
    Expected outcome is:
    ```bash
    Initializing the handler functions...
    Handler functions have been initialized
    Listening on localhost:8081...
    ```
**Note**: By default, the application will be executed using handler functions.

* Run the application with FileServer handler:
    ```bash
    ./golang-web-server -use-handler-functions=false
    ```
    Expected outcome is:
    ```bash
    Initializing handler with FileServer
    Handler has been initialized
    Listening on localhost:8081...
    ```

### Testing


