// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/pk80/student/swagAPI/models"
)

// PostCustomerCreatedCode is the HTTP code returned for type PostCustomerCreated
const PostCustomerCreatedCode int = 201

/*PostCustomerCreated Customer created

swagger:response postCustomerCreated
*/
type PostCustomerCreated struct {
}

// NewPostCustomerCreated creates PostCustomerCreated with default headers values
func NewPostCustomerCreated() *PostCustomerCreated {

	return &PostCustomerCreated{}
}

// WriteResponse to the client
func (o *PostCustomerCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostCustomerInternalServerErrorCode is the HTTP code returned for type PostCustomerInternalServerError
const PostCustomerInternalServerErrorCode int = 500

/*PostCustomerInternalServerError Internal Server Error

swagger:response postCustomerInternalServerError
*/
type PostCustomerInternalServerError struct {
}

// NewPostCustomerInternalServerError creates PostCustomerInternalServerError with default headers values
func NewPostCustomerInternalServerError() *PostCustomerInternalServerError {

	return &PostCustomerInternalServerError{}
}

// WriteResponse to the client
func (o *PostCustomerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

/*PostCustomerDefault error

swagger:response postCustomerDefault
*/
type PostCustomerDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostCustomerDefault creates PostCustomerDefault with default headers values
func NewPostCustomerDefault(code int) *PostCustomerDefault {
	if code <= 0 {
		code = 500
	}

	return &PostCustomerDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post customer default response
func (o *PostCustomerDefault) WithStatusCode(code int) *PostCustomerDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post customer default response
func (o *PostCustomerDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post customer default response
func (o *PostCustomerDefault) WithPayload(payload *models.Error) *PostCustomerDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post customer default response
func (o *PostCustomerDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCustomerDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}