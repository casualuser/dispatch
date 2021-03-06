///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/dispatch/pkg/application-manager/gen/models"
)

// NewAddAppParams creates a new AddAppParams object
// with the default values initialized.
func NewAddAppParams() *AddAppParams {
	var ()
	return &AddAppParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddAppParamsWithTimeout creates a new AddAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddAppParamsWithTimeout(timeout time.Duration) *AddAppParams {
	var ()
	return &AddAppParams{

		timeout: timeout,
	}
}

// NewAddAppParamsWithContext creates a new AddAppParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddAppParamsWithContext(ctx context.Context) *AddAppParams {
	var ()
	return &AddAppParams{

		Context: ctx,
	}
}

// NewAddAppParamsWithHTTPClient creates a new AddAppParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddAppParamsWithHTTPClient(client *http.Client) *AddAppParams {
	var ()
	return &AddAppParams{
		HTTPClient: client,
	}
}

/*AddAppParams contains all the parameters to send to the API endpoint
for the add app operation typically these are written to a http.Request
*/
type AddAppParams struct {

	/*Body
	  Application object

	*/
	Body *models.Application

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add app params
func (o *AddAppParams) WithTimeout(timeout time.Duration) *AddAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add app params
func (o *AddAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add app params
func (o *AddAppParams) WithContext(ctx context.Context) *AddAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add app params
func (o *AddAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add app params
func (o *AddAppParams) WithHTTPClient(client *http.Client) *AddAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add app params
func (o *AddAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add app params
func (o *AddAppParams) WithBody(body *models.Application) *AddAppParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add app params
func (o *AddAppParams) SetBody(body *models.Application) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
