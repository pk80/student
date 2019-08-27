package main

// reference: https://ops.tips/blog/a-swagger-golang-hello-world/

import (
	"log"

	"github.com/go-openapi/loads"

	"bitbucket.org/qwest-io/exp-api/api/tx"
	"bitbucket.org/qwest-io/exp-api/gen/restapi"
	operations "bitbucket.org/qwest-io/exp-api/gen/restapi/operations"
	operationsbc "bitbucket.org/qwest-io/exp-api/gen/restapi/operations/bofa_chk"
	operationstx "bitbucket.org/qwest-io/exp-api/gen/restapi/operations/tx"
)

func main() {

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewExpAPI(swaggerSpec)
	server := restapi.NewServer(api)

	server.Port = 8080

	api.TxGetBulkHandler = operationstx.GetBulkHandlerFunc(tx.GetBulkHandler)
	api.TxDeleteHandler = operationstx.DeleteHandlerFunc(tx.DeleteHandler)
	api.TxGetHandler = operationstx.GetHandlerFunc(tx.GetHandler)
	api.BofaChkBofaChkImportHandler = operationsbc.BofaChkImportHandlerFunc(tx.BofaChkImportHandler)
	api.BofaChkBofaChkUpdateHandler = operationsbc.BofaChkUpdateHandlerFunc(tx.BofaChkUpdateHandler)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
