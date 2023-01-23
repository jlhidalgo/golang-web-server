package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/jlhidalgo/golang-web-server/pkg/handler"
)

var useHandlerFunctions bool

func main() {
	flag.BoolVar(&useHandlerFunctions, "use-handler-functions", true, "Runs the Web Browser with handler functions")
	flag.Parse()

	if useHandlerFunctions {
		handler.InitializeHandlerFunctions()
	} else {
		handler.InitializeHandler()
	}

	log.Fatal(http.ListenAndServe(":8081", nil))
}
