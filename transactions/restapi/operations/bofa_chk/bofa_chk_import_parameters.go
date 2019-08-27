// Code generated by go-swagger; DO NOT EDIT.

package bofa_chk

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// NewBofaChkImportParams creates a new BofaChkImportParams object
// no default values defined in spec.
func NewBofaChkImportParams() BofaChkImportParams {

	return BofaChkImportParams{}
}

// BofaChkImportParams contains all the bound params for the bofa chk import operation
// typically these are obtained from a http.Request
//
// swagger:parameters bofaChkImport
type BofaChkImportParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: formData
	*/
	Csvfile io.ReadCloser
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewBofaChkImportParams() beforehand.
func (o *BofaChkImportParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}

	csvfile, csvfileHeader, err := r.FormFile("csvfile")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "csvfile", err))
	} else if err := o.bindCsvfile(csvfile, csvfileHeader); err != nil {
		// Required: true
		res = append(res, err)
	} else {
		o.Csvfile = &runtime.File{Data: csvfile, Header: csvfileHeader}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCsvfile binds file parameter Csvfile.
//
// The only supported validations on files are MinLength and MaxLength
func (o *BofaChkImportParams) bindCsvfile(file multipart.File, header *multipart.FileHeader) error {
	return nil
}