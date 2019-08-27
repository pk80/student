// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/pk80/student/transactions/restapi/operations/bofa_chk"
	"github.com/pk80/student/transactions/restapi/operations/cat"
	"github.com/pk80/student/transactions/restapi/operations/tx"
)

// NewTransactionsAPI creates a new Transactions instance
func NewTransactionsAPI(spec *loads.Document) *TransactionsAPI {
	return &TransactionsAPI{
		handlers:              make(map[string]map[string]http.Handler),
		formats:               strfmt.Default,
		defaultConsumes:       "application/json",
		defaultProduces:       "application/json",
		customConsumers:       make(map[string]runtime.Consumer),
		customProducers:       make(map[string]runtime.Producer),
		ServerShutdown:        func() {},
		spec:                  spec,
		ServeError:            errors.ServeError,
		BasicAuthenticator:    security.BasicAuth,
		APIKeyAuthenticator:   security.APIKeyAuth,
		BearerAuthenticator:   security.BearerAuth,
		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,
		JSONProducer:          runtime.JSONProducer(),
		BofaChkBofaChkImportHandler: bofa_chk.BofaChkImportHandlerFunc(func(params bofa_chk.BofaChkImportParams) middleware.Responder {
			return middleware.NotImplemented("operation BofaChkBofaChkImport has not yet been implemented")
		}),
		BofaChkBofaChkUpdateHandler: bofa_chk.BofaChkUpdateHandlerFunc(func(params bofa_chk.BofaChkUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation BofaChkBofaChkUpdate has not yet been implemented")
		}),
		TxDeleteHandler: tx.DeleteHandlerFunc(func(params tx.DeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation TxDelete has not yet been implemented")
		}),
		CatDeleteBulkHandler: cat.DeleteBulkHandlerFunc(func(params cat.DeleteBulkParams) middleware.Responder {
			return middleware.NotImplemented("operation CatDeleteBulk has not yet been implemented")
		}),
		CatDeleteCatHandler: cat.DeleteCatHandlerFunc(func(params cat.DeleteCatParams) middleware.Responder {
			return middleware.NotImplemented("operation CatDeleteCat has not yet been implemented")
		}),
		TxGetHandler: tx.GetHandlerFunc(func(params tx.GetParams) middleware.Responder {
			return middleware.NotImplemented("operation TxGet has not yet been implemented")
		}),
		TxGetBulkHandler: tx.GetBulkHandlerFunc(func(params tx.GetBulkParams) middleware.Responder {
			return middleware.NotImplemented("operation TxGetBulk has not yet been implemented")
		}),
		CatGetBulkCatHandler: cat.GetBulkCatHandlerFunc(func(params cat.GetBulkCatParams) middleware.Responder {
			return middleware.NotImplemented("operation CatGetBulkCat has not yet been implemented")
		}),
		CatGetCatHandler: cat.GetCatHandlerFunc(func(params cat.GetCatParams) middleware.Responder {
			return middleware.NotImplemented("operation CatGetCat has not yet been implemented")
		}),
		CatImportCatHandler: cat.ImportCatHandlerFunc(func(params cat.ImportCatParams) middleware.Responder {
			return middleware.NotImplemented("operation CatImportCat has not yet been implemented")
		}),
	}
}

/*TransactionsAPI This is a sample api for all sorts of transactions. And all transactions are categorized based upon type of account operation. >>TODO<< */
type TransactionsAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for a "multipart/form-data" mime type
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// BofaChkBofaChkImportHandler sets the operation handler for the bofa chk import operation
	BofaChkBofaChkImportHandler bofa_chk.BofaChkImportHandler
	// BofaChkBofaChkUpdateHandler sets the operation handler for the bofa chk update operation
	BofaChkBofaChkUpdateHandler bofa_chk.BofaChkUpdateHandler
	// TxDeleteHandler sets the operation handler for the delete operation
	TxDeleteHandler tx.DeleteHandler
	// CatDeleteBulkHandler sets the operation handler for the delete bulk operation
	CatDeleteBulkHandler cat.DeleteBulkHandler
	// CatDeleteCatHandler sets the operation handler for the delete cat operation
	CatDeleteCatHandler cat.DeleteCatHandler
	// TxGetHandler sets the operation handler for the get operation
	TxGetHandler tx.GetHandler
	// TxGetBulkHandler sets the operation handler for the get bulk operation
	TxGetBulkHandler tx.GetBulkHandler
	// CatGetBulkCatHandler sets the operation handler for the get bulk cat operation
	CatGetBulkCatHandler cat.GetBulkCatHandler
	// CatGetCatHandler sets the operation handler for the get cat operation
	CatGetCatHandler cat.GetCatHandler
	// CatImportCatHandler sets the operation handler for the import cat operation
	CatImportCatHandler cat.ImportCatHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *TransactionsAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *TransactionsAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *TransactionsAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *TransactionsAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *TransactionsAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *TransactionsAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *TransactionsAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the TransactionsAPI
func (o *TransactionsAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BofaChkBofaChkImportHandler == nil {
		unregistered = append(unregistered, "bofa_chk.BofaChkImportHandler")
	}

	if o.BofaChkBofaChkUpdateHandler == nil {
		unregistered = append(unregistered, "bofa_chk.BofaChkUpdateHandler")
	}

	if o.TxDeleteHandler == nil {
		unregistered = append(unregistered, "tx.DeleteHandler")
	}

	if o.CatDeleteBulkHandler == nil {
		unregistered = append(unregistered, "cat.DeleteBulkHandler")
	}

	if o.CatDeleteCatHandler == nil {
		unregistered = append(unregistered, "cat.DeleteCatHandler")
	}

	if o.TxGetHandler == nil {
		unregistered = append(unregistered, "tx.GetHandler")
	}

	if o.TxGetBulkHandler == nil {
		unregistered = append(unregistered, "tx.GetBulkHandler")
	}

	if o.CatGetBulkCatHandler == nil {
		unregistered = append(unregistered, "cat.GetBulkCatHandler")
	}

	if o.CatGetCatHandler == nil {
		unregistered = append(unregistered, "cat.GetCatHandler")
	}

	if o.CatImportCatHandler == nil {
		unregistered = append(unregistered, "cat.ImportCatHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *TransactionsAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *TransactionsAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *TransactionsAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *TransactionsAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *TransactionsAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *TransactionsAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the transactions API
func (o *TransactionsAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *TransactionsAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/tx/bofa-chk"] = bofa_chk.NewBofaChkImport(o.context, o.BofaChkBofaChkImportHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/tx/bofa-chk"] = bofa_chk.NewBofaChkUpdate(o.context, o.BofaChkBofaChkUpdateHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/tx"] = tx.NewDelete(o.context, o.TxDeleteHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/cat"] = cat.NewDeleteBulk(o.context, o.CatDeleteBulkHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/cat/{id}"] = cat.NewDeleteCat(o.context, o.CatDeleteCatHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tx/{id}"] = tx.NewGet(o.context, o.TxGetHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tx"] = tx.NewGetBulk(o.context, o.TxGetBulkHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/cat"] = cat.NewGetBulkCat(o.context, o.CatGetBulkCatHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/cat/{id}"] = cat.NewGetCat(o.context, o.CatGetCatHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/cat"] = cat.NewImportCat(o.context, o.CatImportCatHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *TransactionsAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *TransactionsAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *TransactionsAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *TransactionsAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}