// Code generated by go-swagger; DO NOT EDIT.

package tx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetBulkHandlerFunc turns a function with the right signature into a get bulk handler
type GetBulkHandlerFunc func(GetBulkParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBulkHandlerFunc) Handle(params GetBulkParams) middleware.Responder {
	return fn(params)
}

// GetBulkHandler interface for that can handle valid get bulk params
type GetBulkHandler interface {
	Handle(GetBulkParams) middleware.Responder
}

// NewGetBulk creates a new http.Handler for the get bulk operation
func NewGetBulk(ctx *middleware.Context, handler GetBulkHandler) *GetBulk {
	return &GetBulk{Context: ctx, Handler: handler}
}

/*GetBulk swagger:route GET /tx tx getBulk

Get transactions for the given `month`, `year` and `source` from `txData`

*/
type GetBulk struct {
	Context *middleware.Context
	Handler GetBulkHandler
}

func (o *GetBulk) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBulkParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
