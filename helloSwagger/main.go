package main

import (
	"flag"
	"log"

	"github.com/go-openapi/loads"
	"github.com/pk80/student/helloSwagger/apihandlers"
	"github.com/pk80/student/helloSwagger/restapi"
	"github.com/pk80/student/helloSwagger/restapi/operations"
)

var portFlag = flag.Int("port", 8080, "Port to run this service on")

func main() {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewHelloAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag

	// Set Handle
	api.GetHostnameHandler = operations.GetHostnameHandlerFunc(apihandlers.GetHostnameHandler)

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
