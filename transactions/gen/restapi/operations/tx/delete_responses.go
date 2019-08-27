// Code generated by go-swagger; DO NOT EDIT.

package tx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteOKCode is the HTTP code returned for type DeleteOK
const DeleteOKCode int = 200

/*DeleteOK OK

swagger:response deleteOK
*/
type DeleteOK struct {
}

// NewDeleteOK creates DeleteOK with default headers values
func NewDeleteOK() *DeleteOK {

	return &DeleteOK{}
}

// WriteResponse to the client
func (o *DeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteNoContentCode is the HTTP code returned for type DeleteNoContent
const DeleteNoContentCode int = 204

/*DeleteNoContent Deleted

swagger:response deleteNoContent
*/
type DeleteNoContent struct {
}

// NewDeleteNoContent creates DeleteNoContent with default headers values
func NewDeleteNoContent() *DeleteNoContent {

	return &DeleteNoContent{}
}

// WriteResponse to the client
func (o *DeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteBadRequestCode is the HTTP code returned for type DeleteBadRequest
const DeleteBadRequestCode int = 400

/*DeleteBadRequest Bad Request

swagger:response deleteBadRequest
*/
type DeleteBadRequest struct {
}

// NewDeleteBadRequest creates DeleteBadRequest with default headers values
func NewDeleteBadRequest() *DeleteBadRequest {

	return &DeleteBadRequest{}
}

// WriteResponse to the client
func (o *DeleteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// DeleteInternalServerErrorCode is the HTTP code returned for type DeleteInternalServerError
const DeleteInternalServerErrorCode int = 500

/*DeleteInternalServerError Internal Server Error

swagger:response deleteInternalServerError
*/
type DeleteInternalServerError struct {
}

// NewDeleteInternalServerError creates DeleteInternalServerError with default headers values
func NewDeleteInternalServerError() *DeleteInternalServerError {

	return &DeleteInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}