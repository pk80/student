package apihandlers

import (
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pk80/student/helloSwagger/restapi/operations"
)

// GetHostnameHandler ...
// Handlers
func GetHostnameHandler(params operations.GetHostnameParams) middleware.Responder {

	response, err := os.Hostname()
	if err != nil {
		return operations.NewGetHostnameOK().WithPayload("Failed  to retrieve hostname")
	}

	return operations.NewGetHostnameOK().WithPayload(response)
}
