// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/pk80/student/transactions/restapi/operations"
	"github.com/pk80/student/transactions/restapi/operations/bofa_chk"
	"github.com/pk80/student/transactions/restapi/operations/cat"
	"github.com/pk80/student/transactions/restapi/operations/tx"
)

//go:generate swagger generate server --target ../../transactions --name Transactions --spec ../api.yml

func configureFlags(api *operations.TransactionsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TransactionsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	if api.BofaChkBofaChkImportHandler == nil {
		api.BofaChkBofaChkImportHandler = bofa_chk.BofaChkImportHandlerFunc(func(params bofa_chk.BofaChkImportParams) middleware.Responder {
			return middleware.NotImplemented("operation bofa_chk.BofaChkImport has not yet been implemented")
		})
	}
	if api.BofaChkBofaChkUpdateHandler == nil {
		api.BofaChkBofaChkUpdateHandler = bofa_chk.BofaChkUpdateHandlerFunc(func(params bofa_chk.BofaChkUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation bofa_chk.BofaChkUpdate has not yet been implemented")
		})
	}
	if api.TxDeleteHandler == nil {
		api.TxDeleteHandler = tx.DeleteHandlerFunc(func(params tx.DeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation tx.Delete has not yet been implemented")
		})
	}
	if api.CatDeleteBulkHandler == nil {
		api.CatDeleteBulkHandler = cat.DeleteBulkHandlerFunc(func(params cat.DeleteBulkParams) middleware.Responder {
			return middleware.NotImplemented("operation cat.DeleteBulk has not yet been implemented")
		})
	}
	if api.CatDeleteCatHandler == nil {
		api.CatDeleteCatHandler = cat.DeleteCatHandlerFunc(func(params cat.DeleteCatParams) middleware.Responder {
			return middleware.NotImplemented("operation cat.DeleteCat has not yet been implemented")
		})
	}
	if api.TxGetHandler == nil {
		api.TxGetHandler = tx.GetHandlerFunc(func(params tx.GetParams) middleware.Responder {
			return middleware.NotImplemented("operation tx.Get has not yet been implemented")
		})
	}
	if api.TxGetBulkHandler == nil {
		api.TxGetBulkHandler = tx.GetBulkHandlerFunc(func(params tx.GetBulkParams) middleware.Responder {
			return middleware.NotImplemented("operation tx.GetBulk has not yet been implemented")
		})
	}
	if api.CatGetBulkCatHandler == nil {
		api.CatGetBulkCatHandler = cat.GetBulkCatHandlerFunc(func(params cat.GetBulkCatParams) middleware.Responder {
			return middleware.NotImplemented("operation cat.GetBulkCat has not yet been implemented")
		})
	}
	if api.CatGetCatHandler == nil {
		api.CatGetCatHandler = cat.GetCatHandlerFunc(func(params cat.GetCatParams) middleware.Responder {
			return middleware.NotImplemented("operation cat.GetCat has not yet been implemented")
		})
	}
	if api.CatImportCatHandler == nil {
		api.CatImportCatHandler = cat.ImportCatHandlerFunc(func(params cat.ImportCatParams) middleware.Responder {
			return middleware.NotImplemented("operation cat.ImportCat has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
