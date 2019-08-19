// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/pk80/student/08_swagger/restapi/operations"
	"github.com/pk80/student/08_swagger/restapi/operations/deltd"
	"github.com/pk80/student/08_swagger/restapi/operations/gettd"
	"github.com/pk80/student/08_swagger/restapi/operations/posttd"
	"github.com/pk80/student/08_swagger/restapi/operations/puttd"
)

//go:generate swagger generate server --target ../../08_swagger --name TodoList --spec ../swagger.yml

func configureFlags(api *operations.TodoListAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TodoListAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.PosttdAddOneHandler == nil {
		api.PosttdAddOneHandler = posttd.AddOneHandlerFunc(func(params posttd.AddOneParams) middleware.Responder {
			return middleware.NotImplemented("operation posttd.AddOne has not yet been implemented")
		})
	}
	if api.DeltdDestroyOneHandler == nil {
		api.DeltdDestroyOneHandler = deltd.DestroyOneHandlerFunc(func(params deltd.DestroyOneParams) middleware.Responder {
			return middleware.NotImplemented("operation deltd.DestroyOne has not yet been implemented")
		})
	}
	if api.GettdFindTodosHandler == nil {
		api.GettdFindTodosHandler = gettd.FindTodosHandlerFunc(func(params gettd.FindTodosParams) middleware.Responder {
			return middleware.NotImplemented("operation gettd.FindTodos has not yet been implemented")
		})
	}
	if api.PuttdUpdateOneHandler == nil {
		api.PuttdUpdateOneHandler = puttd.UpdateOneHandlerFunc(func(params puttd.UpdateOneParams) middleware.Responder {
			return middleware.NotImplemented("operation puttd.UpdateOne has not yet been implemented")
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
