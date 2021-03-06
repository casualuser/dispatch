///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware/dispatch/pkg/identity-manager/gen/models"
)

// UpdatePolicyOKCode is the HTTP code returned for type UpdatePolicyOK
const UpdatePolicyOKCode int = 200

/*UpdatePolicyOK Successful update

swagger:response updatePolicyOK
*/
type UpdatePolicyOK struct {

	/*
	  In: Body
	*/
	Payload *models.Policy `json:"body,omitempty"`
}

// NewUpdatePolicyOK creates UpdatePolicyOK with default headers values
func NewUpdatePolicyOK() *UpdatePolicyOK {

	return &UpdatePolicyOK{}
}

// WithPayload adds the payload to the update policy o k response
func (o *UpdatePolicyOK) WithPayload(payload *models.Policy) *UpdatePolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update policy o k response
func (o *UpdatePolicyOK) SetPayload(payload *models.Policy) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePolicyBadRequestCode is the HTTP code returned for type UpdatePolicyBadRequest
const UpdatePolicyBadRequestCode int = 400

/*UpdatePolicyBadRequest Invalid input

swagger:response updatePolicyBadRequest
*/
type UpdatePolicyBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdatePolicyBadRequest creates UpdatePolicyBadRequest with default headers values
func NewUpdatePolicyBadRequest() *UpdatePolicyBadRequest {

	return &UpdatePolicyBadRequest{}
}

// WithPayload adds the payload to the update policy bad request response
func (o *UpdatePolicyBadRequest) WithPayload(payload *models.Error) *UpdatePolicyBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update policy bad request response
func (o *UpdatePolicyBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePolicyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePolicyNotFoundCode is the HTTP code returned for type UpdatePolicyNotFound
const UpdatePolicyNotFoundCode int = 404

/*UpdatePolicyNotFound Policy not found

swagger:response updatePolicyNotFound
*/
type UpdatePolicyNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdatePolicyNotFound creates UpdatePolicyNotFound with default headers values
func NewUpdatePolicyNotFound() *UpdatePolicyNotFound {

	return &UpdatePolicyNotFound{}
}

// WithPayload adds the payload to the update policy not found response
func (o *UpdatePolicyNotFound) WithPayload(payload *models.Error) *UpdatePolicyNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update policy not found response
func (o *UpdatePolicyNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePolicyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePolicyInternalServerErrorCode is the HTTP code returned for type UpdatePolicyInternalServerError
const UpdatePolicyInternalServerErrorCode int = 500

/*UpdatePolicyInternalServerError Internal error

swagger:response updatePolicyInternalServerError
*/
type UpdatePolicyInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdatePolicyInternalServerError creates UpdatePolicyInternalServerError with default headers values
func NewUpdatePolicyInternalServerError() *UpdatePolicyInternalServerError {

	return &UpdatePolicyInternalServerError{}
}

// WithPayload adds the payload to the update policy internal server error response
func (o *UpdatePolicyInternalServerError) WithPayload(payload *models.Error) *UpdatePolicyInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update policy internal server error response
func (o *UpdatePolicyInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePolicyInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
