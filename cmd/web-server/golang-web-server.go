package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jlhidalgo/golang-web-server/pkg/handler"
)

const staticDirectory string = "./web/static"

var useHandlerFunctions bool

func main() {
	flag.BoolVar(&useHandlerFunctions, "use-handler-functions", true, "Runs the Web Browser with handler functions")
	flag.Parse()

	if useHandlerFunctions {
		handler.InitializeHandlerFunctions(staticDirectory)
	} else {
		handler.InitializeHandler(staticDirectory)
	}

	fmt.Println("Listening on localhost:8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
