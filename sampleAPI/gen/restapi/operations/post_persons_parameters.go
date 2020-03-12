// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "github.com/pk80/student/sampleAPI/gen/models"
)

// NewPostPersonsParams creates a new PostPersonsParams object
// no default values defined in spec.
func NewPostPersonsParams() PostPersonsParams {

	return PostPersonsParams{}
}

// PostPersonsParams contains all the bound params for the post persons operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostPersons
type PostPersonsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The person to create
	  In: body
	*/
	Person *models.Person
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostPersonsParams() beforehand.
func (o *PostPersonsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Person
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("person", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Person = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
