// Code generated by go-swagger; DO NOT EDIT.

package cat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/pk80/student/transactions/models"
)

// DeleteBulkOKCode is the HTTP code returned for type DeleteBulkOK
const DeleteBulkOKCode int = 200

/*DeleteBulkOK OK

swagger:response deleteBulkOK
*/
type DeleteBulkOK struct {
}

// NewDeleteBulkOK creates DeleteBulkOK with default headers values
func NewDeleteBulkOK() *DeleteBulkOK {

	return &DeleteBulkOK{}
}

// WriteResponse to the client
func (o *DeleteBulkOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteBulkCreatedCode is the HTTP code returned for type DeleteBulkCreated
const DeleteBulkCreatedCode int = 201

/*DeleteBulkCreated Created

swagger:response deleteBulkCreated
*/
type DeleteBulkCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Cat `json:"body,omitempty"`
}

// NewDeleteBulkCreated creates DeleteBulkCreated with default headers values
func NewDeleteBulkCreated() *DeleteBulkCreated {

	return &DeleteBulkCreated{}
}

// WithPayload adds the payload to the delete bulk created response
func (o *DeleteBulkCreated) WithPayload(payload *models.Cat) *DeleteBulkCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete bulk created response
func (o *DeleteBulkCreated) SetPayload(payload *models.Cat) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBulkCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteBulkNoContentCode is the HTTP code returned for type DeleteBulkNoContent
const DeleteBulkNoContentCode int = 204

/*DeleteBulkNoContent No Content

swagger:response deleteBulkNoContent
*/
type DeleteBulkNoContent struct {
}

// NewDeleteBulkNoContent creates DeleteBulkNoContent with default headers values
func NewDeleteBulkNoContent() *DeleteBulkNoContent {

	return &DeleteBulkNoContent{}
}

// WriteResponse to the client
func (o *DeleteBulkNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteBulkPartialContentCode is the HTTP code returned for type DeleteBulkPartialContent
const DeleteBulkPartialContentCode int = 206

/*DeleteBulkPartialContent Partial Content

swagger:response deleteBulkPartialContent
*/
type DeleteBulkPartialContent struct {

	/*
	  In: Body
	*/
	Payload *models.Cat `json:"body,omitempty"`
}

// NewDeleteBulkPartialContent creates DeleteBulkPartialContent with default headers values
func NewDeleteBulkPartialContent() *DeleteBulkPartialContent {

	return &DeleteBulkPartialContent{}
}

// WithPayload adds the payload to the delete bulk partial content response
func (o *DeleteBulkPartialContent) WithPayload(payload *models.Cat) *DeleteBulkPartialContent {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete bulk partial content response
func (o *DeleteBulkPartialContent) SetPayload(payload *models.Cat) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBulkPartialContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(206)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteBulkInternalServerErrorCode is the HTTP code returned for type DeleteBulkInternalServerError
const DeleteBulkInternalServerErrorCode int = 500

/*DeleteBulkInternalServerError Internal Server Error

swagger:response deleteBulkInternalServerError
*/
type DeleteBulkInternalServerError struct {
}

// NewDeleteBulkInternalServerError creates DeleteBulkInternalServerError with default headers values
func NewDeleteBulkInternalServerError() *DeleteBulkInternalServerError {

	return &DeleteBulkInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteBulkInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
