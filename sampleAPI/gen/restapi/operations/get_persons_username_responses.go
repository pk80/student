// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/pk80/student/sampleAPI/gen/models"
)

// GetPersonsUsernameOKCode is the HTTP code returned for type GetPersonsUsernameOK
const GetPersonsUsernameOKCode int = 200

/*GetPersonsUsernameOK A person.

swagger:response getPersonsUsernameOK
*/
type GetPersonsUsernameOK struct {

	/*
	  In: Body
	*/
	Payload *models.Person `json:"body,omitempty"`
}

// NewGetPersonsUsernameOK creates GetPersonsUsernameOK with default headers values
func NewGetPersonsUsernameOK() *GetPersonsUsernameOK {

	return &GetPersonsUsernameOK{}
}

// WithPayload adds the payload to the get persons username o k response
func (o *GetPersonsUsernameOK) WithPayload(payload *models.Person) *GetPersonsUsernameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get persons username o k response
func (o *GetPersonsUsernameOK) SetPayload(payload *models.Person) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPersonsUsernameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPersonsUsernameNotFoundCode is the HTTP code returned for type GetPersonsUsernameNotFound
const GetPersonsUsernameNotFoundCode int = 404

/*GetPersonsUsernameNotFound The person doesn't exist.

swagger:response getPersonsUsernameNotFound
*/
type GetPersonsUsernameNotFound struct {
}

// NewGetPersonsUsernameNotFound creates GetPersonsUsernameNotFound with default headers values
func NewGetPersonsUsernameNotFound() *GetPersonsUsernameNotFound {

	return &GetPersonsUsernameNotFound{}
}

// WriteResponse to the client
func (o *GetPersonsUsernameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
