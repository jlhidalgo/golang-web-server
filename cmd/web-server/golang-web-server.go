package main

import (
	"flag"

	"github.com/jlhidalgo/golang-web-server/configs"
	"github.com/jlhidalgo/golang-web-server/pkg/mhttp"
)

var useHandlerFunctions bool

func main() {
	flag.BoolVar(&useHandlerFunctions, "use-handler-functions", true, "Runs the Web Browser with handler functions, use false to run it in FileServer mode.")
	flag.Parse()

	server := mhttp.NewServer(configs.SERVER_STATIC_DIRECTORY, configs.SERVER_URL, configs.SERVER_PORT)

	if useHandlerFunctions {
		server.InitializeHandlerFunctions()
	} else {
		server.InitializeFileServer()
	}
	server.ListenAndServe()

}
