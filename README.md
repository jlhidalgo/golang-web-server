# Mini Web Server in Go

A very simple implementation of a **Web Server** in **Go** programming language.

## Description

This project implements a very basic Web Server using nothing else than the [Go Standard Library](https://pkg.go.dev/std). This was created for educational purposes in order to explore the basic functions that are provided by the standard library for *networking*, specifically the [net/http](https://pkg.go.dev/net/http@go1.19.5#pkg-overview) package. The server is capable of performing these operations: 
* listen to a specific port for incoming HTTP requests, and 
* process the requests and then reply with the corresponding HTTP responses.

The application can serve the requests by two different ways:
* handler functions: the application registers individual functions to process the requests depending on specific patterns.
* FileServer handler: Serves HTTP requests with the contents of the file system rooted at root.

This repository also includes a [Postman collection](/tests/golang-mini-web-server.postman_collection.json) that can be used for testing the functionality of the Web server.

## Getting started

### Dependencies

This is required for building and compiling the application:
* Go version 1.18.x or higher

Optionally you will need Postman if you want to run the tests provided [here](/tests/golang-mini-web-server.postman_collection.json).
* Postman

### Installing

1. Clone this repository
2. Make the repository folder your working directory
3. Initialize the module for the application by running these commands from the terminal (give a proper name to your module)
    ```bash
    go mod init module_name
    go mod vendor
    ```
4. Update the package reference in the *import* section inside the `cmd/web-server/golang-web-server.go` file with the *module_name* that you used in the previous step:
    ```golang
    "module_name/pkg/handler"
    ```
5. Compile and build the application to ensure everything works fine.
    
    In Unix-like systems just execute this script:
    ```bash
    ./scripts/build.sh
    ```

    Or you can just execute this command as well:
    ```bash
    go build -o ./bin/web-server cmd/web-server/golang-web-server.go
    ```
    This should generate a binary file named `web-server` in the `bin` directory.

### Executing the program

* The simplest way is to run the `web-server` binary file directly:
    ```bash
    ./bin/web-server
    ```
    Expected outcome is:
    ```bash
    Initializing the handler functions...
    Handler functions have been initialized
    Listening on localhost:8081...
    ```
**Note**: By default, the application will be executed using handler functions and will listen to port 8081.

* Alternatively you can also run the application with FileServer handler:
    ```bash
    ./bin/web-server -use-handler-functions=false
    ```
    Expected outcome is:
    ```bash
    Initializing handler with FileServer
    Handler has been initialized
    Listening on localhost:8081...
    ```

### Testing

There are two approaches for interacting and testing the application, this can be done through a Web browser or Postman.

#### Web browser

Open a Web browser and go to http://localhost:8081. The `index` page should be loaded if everything is working as expected. Examine the source code for the requests that can be processed by the Web server, these vary depending on how the application is serving the requests (handler functions or FileServer handler).

#### Postman

Import the [collection](./tests/golang-mini-web-server.postman_collection.json) into Postman and then open the `Golang Mini Web Server` collection. The requests of this collection are grouped into two different folders, one is for handler functions implementation and the other one for FileServer handler implementation. You can run the requests within the folder that matches the current mode of the Web server.

## License
This project is licensed under the MIT License - see the LICENSE file for details

## Acknowledgments
Inspiration, code snippets, etc.

* [Creating a simple Web server with Golang](https://tutorialedge.net/golang/creating-simple-web-server-with-golang/)
* [A simple README.md template](https://gist.github.com/DomPizzie/7a5ff55ffa9081f2de27c315f5018afc)
* [Standard Go project layout](https://github.com/golang-standards/project-layout)