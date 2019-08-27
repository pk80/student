// Code generated by go-swagger; DO NOT EDIT.

package cat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetBulkCatHandlerFunc turns a function with the right signature into a get bulk cat handler
type GetBulkCatHandlerFunc func(GetBulkCatParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBulkCatHandlerFunc) Handle(params GetBulkCatParams) middleware.Responder {
	return fn(params)
}

// GetBulkCatHandler interface for that can handle valid get bulk cat params
type GetBulkCatHandler interface {
	Handle(GetBulkCatParams) middleware.Responder
}

// NewGetBulkCat creates a new http.Handler for the get bulk cat operation
func NewGetBulkCat(ctx *middleware.Context, handler GetBulkCatHandler) *GetBulkCat {
	return &GetBulkCat{Context: ctx, Handler: handler}
}

/*GetBulkCat swagger:route GET /cat/ cat getBulkCat

Get all the Category records `catList` into the database `catData`

*/
type GetBulkCat struct {
	Context *middleware.Context
	Handler GetBulkCatHandler
}

func (o *GetBulkCat) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBulkCatParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}