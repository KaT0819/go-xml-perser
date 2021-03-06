// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetTestOKCode is the HTTP code returned for type GetTestOK
const GetTestOKCode int = 200

/*GetTestOK Success

swagger:response getTestOK
*/
type GetTestOK struct {
	/*

	 */
	AccessControlAllowHeaders string `json:"Access-Control-Allow-Headers"`
	/*

	 */
	AccessControlAllowMethods string `json:"Access-Control-Allow-Methods"`
	/*

	 */
	AccessControlAllowOrigin string `json:"Access-Control-Allow-Origin"`

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetTestOK creates GetTestOK with default headers values
func NewGetTestOK() *GetTestOK {

	return &GetTestOK{}
}

// WithAccessControlAllowHeaders adds the accessControlAllowHeaders to the get test o k response
func (o *GetTestOK) WithAccessControlAllowHeaders(accessControlAllowHeaders string) *GetTestOK {
	o.AccessControlAllowHeaders = accessControlAllowHeaders
	return o
}

// SetAccessControlAllowHeaders sets the accessControlAllowHeaders to the get test o k response
func (o *GetTestOK) SetAccessControlAllowHeaders(accessControlAllowHeaders string) {
	o.AccessControlAllowHeaders = accessControlAllowHeaders
}

// WithAccessControlAllowMethods adds the accessControlAllowMethods to the get test o k response
func (o *GetTestOK) WithAccessControlAllowMethods(accessControlAllowMethods string) *GetTestOK {
	o.AccessControlAllowMethods = accessControlAllowMethods
	return o
}

// SetAccessControlAllowMethods sets the accessControlAllowMethods to the get test o k response
func (o *GetTestOK) SetAccessControlAllowMethods(accessControlAllowMethods string) {
	o.AccessControlAllowMethods = accessControlAllowMethods
}

// WithAccessControlAllowOrigin adds the accessControlAllowOrigin to the get test o k response
func (o *GetTestOK) WithAccessControlAllowOrigin(accessControlAllowOrigin string) *GetTestOK {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
	return o
}

// SetAccessControlAllowOrigin sets the accessControlAllowOrigin to the get test o k response
func (o *GetTestOK) SetAccessControlAllowOrigin(accessControlAllowOrigin string) {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
}

// WithPayload adds the payload to the get test o k response
func (o *GetTestOK) WithPayload(payload string) *GetTestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get test o k response
func (o *GetTestOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Access-Control-Allow-Headers

	accessControlAllowHeaders := o.AccessControlAllowHeaders
	if accessControlAllowHeaders != "" {
		rw.Header().Set("Access-Control-Allow-Headers", accessControlAllowHeaders)
	}

	// response header Access-Control-Allow-Methods

	accessControlAllowMethods := o.AccessControlAllowMethods
	if accessControlAllowMethods != "" {
		rw.Header().Set("Access-Control-Allow-Methods", accessControlAllowMethods)
	}

	// response header Access-Control-Allow-Origin

	accessControlAllowOrigin := o.AccessControlAllowOrigin
	if accessControlAllowOrigin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", accessControlAllowOrigin)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetTestUnauthorizedCode is the HTTP code returned for type GetTestUnauthorized
const GetTestUnauthorizedCode int = 401

/*GetTestUnauthorized Unauthorized

swagger:response getTestUnauthorized
*/
type GetTestUnauthorized struct {
	/*

	 */
	AccessControlAllowHeaders string `json:"Access-Control-Allow-Headers"`
	/*

	 */
	AccessControlAllowMethods string `json:"Access-Control-Allow-Methods"`
	/*

	 */
	AccessControlAllowOrigin string `json:"Access-Control-Allow-Origin"`

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetTestUnauthorized creates GetTestUnauthorized with default headers values
func NewGetTestUnauthorized() *GetTestUnauthorized {

	return &GetTestUnauthorized{}
}

// WithAccessControlAllowHeaders adds the accessControlAllowHeaders to the get test unauthorized response
func (o *GetTestUnauthorized) WithAccessControlAllowHeaders(accessControlAllowHeaders string) *GetTestUnauthorized {
	o.AccessControlAllowHeaders = accessControlAllowHeaders
	return o
}

// SetAccessControlAllowHeaders sets the accessControlAllowHeaders to the get test unauthorized response
func (o *GetTestUnauthorized) SetAccessControlAllowHeaders(accessControlAllowHeaders string) {
	o.AccessControlAllowHeaders = accessControlAllowHeaders
}

// WithAccessControlAllowMethods adds the accessControlAllowMethods to the get test unauthorized response
func (o *GetTestUnauthorized) WithAccessControlAllowMethods(accessControlAllowMethods string) *GetTestUnauthorized {
	o.AccessControlAllowMethods = accessControlAllowMethods
	return o
}

// SetAccessControlAllowMethods sets the accessControlAllowMethods to the get test unauthorized response
func (o *GetTestUnauthorized) SetAccessControlAllowMethods(accessControlAllowMethods string) {
	o.AccessControlAllowMethods = accessControlAllowMethods
}

// WithAccessControlAllowOrigin adds the accessControlAllowOrigin to the get test unauthorized response
func (o *GetTestUnauthorized) WithAccessControlAllowOrigin(accessControlAllowOrigin string) *GetTestUnauthorized {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
	return o
}

// SetAccessControlAllowOrigin sets the accessControlAllowOrigin to the get test unauthorized response
func (o *GetTestUnauthorized) SetAccessControlAllowOrigin(accessControlAllowOrigin string) {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
}

// WithPayload adds the payload to the get test unauthorized response
func (o *GetTestUnauthorized) WithPayload(payload string) *GetTestUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get test unauthorized response
func (o *GetTestUnauthorized) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTestUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Access-Control-Allow-Headers

	accessControlAllowHeaders := o.AccessControlAllowHeaders
	if accessControlAllowHeaders != "" {
		rw.Header().Set("Access-Control-Allow-Headers", accessControlAllowHeaders)
	}

	// response header Access-Control-Allow-Methods

	accessControlAllowMethods := o.AccessControlAllowMethods
	if accessControlAllowMethods != "" {
		rw.Header().Set("Access-Control-Allow-Methods", accessControlAllowMethods)
	}

	// response header Access-Control-Allow-Origin

	accessControlAllowOrigin := o.AccessControlAllowOrigin
	if accessControlAllowOrigin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", accessControlAllowOrigin)
	}

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetTestNotFoundCode is the HTTP code returned for type GetTestNotFound
const GetTestNotFoundCode int = 404

/*GetTestNotFound Not Found

swagger:response getTestNotFound
*/
type GetTestNotFound struct {
	/*

	 */
	AccessControlAllowHeaders string `json:"Access-Control-Allow-Headers"`
	/*

	 */
	AccessControlAllowMethods string `json:"Access-Control-Allow-Methods"`
	/*

	 */
	AccessControlAllowOrigin string `json:"Access-Control-Allow-Origin"`
}

// NewGetTestNotFound creates GetTestNotFound with default headers values
func NewGetTestNotFound() *GetTestNotFound {

	return &GetTestNotFound{}
}

// WithAccessControlAllowHeaders adds the accessControlAllowHeaders to the get test not found response
func (o *GetTestNotFound) WithAccessControlAllowHeaders(accessControlAllowHeaders string) *GetTestNotFound {
	o.AccessControlAllowHeaders = accessControlAllowHeaders
	return o
}

// SetAccessControlAllowHeaders sets the accessControlAllowHeaders to the get test not found response
func (o *GetTestNotFound) SetAccessControlAllowHeaders(accessControlAllowHeaders string) {
	o.AccessControlAllowHeaders = accessControlAllowHeaders
}

// WithAccessControlAllowMethods adds the accessControlAllowMethods to the get test not found response
func (o *GetTestNotFound) WithAccessControlAllowMethods(accessControlAllowMethods string) *GetTestNotFound {
	o.AccessControlAllowMethods = accessControlAllowMethods
	return o
}

// SetAccessControlAllowMethods sets the accessControlAllowMethods to the get test not found response
func (o *GetTestNotFound) SetAccessControlAllowMethods(accessControlAllowMethods string) {
	o.AccessControlAllowMethods = accessControlAllowMethods
}

// WithAccessControlAllowOrigin adds the accessControlAllowOrigin to the get test not found response
func (o *GetTestNotFound) WithAccessControlAllowOrigin(accessControlAllowOrigin string) *GetTestNotFound {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
	return o
}

// SetAccessControlAllowOrigin sets the accessControlAllowOrigin to the get test not found response
func (o *GetTestNotFound) SetAccessControlAllowOrigin(accessControlAllowOrigin string) {
	o.AccessControlAllowOrigin = accessControlAllowOrigin
}

// WriteResponse to the client
func (o *GetTestNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Access-Control-Allow-Headers

	accessControlAllowHeaders := o.AccessControlAllowHeaders
	if accessControlAllowHeaders != "" {
		rw.Header().Set("Access-Control-Allow-Headers", accessControlAllowHeaders)
	}

	// response header Access-Control-Allow-Methods

	accessControlAllowMethods := o.AccessControlAllowMethods
	if accessControlAllowMethods != "" {
		rw.Header().Set("Access-Control-Allow-Methods", accessControlAllowMethods)
	}

	// response header Access-Control-Allow-Origin

	accessControlAllowOrigin := o.AccessControlAllowOrigin
	if accessControlAllowOrigin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", accessControlAllowOrigin)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
