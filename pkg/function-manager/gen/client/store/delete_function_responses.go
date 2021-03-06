///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/dispatch/pkg/function-manager/gen/models"
)

// DeleteFunctionReader is a Reader for the DeleteFunction structure.
type DeleteFunctionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFunctionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteFunctionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteFunctionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteFunctionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteFunctionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteFunctionOK creates a DeleteFunctionOK with default headers values
func NewDeleteFunctionOK() *DeleteFunctionOK {
	return &DeleteFunctionOK{}
}

/*DeleteFunctionOK handles this case with default header values.

Successful operation
*/
type DeleteFunctionOK struct {
	Payload *models.Function
}

func (o *DeleteFunctionOK) Error() string {
	return fmt.Sprintf("[DELETE /function/{functionName}][%d] deleteFunctionOK  %+v", 200, o.Payload)
}

func (o *DeleteFunctionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Function)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFunctionBadRequest creates a DeleteFunctionBadRequest with default headers values
func NewDeleteFunctionBadRequest() *DeleteFunctionBadRequest {
	return &DeleteFunctionBadRequest{}
}

/*DeleteFunctionBadRequest handles this case with default header values.

Invalid Name supplied
*/
type DeleteFunctionBadRequest struct {
	Payload *models.Error
}

func (o *DeleteFunctionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /function/{functionName}][%d] deleteFunctionBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFunctionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFunctionNotFound creates a DeleteFunctionNotFound with default headers values
func NewDeleteFunctionNotFound() *DeleteFunctionNotFound {
	return &DeleteFunctionNotFound{}
}

/*DeleteFunctionNotFound handles this case with default header values.

Function not found
*/
type DeleteFunctionNotFound struct {
	Payload *models.Error
}

func (o *DeleteFunctionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /function/{functionName}][%d] deleteFunctionNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFunctionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFunctionInternalServerError creates a DeleteFunctionInternalServerError with default headers values
func NewDeleteFunctionInternalServerError() *DeleteFunctionInternalServerError {
	return &DeleteFunctionInternalServerError{}
}

/*DeleteFunctionInternalServerError handles this case with default header values.

Internal error
*/
type DeleteFunctionInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteFunctionInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /function/{functionName}][%d] deleteFunctionInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFunctionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
