// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersTokensBulkDeleteReader is a Reader for the UsersTokensBulkDelete structure.
type UsersTokensBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersTokensBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUsersTokensBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersTokensBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersTokensBulkDeleteNoContent creates a UsersTokensBulkDeleteNoContent with default headers values
func NewUsersTokensBulkDeleteNoContent() *UsersTokensBulkDeleteNoContent {
	return &UsersTokensBulkDeleteNoContent{}
}

/*
UsersTokensBulkDeleteNoContent describes a response with status code 204, with default header values.

UsersTokensBulkDeleteNoContent users tokens bulk delete no content
*/
type UsersTokensBulkDeleteNoContent struct {
}

// IsSuccess returns true when this users tokens bulk delete no content response has a 2xx status code
func (o *UsersTokensBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users tokens bulk delete no content response has a 3xx status code
func (o *UsersTokensBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users tokens bulk delete no content response has a 4xx status code
func (o *UsersTokensBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this users tokens bulk delete no content response has a 5xx status code
func (o *UsersTokensBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this users tokens bulk delete no content response a status code equal to that given
func (o *UsersTokensBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UsersTokensBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/tokens/][%d] usersTokensBulkDeleteNoContent ", 204)
}

func (o *UsersTokensBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /users/tokens/][%d] usersTokensBulkDeleteNoContent ", 204)
}

func (o *UsersTokensBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUsersTokensBulkDeleteDefault creates a UsersTokensBulkDeleteDefault with default headers values
func NewUsersTokensBulkDeleteDefault(code int) *UsersTokensBulkDeleteDefault {
	return &UsersTokensBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
UsersTokensBulkDeleteDefault describes a response with status code -1, with default header values.

UsersTokensBulkDeleteDefault users tokens bulk delete default
*/
type UsersTokensBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users tokens bulk delete default response
func (o *UsersTokensBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this users tokens bulk delete default response has a 2xx status code
func (o *UsersTokensBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users tokens bulk delete default response has a 3xx status code
func (o *UsersTokensBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users tokens bulk delete default response has a 4xx status code
func (o *UsersTokensBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users tokens bulk delete default response has a 5xx status code
func (o *UsersTokensBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users tokens bulk delete default response a status code equal to that given
func (o *UsersTokensBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsersTokensBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /users/tokens/][%d] users_tokens_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *UsersTokensBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /users/tokens/][%d] users_tokens_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *UsersTokensBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersTokensBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}