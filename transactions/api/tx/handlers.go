package tx

import (
	"fmt"

	middleware "github.com/go-openapi/runtime/middleware"

	"bitbucket.org/qwest-io/exp-api/gen/models"
	operationsbc "bitbucket.org/qwest-io/exp-api/gen/restapi/operations/bofa_chk"
	operations "bitbucket.org/qwest-io/exp-api/gen/restapi/operations/tx"
)

// GetBulkHandler ...
// [GET] tx/?month=YYYYMM&source=type
func GetBulkHandler(params operations.GetBulkParams) middleware.Responder {

	fmt.Printf("received month %v, source %v\n", params.Month, params.Source)

	payload := models.Transactions{&models.Transaction{Notes: "sample notes 3"}}
	return operations.NewGetBulkOK().WithPayload(payload)
}

// DeleteHandler ...
// [DELETE] tx/?month=YYYYMM&source=type
// Handler(params operations.GetBulkParams) middleware.Responder
func DeleteHandler(params operations.DeleteParams) middleware.Responder {

	fmt.Printf("received month %v, source %v\n", params.Month, params.Source)

	payload := models.Transactions{&models.Transaction{Notes: "delete sample notes 3"}}
	return operations.NewGetBulkOK().WithPayload(payload)
}

// GetHandler ...
// [GET] tx/{id}
func GetHandler(params operations.GetParams) middleware.Responder {

	fmt.Printf("transaction id %v\n", params.ID)

	payload := &models.Transaction{Notes: "sample transaction with id"}
	return operations.NewGetOK().WithPayload(payload)
}

// BofaChkImportHandler ...
// [POST] tx/bofa_chk/
// form-data: {csvfile, <file>}
func BofaChkImportHandler(params operationsbc.BofaChkImportParams) middleware.Responder {

	fmt.Printf("importing file %v\n", params.Csvfile)

	payload := models.TxData{&models.TxDataItems0{Sync: "csv file upload"}}
	return operationsbc.NewBofaChkImportOK().WithPayload(payload)
}

// BofaChkUpdateHandler ...
// [PUT] tx/bofa_chk/?lastId=123&source=bofa_chk
// form-data: {upfile, <file>}
func BofaChkUpdateHandler(params operationsbc.BofaChkUpdateParams) middleware.Responder {

	fmt.Printf("importing file %v from source %v with an lastId %v\n", params.Upfile, params.Source, params.LastID)

	payload := models.TxData{&models.TxDataItems0{Sync: "csv file update by upload"}}
	return operationsbc.NewBofaChkImportOK().WithPayload(payload)
}
