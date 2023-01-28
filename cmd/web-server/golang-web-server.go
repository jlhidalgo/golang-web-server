package main

import (
	"flag"

	"github.com/jlhidalgo/golang-web-server/pkg/mhttp"
)

const (
	staticDirectory string = "./web/static"
	url             string = "localhost"
	port            string = "8081"
)

var useHandlerFunctions bool

func main() {
	flag.BoolVar(&useHandlerFunctions, "use-handler-functions", true, "Runs the Web Browser with handler functions, use false to run it in FileServer mode.")
	flag.Parse()

	server := mhttp.NewServer(staticDirectory, url, port)

	if useHandlerFunctions {
		server.InitializeHandlerFunctions()
	} else {
		server.InitializeFileServer()
	}
	server.ListenAndServe()

}
