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

	"github.com/netbox-community/go-netbox/netbox/models"
)

// UsersTokensBulkUpdateReader is a Reader for the UsersTokensBulkUpdate structure.
type UsersTokensBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersTokensBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersTokensBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersTokensBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersTokensBulkUpdateOK creates a UsersTokensBulkUpdateOK with default headers values
func NewUsersTokensBulkUpdateOK() *UsersTokensBulkUpdateOK {
	return &UsersTokensBulkUpdateOK{}
}

/* UsersTokensBulkUpdateOK describes a response with status code 200, with default header values.

UsersTokensBulkUpdateOK users tokens bulk update o k
*/
type UsersTokensBulkUpdateOK struct {
	Payload *models.Token
}

func (o *UsersTokensBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /users/tokens/][%d] usersTokensBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *UsersTokensBulkUpdateOK) GetPayload() *models.Token {
	return o.Payload
}

func (o *UsersTokensBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersTokensBulkUpdateDefault creates a UsersTokensBulkUpdateDefault with default headers values
func NewUsersTokensBulkUpdateDefault(code int) *UsersTokensBulkUpdateDefault {
	return &UsersTokensBulkUpdateDefault{
		_statusCode: code,
	}
}

/* UsersTokensBulkUpdateDefault describes a response with status code -1, with default header values.

UsersTokensBulkUpdateDefault users tokens bulk update default
*/
type UsersTokensBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users tokens bulk update default response
func (o *UsersTokensBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *UsersTokensBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /users/tokens/][%d] users_tokens_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *UsersTokensBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersTokensBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
