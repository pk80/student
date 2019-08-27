// Code generated by go-swagger; DO NOT EDIT.

package cat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteBulkHandlerFunc turns a function with the right signature into a delete bulk handler
type DeleteBulkHandlerFunc func(DeleteBulkParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteBulkHandlerFunc) Handle(params DeleteBulkParams) middleware.Responder {
	return fn(params)
}

// DeleteBulkHandler interface for that can handle valid delete bulk params
type DeleteBulkHandler interface {
	Handle(DeleteBulkParams) middleware.Responder
}

// NewDeleteBulk creates a new http.Handler for the delete bulk operation
func NewDeleteBulk(ctx *middleware.Context, handler DeleteBulkHandler) *DeleteBulk {
	return &DeleteBulk{Context: ctx, Handler: handler}
}

/*DeleteBulk swagger:route DELETE /cat cat deleteBulk

Delete all category related transaction in bulk

*/
type DeleteBulk struct {
	Context *middleware.Context
	Handler DeleteBulkHandler
}

func (o *DeleteBulk) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteBulkParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
