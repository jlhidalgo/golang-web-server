package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang-web-server.com/web-server/handler"
	"golang-web-server.com/web-server/utils"
)

var valid_args = []string{"handler_functions", "handler_directory"}

func printUsage() {
	fmt.Println("Usage: ./golang-web-server option")
	fmt.Println("  options:")
	fmt.Println("    handler_functions")
	fmt.Println("    handler_directory")
	fmt.Println("  example:")
	fmt.Println("./golang-web-server handler_directory")
}

func main() {
	args := os.Args[1:]
	if len(args) <= 0 || !utils.Contains(valid_args, args[0]) {
		printUsage()
	} else {

		switch args[0] {
		case "handler_functions":
			handler.InitializeHandlerFunctions()
		case "handler_directory":
			handler.InitializeHandler()
		}

		log.Fatal(http.ListenAndServe(":8081", nil))
	}
}
